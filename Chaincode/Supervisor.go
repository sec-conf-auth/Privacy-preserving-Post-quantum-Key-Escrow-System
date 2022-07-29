package main

import (
	"bytes"
	"strings"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strconv"
	"crypto/x509"
	"encoding/pem"
	//"time"
	"encoding/base64"
	"math/big"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"crypto/sha256"
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"log"
	"time"
	"encoding/binary"
	"math"
	_"github.com/Re-volution/sizestruct"
	crypt_rand "crypto/rand"
	"math/rand"
)

type Sup struct{}

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
//generate Merkle Root
func GenerateZKMerkleRoot(dataList [][]byte, seed []byte) []byte {
	sequence := GenerateSequence256(seed, len(dataList))
	newDataList := make([][]byte, len(dataList))
	for i, data := range dataList {
		newDataList[i] = append(sequence[i], data...)
	}
	return CalcMerkleRoot(newDataList)
}

//Genertate Merkle Evidence
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

func GenerateSequence256(seed []byte, Sup int) [][]byte {
	result := [][]byte{}
	current := seed
	for i := 0; i < Sup; i++ {
		current = getHash(current)
		result = append(result, current)
	}
	return result
}
func GenerateSequence128(seed []byte, Sup int) [][]byte {
	result := [][]byte{}
	current := seed
	for i := 0; i < Sup; i++ {
		current = getHash(current)
		result = append(result, current[0:16])
	}
	return result
}
func GenerateSequenceUInt64(seed []byte, Sup int) []uint64 {
	result := []uint64{}
	current := seed
	for i := 0; i < Sup; i++ {
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

// Pove the evidence
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

//
func CalcMerkleEvidence(dataList [][]byte, idx uint64) (*Evidence, error) {
	evidence := &Evidence{Index: idx, RawData: dataList[idx]}
	merkleList := BuildMerkleTreeStore(dataList)
	path := [][]byte{}
	dep := int(math.Log2(float64(len(merkleList) + 1)))

	evidence.MerkleRoot = merkleList[len(merkleList)-1]
	merkleList = merkleList[:len(merkleList)-1]
	for i := 1; i < dep; i++ {
		levelSup := int(math.Pow(2, float64(i)))           //
		levelList := merkleList[len(merkleList)-levelSup:] //
		merkleList = merkleList[:len(merkleList)-levelSup]
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

//////////////////////////////////////////////////////////////////////////////
//get creator of the user
//////////////////////////////////////////////////////////////////////////////
func GetCreator(stub shim.ChaincodeStubInterface) string {
	creatorByte, _ := stub.GetCreator()
	fmt.Println(string(creatorByte))
	certStart := bytes.IndexAny(creatorByte, "-----BEGIN")
	if certStart == -1 {
		fmt.Println("No certificate found")
		return ""	
	}
	certText := creatorByte[certStart:]
	bl, _ := pem.Decode(certText)
	if bl == nil {
		fmt.Println("Could not decode the PEM structure")
		return ""
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate failed")
		return ""
	}
	uname := cert.Subject.CommonName
	return uname
}
//use AES to decrypt the message
func AES_Decrypt(encryptedDate []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	privDate := make([]byte, len(encryptedDate))
	blockMode.CryptBlocks(privDate, encryptedDate)
	privDate = PKCS7UnPadding(privDate)
	return privDate, nil
}


func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


func (t *Sup) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Success invoke and not opter!!"))
}

//////////////////////////////////////////////////////////////////////////////
//The main entrance of the chaincode including all the APIs that can be invoked by the clients and other chaincodes.
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	 if fn == "Recover_Final_SessKey" {
		return t.Recover_Final_SessKey(stub, args)
	}else if fn == "Get_Final_SessKey" {
		return t.Get_Final_SessKey(stub, args)
	}else if fn == "Gen_Super_KeyPair"{
		return t.Gen_Super_KeyPair(stub, args)
	}else if fn =="Get_Super_PubKey"{
		return t.Get_Super_PubKey(stub, args)
	}else if fn=="Anonymization"{
		return t.Anonymization(stub, args)
	}else if fn=="Get_Anonymized_EK"{
		return t.Get_Anonymized_EK(stub, args)
	}else if fn=="Init_Global_Setup"{
		return t.Init_Global_Setup(stub, args)
	}else if fn=="Get_Global_Setup"{
		return t.Get_Global_Setup(stub, args)
	}else if fn=="Get_Global_Time"{
		return t.Get_Global_Time(stub,args)
	}else if fn=="Get_Anony_Setup"{
		return t.Get_Anony_Setup(stub,args)
	}
	return shim.Error("Recevied unkown function invocation")
}

type esPKKeyid struct{
	  x string;
}

type collection struct{
	  x string;
	 
}
type xyKeyid struct{
	  x string;
	 
}

type id struct{
	  x string; 
}
type ystring struct{
	  x string;
	 
}

type Data struct{
	  x *big.Int;
	  y *big.Int
}

type nameid struct{
	  x string;  
}

type property struct{
	  x string;  
}


//////////////////////////////////////////////////////////////////////////////
//Read key names of the public keys and Private Data
//Collection names from GS key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Global_Time(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	t1:=time.Now()
	fmt.Println("time",t1)
	x,y,z:=t1.Date()
	str1 := fmt.Sprintf("%d", x)
	str2 := fmt.Sprintf("%d", z)
	fmt.Println("year",str1)
	k:=y.String()
	if k=="January"{
	k="01"	
	}else if k=="February"{
	k="02"
	}else if k=="March"{
	k="03"
	}else if k=="April"{
	k="04"
	}else if k=="May"{
	k="05"
	}else if k=="June"{
	k="06"
	}else if k=="July"{
	k="07"
	}else if k=="August"{
	k="08"
	}else if k=="September"{
	k="09"
	}else if k=="October"{
	k="10"
	}else if k=="November"{
	k="11"
	}else if k=="December"{
	k="12"
	}
	fmt.Println("month",k)
	fmt.Println("day",str2)
	s:=str1+k+str2
	fmt.Println("time",s)
	return shim.Success([]byte(s))
}

//////////////////////////////////////////////////////////////////////////////
//Initialize GS key, under which key names of the public keys of supervisor and escrow agents together with
//supervisor-related Private Data Collection names are stored
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Init_Global_Setup(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	t0 := time.Now()
	t1,_:=strconv.Atoi(args[0])//t
	args0:=strconv.Itoa(t1)
	n,_:=strconv.Atoi(args[1])//n
	args1:=strconv.Itoa(n)
	args2:=args[2]//kem name
	args3:="SupPublicKey"//Sup pub key
	args4:="EsPublicKey01"	
	args5:="EsPublicKey02"
	args6:="EsPublicKey03"
	args7:="EsPublicKey04"
	args8:="EsPublicKey05"
	args9:="IS_PDC"
	args10:="EA1_S_PDC"
	args11:="EA2_S_PDC"
	args12:="EA3_S_PDC"
	args13:="EA4_S_PDC"
	args14:="EA5_S_PDC"
	Collection2:=make([]collection,5)
	Collection2[0].x=args10
	Collection2[1].x=args11
	Collection2[2].x=args12
	Collection2[3].x=args13
	Collection2[4].x=args14
	Collection1:=make([]collection,t1)
	for i:=0;i<t1;i++{
	m,_:=strconv.Atoi(args[i+3])//select which escrowAgent
	Collection1[i].x=Collection2[m-1].x
	}
	collectionBytes:=[]byte(Collection1[0].x)
	for i:=1;i<t1;i++{
	collectionBytes=bytes.Join([][]byte{[]byte(collectionBytes),[]byte(Collection1[i].x)},[]byte(";;;;;"))
	}
	combineBytes:=bytes.Join([][]byte{[]byte(args0),[]byte(args1),[]byte(args2),[]byte(args3),[]byte(args4),[]byte(args5),[]byte(args6),[]byte(args7),[]byte(args8),[]byte(args9),[]byte(args10),[]byte(args11),[]byte(args12),[]byte(args13),[]byte(args14),[]byte(collectionBytes)}, []byte(";;;;;"))
	_ = stub.PutState("GlobalSetup", combineBytes)
	elapsed := time.Since(t0)
	alltime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	fmt.Println("\nthe alltime is :",alltime)
	fmt.Println("combineBytes:",len(combineBytes))
	return shim.Success([]byte("Init Successfully! Init time:"+alltime))
}


//////////////////////////////////////////////////////////////////////////////
//Anonymize the source data record name, store the
//pseudonym pName under AS key and transfer part of
//the source data (relating to the escrow agents) under
//the EK keys
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Anonymization(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	t0 := time.Now()
	t1 := time.Now()
	collection1:=args[0]
	queryArgs := [][]byte{[]byte("Get_Record_Name"),[]byte(collection1)}
	response := stub.InvokeChaincode("DataSource", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
		}
	senderKeyid:=string(response.Payload)
	fmt.Println("\nthe senderKeyid is :",senderKeyid)
	k,_:=strconv.Atoi(args[1])
	Property:=make([]property,5)
	for i:=0;i<5;i++{
	Property[i].x=args[i+2]
	}
	response2,_:= stub.GetState("GlobalSetup")
	enConBytes:=bytes.Split(response2, []byte(";;;;;"))
	l1string:=string(enConBytes[1])
	l1,_:=strconv.Atoi(l1string)
	Collection:=make([]collection,l1)
	for i:=0;i<l1;i++{
	Collection[i].x=string(enConBytes[i+10])
	}
	queryArgs = [][]byte{[]byte("Get_Source_Data"),[]byte(senderKeyid),[]byte(collection1)}
	response1 := stub.InvokeChaincode("DataSource", queryArgs, "mychannel")
	if response1.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response1.Payload))
		}
	elapsed := time.Since(t1)
	getSecrettime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	t3 := time.Now()
	lenth:=len(senderKeyid)
	splitylen:=lenth-8
	splitkey:=strings.Split(senderKeyid,senderKeyid[splitylen:])
	args0:=splitkey[0]
	fmt.Println("\nthe args0 is :",args0)
	max := new(big.Int).Lsh(big.NewInt(1), 256)
	min := new(big.Int).Lsh(big.NewInt(1), 255)
	fmt.Println("\nthe max is ",max)
	fmt.Println("\nthe min is ",min)
	l:=-1
	k1:= big.NewInt(1)
	for l==-1{
	coeNumber, _ := crypt_rand.Int(crypt_rand.Reader, max)
	fmt.Println("\nthe coeNumber is :",coeNumber)
	k1=coeNumber
	l=k1.Cmp(min)
	}
	m:=k1.String()
	//////////////////////////////////////////////////////////////////////////////
	//Generate a false name
	//////////////////////////////////////////////////////////////////////////////
	s:=args0+m
    	h := sha256.New()
    	h.Write([]byte(s))
    	bs := h.Sum(nil)
	pName:=string(bs)
	datalist1:="True name:"+args0
	datalist2:="False name:"+pName
	datalist3:="Country:"+Property[0].x
	datalist4:="Organization:"+Property[1].x
	datalist5:="Locality:"+Property[2].x
	datalist6:="StreetAddress:"+Property[3].x
	datalist7:="PostalCode"+Property[4].x
	dataList := [][]byte{
		[]byte(datalist1),
		[]byte(datalist2),
		[]byte(datalist3),
		[]byte(datalist4),
		[]byte(datalist5),
		[]byte(datalist6),
		[]byte(datalist7),
	}
	rand.Seed(time.Now().UnixNano())
	seed := strconv.Itoa(rand.Int())
	fmt.Println("random seed：", seed)
	root := GenerateZKMerkleRoot(dataList, []byte(seed))
	k2:=IntToBytes(k)
	k3 := uint64(binary.BigEndian.Uint64(k2))
	//////////////////////////////////////////////////////////////////////////////
	//Generate Merkle Evidence
	//////////////////////////////////////////////////////////////////////////////
	evidence, _ := GenerateZKMerkleEvidence(dataList, []byte(seed), k3)
	elapsed = time.Since(t3)
	GenEvidencetime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	t4 := time.Now()
	combinebytes:=marshal(evidence,root)
	fmt.Println("\nthe response1.Payload is ",response1.Payload)
	ComBytes:=bytes.Split(response1.Payload, []byte(")))))"))
	_ = stub.PutPrivateData(collection1,"prime",ComBytes[1])
	fmt.Println("\nthe combinebytes is ",combinebytes)
	combineBytes:=bytes.Join([][]byte{ComBytes[0],combinebytes}, []byte(":::::"))
	fmt.Println("\nthe combineBytes is ",combineBytes)
	EscrowedKey:= fmt.Sprintf("%X", pName)
	fmt.Println("key ID：",EscrowedKey)
	//////////////////////////////////////////////////////////////////////////////
	//put combinebytes on chain
	//////////////////////////////////////////////////////////////////////////////
	for i:=0;i<l1;i++{
		err := stub.PutPrivateData(Collection[i].x,EscrowedKey,combineBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		_= stub.PutPrivateData(Collection[i].x,"AnonymitySetup",[]byte(EscrowedKey))
	}
	elapsed = time.Since(t4)
	putonchaintime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	fmt.Println("combineBytes:",len(combineBytes))
	fmt.Println("evidence:",len(combinebytes))
	elapsed = time.Since(t0)
	alltime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	return shim.Success([]byte("The name Anonymization is complete! All time :"+alltime+"Get Secret time:"+getSecrettime+"Gen Evidence time:"+GenEvidencetime+"put on chain time"+putonchaintime))
}

