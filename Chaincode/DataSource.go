
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"crypto/sha256"
	"encoding/binary"
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"strconv"
	crypt_rand "crypto/rand"
	"math/big"
	"log"
)

type Sender struct{}


func AES_Encrypt(privData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	blockSize := block.BlockSize()
	privData = PKCS7Padding(privData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encryptedDate := make([]byte, len(privData))
	blockMode.CryptBlocks(encryptedDate, privData)
	return encryptedDate, nil
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (t *Sender) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Success invoke and not opter!!"))
}

func (t *Sender) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "Gen_Source_Data" {
		return t.Gen_Source_Data(stub, args)
	} else if fn == "Get_Source_Data" {
		return t.Get_Source_Data(stub, args)
	} else if fn=="Get_Pub_Key" {
		return t.Get_Pub_Key(stub, args)
	} else if fn=="Get_Record_Name"{
		return t.Get_Record_Name(stub, args)
	}
	return shim.Error("Recevied unkown function invocation")
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

//////////////////////////////////////////////////////////////////////////////
//Read all the public keys of the supervisor and escrow
//agents based on GS key (bound to Supervisor chaincode)
//////////////////////////////////////////////////////////////////////////////
func (t *Sender) Get_Pub_Key(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	queryArgs := [][]byte{[]byte("Get_Global_Setup")}
	response := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	enConBytes:=bytes.Split(response.Payload, []byte(";;;;;"))
	n1:=string(enConBytes[1])
	n,_:=strconv.Atoi(n1)
	fmt.Println("n:",n)
	for i:=0;i<n;i++{
	keyid:=string(enConBytes[i+4])
	fmt.Println("keyid:",keyid)
	queryArgs := [][]byte{[]byte("Get_EA_PubKey"), []byte(keyid)}
	response := stub.InvokeChaincode("EscrowAgent", queryArgs, "mychannel")
	pubKey:=response.Payload
		if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
		}
	fmt.Println("response.Payload:",pubKey)
	err := stub.PutState(keyid, pubKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	}
	return shim.Success([]byte("Get pubkey successfully!!!"))
}

func powerf(x *big.Int,n int)*big.Int{
	ans:=big.NewInt(1)
	for n!=0{
	ans=ans.Mul(ans,x)
	n--
	}
	return ans
}
type Data struct{
	  x *big.Int;
	  y *big.Int
}
type CT struct{
	  x []byte;
	 
}
type SS struct{
	  x []byte;
	 
}
type Response struct{
	  x []byte;
	 
}

type esPKKeyid struct{
	  x string;
	 
}

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

//////////////////////////////////////////////////////////////////////////////
//Retrieve the secret data M from off-chain private
//database, generate and upload source data (i.e., store
//the encrypted M and escrowed session key separately
//under SD4S and SD4I keys)
//////////////////////////////////////////////////////////////////////////////
func (t *Sender) Gen_Source_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	t0 := time.Now()
	//////////////////////////////////////////////////////////////////////////////
	//Get the private message and KS through the getTransient function of shim
	// and the type of message and KS is []byte.
	//////////////////////////////////////////////////////////////////////////////
	tMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error(fmt.Sprintf("Could not retrieve transient, err %s", err))
	}
	messageBytes, ok := tMap["MESSAGE"]
	if !ok {
		return shim.Error(fmt.Sprintf("Expected transient message"))
	}
	queryArgs := [][]byte{[]byte("Get_Global_Setup")}
	response2 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response2.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response2.Payload))
	}
	enConBytes:=bytes.Split(response2.Payload, []byte(";;;;;"))
	elapsed := time.Since(t0)
	alltime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	ttime := time.Now()
	queryArgs = [][]byte{[]byte("Get_Global_Time")}
	response1 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response1.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response1.Payload))
	}
	Time:=string(response1.Payload)
	elapsedtime := time.Since(ttime)
	gettime := strconv.FormatFloat(elapsedtime.Seconds(), 'E', -1, 64)
	userKeyid := args[0]
	senderKeyid := userKeyid+Time
	kemName := string(enConBytes[2]) 
	collection1:=args[1]
	collection2:=args[2]
	lstring:=string(enConBytes[1])
	kstring:=string(enConBytes[0])
	l,_:=strconv.Atoi(lstring)
	k,_:=strconv.Atoi(kstring)
	EsPKKeyid:= make([]esPKKeyid,l)
	response:= make([]Response,l)
	SSecret:=make([]SS,l)
	CT :=make([]CT,l)
	server := oqs.KeyEncapsulation{}
	defer server.Clean() 
	if err := server.Init(kemName, nil); err != nil {
		return shim.Error(err.Error())
	}
	for i:=0;i<l;i++{
	EsPKKeyid[i].x = string(enConBytes[i+4])
	}
		
	//////////////////////////////////////////////////////////////////////////////
	//Get the publicKey from Chain
	//////////////////////////////////////////////////////////////////////////////

	for i:=0;i<l;i++{
	response[i].x, _ = stub.GetState(EsPKKeyid[i].x)
	//////////////////////////////////////////////////////////////////////////////
	//Use EncapSecret function to generate CT[i],SS[i]
	//////////////////////////////////////////////////////////////////////////////
	CT[i].x, SSecret[i].x,  err = GenEncapUploadSessionKey(kemName,response[i].x)
	if err != nil {
		fmt.Println("\nGenEncapUploadSessionKey fails: ", err)
	}
		if len(SSecret[i].x) >= 32 {
		hash := sha256.New()
		hash.Write(SSecret[i].x)
		SSecret[i].x = hash.Sum(nil)
		}
	}
	Suppubkey := string(enConBytes[3])
	queryArgs = [][]byte{[]byte("Get_Super_PubKey"),[]byte(Suppubkey)}
	response3 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response3.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response3.Payload))
	}
	t1:=time.Now()
	countCT, SessKey_B, errs := GenEncapUploadSessionKey(kemName,response3.Payload)
	elapsed1 := time.Since(t1)
	encapSecretTime := strconv.FormatFloat(elapsed1.Seconds(), 'E', -1, 64)	
	if errs != nil {
		fmt.Println("\nGenEncapUploadSessionKey fails: ", errs)
	}
	fmt.Println("-------EncapSecretTime:"+encapSecretTime)
	t2 := time.Now()
	if len(SessKey_B) >= 32 {
		hash := sha256.New()
		hash.Write(SessKey_B)
		SessKey_B = hash.Sum(nil)
	}
	SessKey_A := make([]byte, len(SessKey_B))
	fmt.Println("len(SessKey_B):",len(SessKey_B))
   	 _, err = crypt_rand.Read(SessKey_A)
    	// Note that err == nil only if we read len(b) bytes.
    	if err != nil {
        fmt.Println("error",err)
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
	//////////////////////////////////////////////////////////////////////////////
	//Use SK to encrypt message through the AES function.
	//////////////////////////////////////////////////////////////////////////////
	encryptedMessage, err:= AES_Encrypt(messageBytes, SessKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	elapsed2 := time.Since(t2)
	EncryptSecretTime := strconv.FormatFloat(elapsed2.Seconds(), 'E', -1, 64)
	fmt.Println("Encrypt messagetime"+EncryptSecretTime)
	t3 := time.Now()
	max := big.NewInt(1)
	max=max.SetBytes(SessKey_A)
	fmt.Println("max=",max)
	b:=0
	//////////////////////////////////////////////////////////////////////////////
	//Generate the prime Number.
	//////////////////////////////////////////////////////////////////////////////

	te:=1
	r:=big.NewInt(1)
	for te>0{
	r,_ = crypt_rand.Prime(crypt_rand.Reader,256) 
	if r.ProbablyPrime(20) {
		fmt.Print("%d is probably prime\n", r)
					}else{
	break
					}
        
	te=max.Cmp(r)
	}
    	
	fmt.Println("r=",r)
	All:= make([]Data,l)
	coe:=make([]*big.Int,k)
	//////////////////////////////////////////////////////////////////////////////
	//Random generation coefficient
	//////////////////////////////////////////////////////////////////////////////
	for q:=0;q<k-1;q++{
	coeNumber, _ := crypt_rand.Int(crypt_rand.Reader, max)
	coe[q+1]=coeNumber
	fmt.Println("coe[q+1]=",coe[q+1],"q+1=",q+1)
	}
	//////////////////////////////////////////////////////////////////////////////
	//Use the SS[i] as the value of X and then generate the Y
	//////////////////////////////////////////////////////////////////////////////
	for j:=0;j<l;j++ {
	x:=big.NewInt(1)
	x=x.SetBytes(SSecret[j].x)
	secret:=big.NewInt(1)
	secret=secret.SetBytes(SessKey_A)
		for z:=1;z<k;z++{
		a:=coe[z]
		pe:=powerf(x,z)
		a2:=big.NewInt(1)
		a2.Mul(a,pe)
		secret.Add(secret,a2)
		}
	
	fmt.Println("x=",x,"y=",secret.Mod(secret,r))
	All[b].x=x
	All[b].y=secret
	b++
	}
	elapsed3 := time.Since(t3)
	deviceSecrettime := strconv.FormatFloat(elapsed3.Seconds(), 'E', -1, 64)
	fmt.Println("device Secret time"+deviceSecrettime)
	t4 := time.Now()
	v:=r.Bytes()
	ycombineBytes:=All[0].y.Bytes()
	ycombineBytes = bytes.Join([][]byte{CT[0].x,ycombineBytes}, []byte(";;;;;"))
	for i:=1;i<l;i++{
	o:=All[i].y.Bytes()
	ycombineBytes = bytes.Join([][]byte{ycombineBytes,CT[i].x,o}, []byte(";;;;;"))
	}
	SD4S:=senderKeyid
	SD4I:=senderKeyid
	err = stub.PutPrivateData(collection2, SD4I, encryptedMessage)
	_ = stub.PutPrivateData(collection2, "PersonalSetup", []byte(SD4I))
	if err != nil {
		return shim.Error(err.Error())
	}
	//////////////////////////////////////////////////////////////////////////////
	//Put the combine messages and CT on chain
	//////////////////////////////////////////////////////////////////////////////
	ycombineBytes = bytes.Join([][]byte{ycombineBytes,countCT}, []byte("-----"))
	combineBytes:=bytes.Join([][]byte{ycombineBytes,v}, []byte(")))))"))
	err = stub.PutPrivateData(collection1,SD4S, combineBytes)
	_ = stub.PutPrivateData(collection1, "PersonalSetup", []byte(SD4S))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("[]byte(senderKeyid):",len([]byte(senderKeyid)))
	fmt.Println("encryptedMessage:",len(encryptedMessage))
	fmt.Println("y:",len(All[0].y.Bytes()))
	fmt.Println("CT:",len(CT[0].x))
	fmt.Println("v:",len(v))
	fmt.Println("combineBytes:",len(combineBytes))
	elapsed4 := time.Since(t4)
	chaintime := strconv.FormatFloat(elapsed4.Seconds(), 'E', -1, 64)
	return shim.Success([]byte("Init parameter:"+ alltime+"get time:"+gettime+"Encap Secret Time:"+encapSecretTime+"Encrypt message time:"+EncryptSecretTime+"device Secret time"+deviceSecrettime+"put on chain"+chaintime))
}

//////////////////////////////////////////////////////////////////////////////
//Read the source data (i.e., the encrypted M and escrowed session key) from SD4S and SD4I keys
//////////////////////////////////////////////////////////////////////////////
func (t *Sender) Get_Source_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	keyid := args[0]
	collection:=args[1]
	result, err := stub.GetPrivateData(collection,keyid)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}

//////////////////////////////////////////////////////////////////////////////
//Read the source data record name (i.e., the concatenation of data source ID id and the time interval t
//during which the record is generated) from RN key
//////////////////////////////////////////////////////////////////////////////
func (t *Sender) Get_Record_Name(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	keyid := "PersonalSetup"
	collection:=args[0]
	result, err := stub.GetPrivateData(collection,keyid)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}




func main() {
	err1 := shim.Start(new(Sender))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}
