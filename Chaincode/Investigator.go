package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strconv"
	"crypto/x509"
	"encoding/pem"
	"time"
	//"encoding/base64"
	_"math/big"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	_"crypto/sha256"
	//"github.com/open-quantum-safe/liboqs-go/oqs"
)

type Recover struct{}

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


func (t *Recover) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Success invoke and not opter!!"))
}

func (t *Recover) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	 if fn == "Dec_Secrect_Data" {
		return t.Dec_Secrect_Data(stub, args)
	}
	return shim.Error("Recevied unkown function invocation")
}




//////////////////////////////////////////////////////////////////////////////
//Decrypt the encrypted secret data C read from SD4I key bound to Data Source chaincode by using
//the final session key read from FK key bound to Supervisor chaincode
//////////////////////////////////////////////////////////////////////////////
func (t *Recover) Dec_Secrect_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	t0 := time.Now()
	colleciton2:=args[0]
	queryArgs := [][]byte{[]byte("Get_Global_Setup")}
	response2 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response2.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response2.Payload))
	}
	enConBytes1:=bytes.Split(response2.Payload, []byte(";;;;;"))
	queryArgs = [][]byte{[]byte("Get_Record_Name"),[]byte(colleciton2)}
	response3 := stub.InvokeChaincode("DataSource", queryArgs, "mychannel")
	if response3.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response3.Payload))
	}
	senderKeyid:=string(response3.Payload)
	AESKeyid := senderKeyid
	colleciton1:=string(enConBytes1[9])
	//////////////////////////////////////////////////////////////////////////////
	//Get the Sender Data from chain
	//////////////////////////////////////////////////////////////////////////////
	queryArgs = [][]byte{[]byte("Get_Source_Data"), []byte(senderKeyid), []byte(colleciton2)}
	response := stub.InvokeChaincode("DataSource", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	enConBytes := response.Payload
	elapsed := time.Since(t0)
	Get_Sender_Datatime := strconv.FormatFloat(elapsed.Seconds(), 'E', -1, 64)
	t2 := time.Now()
	queryArgs = [][]byte{[]byte("Get_Final_SessKey"),[]byte(AESKeyid),[]byte(colleciton1)}
	response1 := stub.InvokeChaincode("Supervisor", queryArgs, "mychannel")
	if response1.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response1.Payload))
		}
	combineSecret:=response1.Payload
	elapsed2 := time.Since(t2)
	Get_AES_Keytime := strconv.FormatFloat(elapsed2.Seconds(), 'E', -1, 64)
	t1 := time.Now()
	messageBytes, err:= AES_Decrypt(enConBytes, combineSecret)
	if err != nil {
		return shim.Error(err.Error())
	}
	elapsed1 := time.Since(t1)
	Decrypttime := strconv.FormatFloat(elapsed1.Seconds(), 'E', -1, 64)
	//fmt.Println(string(messageBytes))
	return shim.Success([]byte("The listener successfully receive the message!!!---MESSAGE:" +  string(messageBytes) +"Decrypt message time :"+Decrypttime+"Get Sender Data time:"+Get_Sender_Datatime+"Get AES Keytime:"+Get_AES_Keytime ))
}


func main() {
	err1 := shim.Start(new(Recover))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}












	