//////////////////////////////////////////////////////////////////////////////
//Read the pseudonym pName from AS key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Anony_Setup(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	collection:=args[0]
	result, err := stub.GetPrivateData(collection,"AnonymitySetup")
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("\nthe result is ",result)
	return shim.Success(result)
}


//////////////////////////////////////////////////////////////////////////////
//Read the anonymized escrowed session key from the EK key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Anonymized_EK(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	keyid := args[0]
	collection:=args[1]
	result, err := stub.GetPrivateData(collection,keyid)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("\nthe result is ",result)
	return shim.Success(result)
}

//////////////////////////////////////////////////////////////////////////////
//Read key names of the public keys and Private Data
//Collection names from GS key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Global_Setup(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	result, err := stub.GetState("GlobalSetup")
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("\nthe result is ",result)
	return shim.Success(result)
}

//////////////////////////////////////////////////////////////////////////////
//Read the supervisor’s public key from public ledger
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Super_PubKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

//////////////////////////////////////////////////////////////////////////////
//Recover the final session key SessKey, and store
//SessKey under the FK key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Recover_Final_SessKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	t0 := time.Now()
	pcollection:=args[0]
	response2,_:= stub.GetState("GlobalSetup")
	enConBytes1:=bytes.Split(response2, []byte(";;;;;"))
	kstring:=string(enConBytes1[0])
	k,_:=strconv.Atoi(kstring)
	queryArgs := [][]byte{[]byte("Get_Record_Name"),[]byte(pcollection)}
	response3 := stub.InvokeChaincode("DataSource", queryArgs, "mychannel")
	if response3.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response3.Payload))
		}
	senderKeyid:=string(response3.Payload)
	kemName:=string(enConBytes1[2])
	FinalKey:=senderKeyid
	Ccollection:=string(enConBytes1[9])
	namecollection:=string(enConBytes1[10])
	esagentKeyid,_:= stub.GetPrivateData(namecollection,"AnonymitySetup")
	EsxyKeyid:=make([]esPKKeyid,k)
	Collection:=make([]collection,k)
	//XyKeyid:=make([]xyKeyid,k)
	Point:= make([]Data,k)
	//ID:=make([]id,k)
	for i:=0;i<k;i++{
	EsxyKeyid[i].x=string(esagentKeyid)
	Collection[i].x=string(enConBytes1[i+15])
	}
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
	//Get the X and Y from private collection
	//////////////////////////////////////////////////////////////////////////////
	t2:= time.Now()
	for i:=0;i<k;i++{
	queryArgs := [][]byte{[]byte("Get_Partial_SessKey"),[]byte(EsxyKeyid[i].x),[]byte(Collection[i].x)}
	response2 := stub.InvokeChaincode("EscrowAgent", queryArgs, "mychannel")
	if response2.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response2.Payload))
		}
	x:=big.NewInt(1)
	y:=big.NewInt(1)
	enConBytes:=bytes.Split(response2.Payload, []byte("-----"))
	fmt.Println("response2.Payload=",enConBytes[0])
	x=x.SetBytes(enConBytes[0])
	fmt.Println("x=",x)
	y=y.SetBytes(enConBytes[1])
	fmt.Println("y=",y)
	Point[i].x=x
	Point[i].y=y
	}
	elapsed2 := time.Since(t2)
	getxytime := strconv.FormatFloat(elapsed2.Seconds(), 'E', -1, 64)
	//////////////////////////////////////////////////////////////////////////////
	//Get the prime number from chain
	//////////////////////////////////////////////////////////////////////////////
	t3:= time.Now()
	result ,_:= stub.GetPrivateData(pcollection,"prime")
	r:=big.NewInt(1)
	r=r.SetBytes(result)
	fmt.Println("r=",r)
	elapsed3 := time.Since(t3)
	getprimetime := strconv.FormatFloat(elapsed3.Seconds(), 'E', -1, 64)
	//////////////////////////////////////////////////////////////////////////////
	//Use the LagrangeInsert to recover the SessKey_A
	//////////////////////////////////////////////////////////////////////////////
	t1 := time.Now()
	Secret:=LagrangeInsert(Point,k,r)
	fmt.Println("Secret=",Secret)
	SessKey_A:=Secret.Bytes()
	elapsed1 := time.Since(t1)
	combinetime := strconv.FormatFloat(elapsed1.Seconds(), 'E', -1, 64)
	t6 := time.Now()
	queryArgs3 := [][]byte{[]byte("Get_Source_Data"), []byte(senderKeyid),[]byte(pcollection)}
	response := stub.InvokeChaincode("DataSource", queryArgs3, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
		}
	enConBytescom:=bytes.Split(response.Payload, []byte(")))))"))
	enConBytesSup:=bytes.Split(enConBytescom[0], []byte("-----"))
	elapsed6 := time.Since(t6)
	Get_xy_Valuetime := strconv.FormatFloat(elapsed6.Seconds(), 'E', -1, 64)
	t5 := time.Now()
	//////////////////////////////////////////////////////////////////////////////
	//Deacap  to get the SessKey_A
	//////////////////////////////////////////////////////////////////////////////
	SessKey_B, _ := DownloadDecapSessionKey(enConBytesSup[1], privateKeyBytes, kemName)
	elapsed5 := time.Since(t5)
	decapSecretTime := strconv.FormatFloat(elapsed5.Seconds(), 'E', -1, 64)
	t4 := time.Now()
	if len(SessKey_B) >= 32 {
		hash := sha256.New()
		hash.Write(SessKey_B)
		SessKey_B = hash.Sum(nil)
	}
	SessKey := make([]byte, len(SessKey_B))
	for i:=0; i<len(SessKey_A); i++ {
		SessKey[i] = SessKey_A[i] ^ SessKey_B[i]
	}
	if len(SessKey) >= 32 {
		hash := sha256.New()
		hash.Write(SessKey)
		SessKey = hash.Sum(nil)
	}
	err = stub.PutPrivateData(Ccollection,FinalKey, SessKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SessKey:",len(SessKey))
	elapsed4 := time.Since(t4)
	putonchainTime := strconv.FormatFloat(elapsed4.Seconds(), 'E', -1, 64)
	return shim.Success([]byte("Init parameter:"+alltime+"recover xy time:"+getxytime+"get prime time:"+getprimetime+"Get court pubkey from chain time:"+Get_xy_Valuetime+"decap secret time:"+decapSecretTime+"combine secret time:"+combinetime+"put on chain Time:"+putonchainTime))
}

