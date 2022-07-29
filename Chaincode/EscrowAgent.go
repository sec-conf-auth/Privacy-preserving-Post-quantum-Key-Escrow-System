package main

import (
	//"crypto/aes"
	//"crypto/cipher"
	"fmt"
	"strconv"
	//"encoding/pem"
	"time"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/open-quantum-safe/liboqs-go/oqs"
	_"log"
	_"io"
	"encoding/binary"
	"math"
	//"encoding/base64"
	_"crypto/x509"
	_"crypto/ecdsa"
	_"crypto/rsa"
	"crypto/sha256"
	_"math/big"
	_"crypto/elliptic"
	_"encoding/pem"
   	"math/big"
    	//"crypto/x509"
	"bytes"
	"encoding/base64" 
	math_rand "math/rand"
//	"encoding/hex"
//	"encoding/json"
//	"encoding/asn1"
	"log"
)

func init() {
    math_rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[math_rand.Intn(len(letterRunes))]
    }
    return string(b)
}


type Dec struct{}

//////////////////////////////////////////////////////////////////////////////
//marshal evidence to bytes
//////////////////////////////////////////////////////////////////////////////
func marshal(evidence *ZKEvidence,root []byte)[]byte{
	evidenceList := make([][]byte, 7)
	evidenceList[0]=evidence.Salt
	evidenceList[1]=evidence.RawData
	x:=len(evidence.MerkleSibling)
	y:=len(evidence.MerkleSibling[0])
	z:=x*y
	key := make([]byte, z)
	m:=0
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
		key[m]=evidence.MerkleSibling[i][j]
		m++         }
		}
	xbyte:=IntToBytes(x)
	ybyte:=IntToBytes(y)
	evidenceList[2]=key
	evidenceList[3]=xbyte
	evidenceList[4]=ybyte
	b:=make([]byte,8)
	binary.BigEndian.PutUint64(b, evidence.Index)
	evidenceList[5]=b
	evidenceList[6]=evidence.MerkleRoot
	evidencecombineBytes := bytes.Join([][]byte{evidenceList[0],evidenceList[1],evidenceList[2],evidenceList[3],evidenceList[4],evidenceList[5],evidenceList[6]}, []byte("+++++"))
	combineBytes:=bytes.Join([][]byte{evidencecombineBytes,root}, []byte("======"))
	return combineBytes
}

//////////////////////////////////////////////////////////////////////////////
//unmarshal bytes to evidence
//////////////////////////////////////////////////////////////////////////////
func unmarshal(combineBytes []byte)(*ZKEvidence,[]byte){
	enConBytes := bytes.Split(combineBytes, []byte("======"))
	root:=enConBytes[1]
	evidenceList:=bytes.Split(enConBytes[0], []byte("+++++"))
	x:=BytesToInt(evidenceList[3])
	y:=BytesToInt(evidenceList[4])
	key := evidenceList[2]
	run := make([][]byte, x)
	// 
	for i := range run {
	run[i] = make([]byte, y)
	}
	m:=0
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			run[i][j]=key[m]
			m++
			}
				}
	index:=uint64(binary.BigEndian.Uint64(evidenceList[5]))
	return &ZKEvidence{
		Salt:          evidenceList[0],
		RawData:       evidenceList[1],
		MerkleSibling: run,
		Index:         index,
		MerkleRoot:    evidenceList[6],
	}, root
}

func IntToBytes(n int) []byte {
    data := int64(n)
    bytebuf := bytes.NewBuffer([]byte{})
    binary.Write(bytebuf, binary.BigEndian, data)
    return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
    bytebuff := bytes.NewBuffer(bys)
    var data int64
    binary.Read(bytebuff, binary.BigEndian, &data)
    return int(data)
}
//Generate Merkle Root
func GenerateZKMerkleRoot(dataList [][]byte, seed []byte) []byte {
	sequence := GenerateSequence256(seed, len(dataList))
	newDataList := make([][]byte, len(dataList))
	for i, data := range dataList {
		newDataList[i] = append(sequence[i], data...)
	}
	return CalcMerkleRoot(newDataList)
}