//////////////////////////////////////////////////////////////////////////////
//Read the final session key from the FK key
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Get_Final_SessKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	keyid := args[0]
	collection:=args[1]
	result, err := stub.GetPrivateData(collection,keyid)
	fmt.Println("result=",result)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}

//internal function:under 3 function are used to do the LagrangeInsert
func LagrangeInsert(Key []Data,k int,r *big.Int)*big.Int{
	s:=big.NewInt(1)
	secret:=big.NewInt(0)
	for j:=0;j<k;j++{
		s=basis(Key,k,j,r)
		d1 :=s.Sign()	
		if d1<0{
		s=s.Add(s,r)
		}
		
	secret=secret.Mod(secret.Add(secret,s.Mul(s,Key[j].y)),r)
	
	}
	fmt.Println("the final secert is:",secret.Mod(secret,r))
	return secret
}


func basis(Key []Data,k int,j int,p *big.Int)*big.Int{
	terms:=make([]*big.Int,k-1)
	i:=0
	for m:=0;m<k;m++{
		if m!=j{
		o:=big.NewInt(1)
		o.Sub(Key[j].x,Key[m].x)
		u :=big.NewInt(1)
		u.ModInverse(o,p)
		l:=big.NewInt(1)
		terms[i]=l.Mul(l.Sub(big.NewInt(0),Key[m].x),u)
		i++
		}
	}
	x:=big.NewInt(1)
	y:=x.Mod(prod(terms,k),p)
	return y
}

func prod(terms []*big.Int,k int)*big.Int{
	product:=big.NewInt(1)
	for i:=0;i<k-1;i++{
	product=product.Mul(product,terms[i])
	}
	return product
}

//////////////////////////////////////////////////////////////////////////////
//Generate one post-quantum public/private key pair
//for the supervisor, upload the public key to public
//ledger, and store the private key off-chain
//////////////////////////////////////////////////////////////////////////////
func (t *Sup) Gen_Super_KeyPair(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	

	// verify identity of the escrower 
	combineBytes, err := stub.GetState("GlobalSetup")
	enConBytes:=bytes.Split(combineBytes, []byte(";;;;;"))
	// We use generateKeyPair function to generate publicKey and privateKey about escrower.
	kemName := string(enConBytes[2]) 
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
	EsPKKeyid:=string(enConBytes[3])
	err = stub.PutState(EsPKKeyid, publicKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("Put public key successfully!!!------privateKey:" + strPriKey))
}

//encap to get CT and sessionkey
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

//decap to get sessionkey
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



func main() {
	err1 := shim.Start(new(Sup))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}












	