//Generate Merkle Evidence
func GenerateZKMerkleEvidence(dataList [][]byte, seed []byte, idx uint64) (*ZKEvidence, error) {
	sequence := GenerateSequence256(seed, len(dataList))
	newDataList := make([][]byte, len(dataList))
	for i, data := range dataList {
		newDataList[i] = append(sequence[i], data...)
	}
	evidence, err := CalcMerkleEvidence(newDataList, idx)
	if err != nil {
		return nil, err
	}
	return &ZKEvidence{
		Salt:          sequence[idx],
		RawData:       evidence.RawData[len(sequence[idx]):],
		MerkleSibling: evidence.MerkleSibling,
		Index:         evidence.Index,
		MerkleRoot:    evidence.MerkleRoot,
	}, nil
}

func ZKProve(evidence *ZKEvidence) bool {
	ev := &Evidence{
		RawData:       append(evidence.Salt, evidence.RawData...),
		MerkleSibling: evidence.MerkleSibling,
		Index:         evidence.Index,
		MerkleRoot:    evidence.MerkleRoot,
	}
	return Prove(ev)
}

type Evidence struct {
	RawData       []byte
	MerkleSibling [][]byte
	Index         uint64
	MerkleRoot    []byte
}
type ZKEvidence struct {
	Salt          []byte
	RawData       []byte
	MerkleSibling [][]byte
	Index         uint64
	MerkleRoot    []byte
}

func GenerateSequence256(seed []byte, count int) [][]byte {
	result := [][]byte{}
	current := seed
	for i := 0; i < count; i++ {
		current = getHash(current)
		result = append(result, current)
	}
	return result
}
func GenerateSequence128(seed []byte, count int) [][]byte {
	result := [][]byte{}
	current := seed
	for i := 0; i < count; i++ {
		current = getHash(current)
		result = append(result, current[0:16])
	}
	return result
}
func GenerateSequenceUInt64(seed []byte, count int) []uint64 {
	result := []uint64{}
	current := seed
	for i := 0; i < count; i++ {
		current = getHash(current)
		x := binary.BigEndian.Uint64(current[0:8])
		result = append(result, x)
	}
	return result
}

func nextPowerOfTwo(n int) int {
	// Return the number if it's already a power of 2.
	if n&(n-1) == 0 {
		return n
	}

	// Figure out and return the next power of two.
	exponent := uint64(math.Log2(float64(n))) + 1
	return 1 << exponent // 2^exponent
}
func getHash(input []byte) []byte {
	h := sha256.New()
	h.Write(input)
	return h.Sum(nil)
}

// hashMerkleBranches takes two hashes, treated as the left and right tree
// nodes, and returns the hash of their concatenation.  This is a helper
// function used to aid in the generation of a merkle tree.
func hashMerkleBranches(left []byte, right []byte) []byte {
	// Concatenate the left and right nodes.
	h := sha256.New()
	h.Write(left)
	h.Write(right)
	return h.Sum(nil)
}

// BuildMerkleTreeStore creates a merkle tree from a slice of transactions,
// stores it using a linear array, and returns a slice of the backing array.  A
// linear array was chosen as opposed to an actual tree structure since it uses
// about half as much memory.  The following describes a merkle tree and how it
// is stored in a linear array.
//
// A merkle tree is a tree in which every non-leaf node is the hash of its
// children nodes.  A diagram depicting how this works for bitcoin transactions
// where h(x) is a double sha256 follows:
//
//	         root = h1234 = h(h12 + h34)
//	        /                           \
//	  h12 = h(h1 + h2)            h34 = h(h3 + h4)
//	   /            \              /            \
//	h1 = h(tx1)  h2 = h(tx2)    h3 = h(tx3)  h4 = h(tx4)
//
// The above stored as a linear array is as follows:
//
// 	[h1 h2 h3 h4 h12 h34 root]
//
// As the above shows, the merkle root is always the last element in the array.
//
// The number of inputs is not always a power of two which results in a
// balanced tree structure as above.  In that case, parent nodes with no
// children are also zero and parent nodes with only a single left node
// are calculated by concatenating the left node with itself before hashing.
// Since this function uses nodes that are pointers to the hashes, empty nodes
// will be nil.
//
// The additional bool parameter indicates if we are generating the merkle tree
// using witness transaction id's rather than regular transaction id's. This
// also presents an additional case wherein the wtxid of the coinbase transaction
// is the zeroHash.
func BuildMerkleTreeStore(transactions [][]byte) [][]byte {
	// Calculate how many entries are required to hold the binary merkle
	// tree as a linear array and create an array of that size.
	nextPoT := nextPowerOfTwo(len(transactions))
	arraySize := nextPoT*2 - 1
	merkles := make([][]byte, arraySize)

	// Create the base transaction hashes and populate the array with them.
	for i, tx := range transactions {
		// If we're computing a witness merkle root, instead of the
		// regular txid, we use the modified wtxid which includes a
		// transaction's witness data within the digest. Additionally,
		// the coinbase's wtxid is all zeroes.
		merkles[i] = getHash(tx)
	}

	// Start the array offset after the last transaction and adjusted to the
	// next power of two.
	offset := nextPoT
	for i := 0; i < arraySize-1; i += 2 {
		switch {
		// When there is no left child node, the parent is nil too.
		case merkles[i] == nil:
			merkles[offset] = nil

		// When there is no right child, the parent is generated by
		// hashing the concatenation of the left child with itself.
		case merkles[i+1] == nil:
			newHash := hashMerkleBranches(merkles[i], merkles[i])
			merkles[offset] = newHash

		// The normal case sets the parent node to the double sha256
		// of the concatentation of the left and right children.
		default:
			newHash := hashMerkleBranches(merkles[i], merkles[i+1])
			merkles[offset] = newHash
		}
		offset++
	}

	return merkles
}

//
func CalcMerkleRoot(dataList [][]byte) []byte {
	hashes := BuildMerkleTreeStore(dataList)
	return hashes[len(hashes)-1]
}

// Prove the evidence
func Prove(evidence *Evidence) bool {
	txid := getHash(evidence.RawData)
	merkleRoot := evidence.MerkleRoot
	intermediateNodes := evidence.MerkleSibling
	index := evidence.Index
	// Shortcut the empty-block case
	if bytes.Equal(txid[:], merkleRoot[:]) && index == 0 && len(intermediateNodes) == 0 {
		return true
	}

	current := txid
	idx := index
	proofLength := len(intermediateNodes)

	numSteps := (proofLength)

	for i := 0; i < numSteps; i++ {
		next := intermediateNodes[i]
		if idx%2 == 1 {
			current = hashMerkleBranches(next, current)
		} else {
			current = hashMerkleBranches(current, next)
		}
		idx >>= 1
	}

	return bytes.Equal(current, merkleRoot)
}

//Calc evidencee
func CalcMerkleEvidence(dataList [][]byte, idx uint64) (*Evidence, error) {
	evidence := &Evidence{Index: idx, RawData: dataList[idx]}
	merkleList := BuildMerkleTreeStore(dataList)
	path := [][]byte{}
	dep := int(math.Log2(float64(len(merkleList) + 1)))

	evidence.MerkleRoot = merkleList[len(merkleList)-1]
	merkleList = merkleList[:len(merkleList)-1]
	for i := 1; i < dep; i++ {
		levelCount := int(math.Pow(2, float64(i)))           //
		levelList := merkleList[len(merkleList)-levelCount:] //
		merkleList = merkleList[:len(merkleList)-levelCount]
		mask := idx >> (dep - i - 1)
		if mask%2 == 0 { //left
			mask++ //
		} else {
			mask-- //
		}
		if levelList[mask] == nil { //
			mask--
		}
		path = leftJoin(levelList[mask], path)
	}
	evidence.MerkleSibling = path
	return evidence, nil
}
func leftJoin(n []byte, list [][]byte) [][]byte {
	result := make([][]byte, len(list)+1)
	result[0] = n
	for i, x := range list {
		result[i+1] = x
	}
	return result
}

func (t *Dec) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Success invoke and not opter!!"))
}
//////////////////////////////////////////////////////////////////////////////
//The main entrance of the chaincode including all the APIs that can be invoked by the clients and other chaincodes.
//////////////////////////////////////////////////////////////////////////////
func (t *Dec) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "Gen_EA_KeyPair" {
		return t.Gen_EA_KeyPair(stub, args)
	} else if fn == "Get_EA_PubKey" {
		return t.Get_EA_PubKey(stub, args)
	} else if fn =="Decap_Partial_SessKey"{
		return t.Decap_Partial_SessKey(stub, args)
	} else if fn=="Get_Partial_SessKey"{
		return t.Get_Partial_SessKey(stub, args)
	}
	return shim.Error("Recevied unkown function invocation")
}

//////////////////////////////////////////////////////////////////////////////
//API invocation interface function:Generate one post-quantum public/private key pair
//for the escrow agent, upload the public key to public
//ledger, and store the private key off-chain
//////////////////////////////////////////////////////////////////////////////
func (t *Dec) Gen_EA_KeyPair(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	t0 := time.Now()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// verify identity of the escrower 

	// We use generateKeyPair function to generate publicKey and privateKey about escrower.
	k,_:=strconv.Atoi(args[0])
	queryArgs := [][]byte{[]byte("Get_Global_Setup")}
	response := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	enConBytes:=bytes.Split(response.Payload, []byte(";;;;;"))
	kemName:=string(enConBytes[2])
	escrow := oqs.KeyEncapsulation{}
	defer escrow.Clean() 
	if err := escrow.Init(kemName, nil); err != nil {
		return shim.Error(err.Error())
	}
	publicKey, err := escrow.GenerateKeyPair()
	if err != nil {
		return shim.Error(err.Error())
	}
	privateKey := escrow.ExportSecretKey()
	strPriKey := base64.StdEncoding.EncodeToString(privateKey)
	elapsed := time.Since(t0)
	gentime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	t1:=time.Now()
	EsPKKeyid:=string(enConBytes[k+3])
	err = stub.PutState(EsPKKeyid, publicKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("publicKey:",len(publicKey))
	elapsed1 := time.Since(t1)
	putstatetime := strconv.FormatFloat(elapsed1.Seconds(), 'E', -1, 64)
	return shim.Success([]byte("------privateKey:" + strPriKey+"   The escrow has successfully decaped private Key and public Key!!!------generate keypairs time:"+ gentime+"Put public key on chain :"+putstatetime))
}

//////////////////////////////////////////////////////////////////////////////
//API invocation interface function:Read the escrow agent’s public key from public ledger
//////////////////////////////////////////////////////////////////////////////
func (t *Dec) Get_EA_PubKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	EsPKKeyid := args[0]
	publicKeyBytes, err := stub.GetState(EsPKKeyid)
	if err != nil {
		fmt.Println("Get state error")
		return shim.Error(err.Error())
	}
	return shim.Success(publicKeyBytes)
}
//encap to get sessionkey and CT
func GenEncapUploadSessionKey(algName_pq string,pubKey_pq []byte)([]byte,  []byte, error){
	kemName := algName_pq
	kemer := oqs.KeyEncapsulation{}
	defer kemer.Clean() // clean up even in case of panic
	if err := kemer.Init(kemName, nil); err != nil {
		log.Fatal(err)
		return nil,  nil, err
	}
	/////////////////////////////////////////////////////////////////////
	// Sessionkey is used to be encapsulated by PQ key Encapsulatioon
	/////////////////////////////////////////////////////////////////////
	fmt.Println("\nThe details of the KEM algorithm is:\n", kemer.Details().String())

	encapSessionKey, sessionKey, err := kemer.EncapSecret(pubKey_pq)
	if err != nil {
		log.Fatal(err)
		return nil,  nil, err
	}
	return encapSessionKey, sessionKey, nil
}

//use CT and prikey to decap to get sessionkey 
func DownloadDecapSessionKey(encapSessionKey []byte,  privateKey_pq_KE []byte,algName_pq string)([]byte, error){
	kemer := oqs.KeyEncapsulation{}
	defer kemer.Clean() // clean up even in case of panic
	if err := kemer.Init(algName_pq, privateKey_pq_KE); err != nil {
		log.Fatal(err)
		return nil, err
	}
	sessionKey, err := kemer.DecapSecret(encapSessionKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("\nsessionKey in DecapSecret [length=%d]:\n", len(sessionKey))
	return sessionKey, nil
}


//////////////////////////////////////////////////////////////////////////////
//API invocation interface function:Decapsulate to get one share (xi, yi) of partial session
//key, and store the share under the PK key
//////////////////////////////////////////////////////////////////////////////
func (t *Dec) Decap_Partial_SessKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	t0 := time.Now()
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	queryArgs := [][]byte{[]byte("Get_Global_Setup")}
	response2 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response2.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response2.Payload))
	}
	enConBytes1:=bytes.Split(response2.Payload, []byte(";;;;;"))
	collection:=string(enConBytes1[id+9])
	queryArgs = [][]byte{[]byte("Get_Anony_Setup"),[]byte(collection)}
	response1 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response1.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response1.Payload))
	}
	senderKeyid := string(response1.Payload)
	PartialKey := string(response1.Payload)
	kemName := string(enConBytes1[2])
	//////////////////////////////////////////////////////////////////////////////
	// Initialize the private key
	//////////////////////////////////////////////////////////////////////////////
	tMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error(fmt.Sprintf("Could not retrieve transient, err %s", err))
	}
	privateKeyBytes, ok := tMap["PRIVATEKEY"]
	if !ok {
		return shim.Error(fmt.Sprintf("Expected transient KS"))
	}
	elapsed := time.Since(t0)
	alltime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	//////////////////////////////////////////////////////////////////////////////
	// get the Supervisor's annonymized data
	//////////////////////////////////////////////////////////////////////////////
	t2:= time.Now()
	queryArgs = [][]byte{[]byte("Get_Anonymized_EK"), []byte(senderKeyid),[]byte(collection)}
	response := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	combinebytes:=bytes.Split(response.Payload, []byte(":::::"))
	t5:=time.Now()
	evidence2,root2:=unmarshal(combinebytes[1])
	if !bytes.Equal(evidence2.MerkleRoot, root2) {
		fmt.Printf("Root[%x]It's not signed by the public security departmentRoot[%x]", evidence2.MerkleRoot, root2)
	}
	pass := ZKProve(evidence2)
	if pass {
		fmt.Printf("Evidence verified! Assert「%s」is true", evidence2.RawData)
	} else {
		return shim.Error(fmt.Sprintf("Evidence verification failure"))
	}
	elapsed5 := time.Since(t5)
	verifyEvidenceTime := strconv.FormatFloat(elapsed5.Seconds(), 'E', -1, 64)
	enConBytesall := bytes.Split(combinebytes[0], []byte("-----"))
	//////////////////////////////////////////////////////////////////////////////
	//use DecapSecret function to decrypt sharedSecret
	//////////////////////////////////////////////////////////////////////////////
	enConBytesxy:=enConBytesall[0]
	enConBytes:=bytes.Split(enConBytesxy, []byte(";;;;;"))
	elapsed2 := time.Since(t2)
	getsenderdataTime := strconv.FormatFloat(elapsed2.Seconds(), 'E', -1, 64)
	t1 := time.Now()
	escrow := oqs.KeyEncapsulation{}
	defer escrow.Clean() 
	if err := escrow.Init(kemName, privateKeyBytes); err != nil {
		return shim.Error(err.Error())
	}
	encapSessionKey := enConBytes[2*(id-1)]
	sk, err := DownloadDecapSessionKey(encapSessionKey, privateKeyBytes, kemName)
	if err != nil {
		return shim.Error(err.Error())
	}
	elapsed1 := time.Since(t1)
	decapSecretTime := strconv.FormatFloat(elapsed1.Seconds(), 'E', -1, 64)
	t3 := time.Now()
	x:=big.NewInt(1)
	if len(sk) >= 32 {
	hash := sha256.New()
	hash.Write(sk)
	sk = hash.Sum(nil)
	}
	x=x.SetBytes(sk) 
	fmt.Println("x=",x)
	//////////////////////////////////////////////////////////////////////////////
	//get the Y value
	//////////////////////////////////////////////////////////////////////////////
	y:=big.NewInt(1)
	y=y.SetBytes(enConBytes[2*(id-1)+1]) 
	fmt.Println("y=",y)
	combineBytes :=bytes.Join([][]byte{sk, enConBytes[2*(id-1)+1]}, []byte("-----"))
	//////////////////////////////////////////////////////////////////////////////
	//put the X and Y on private collection
	//////////////////////////////////////////////////////////////////////////////
	err = stub.PutPrivateData(collection, PartialKey, combineBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("combineBytes:",len(combineBytes))
	elapsed3 := time.Since(t3)
	putonchainTime := strconv.FormatFloat(elapsed3.Seconds(), 'E', -1, 64)
	fmt.Println("prepare parameter:"+ alltime +"get sender data time:"+getsenderdataTime+"---verify Evidence Time:"+verifyEvidenceTime+"-------DecapSecretTime:"+decapSecretTime+"put on chain Time:"+putonchainTime)
	return shim.Success(combineBytes)
}

//////////////////////////////////////////////////////////////////////////////
//API invocation interface function:Read the share (xi, yi) under the PK key
//////////////////////////////////////////////////////////////////////////////
func (t *Dec) Get_Partial_SessKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	keyid := args[0]
	collection:=args[1]
	result, err := stub.GetPrivateData(collection,keyid)
	fmt.Println("result=",result)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}


func main() {
	err1 := shim.Start(new(Dec))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}
