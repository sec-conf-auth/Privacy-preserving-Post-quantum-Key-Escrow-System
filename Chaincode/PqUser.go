package main

import (
	"fmt"
	_"log"
	_"bytes"
	_"io"
	//"encoding/base64"
	_"crypto/x509"
	_"crypto/ecdsa"
	_"crypto/rsa"
	_"crypto/sha256"
	_"math/big"
	_"crypto/rand"
	_"crypto/elliptic"
	"time"
	_"encoding/pem"
	_"encoding/gob"
   	"math/big"
    	"crypto/x509"
	"crypto/x509/pkix"
	"bytes"
	"encoding/pem"
	"crypto/sha256"
	"encoding/gob"
	"encoding/base64" 
	"crypto/rand"
//	"encoding/hex"
//	"encoding/json"
//	"encoding/asn1"
	"log"
/*	
	"net"
	"runtime"
*/
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
//	"ecies"
)
var layout string
type Cert struct {
	version 			string
	serialNumber			*big.Int
	issuer 				pkix.Name
	subject 			pkix.Name
	notBefore,notAfter 		time.Time
	keyUsage 			x509.KeyUsage
	sigName_pq			string
	sig_pq				[]byte


	algName_pq			string
	privateKey_pq_PEM 		[]byte
	pubKey_pq			[]byte
	parentKeySerialNum		string
	extensions			[]pkix.Extension

	isKeyEncap			bool
	isID				bool
	isCA				bool
	isRevoked 			bool

	// Additional info provided when cert was created
	// additional 		[]byte
}
type Cert_NoPrivKey struct {
	version 			string
	serialNumber			*big.Int
	issuer 				pkix.Name
	subject 			pkix.Name
	notBefore,notAfter 		time.Time
	keyUsage 			x509.KeyUsage
	sigName_pq			string
	sig_pq				[]byte


	algName_pq			string
	pubKey_pq			[]byte
	parentKeySerialNum		string
	extensions			[]pkix.Extension

	isKeyEncap			bool
	isID  				bool
	isCA      			bool
	isRevoked 			bool

	// Additional info provided when cert was created
	// additional 		[]byte
}
type certgob struct {
	Version 			string
	SerialNumber			string
	Country				[]string
	Organization  			[]string
	OrganizationalUnit		[]string
	Locality 			[]string
	Province 			[]string
	StreetAddress 			[]string
	PostalCode 			[]string
	CommonName 			string
	NotBefore,NotAfter 		string
	KeyUsage 			x509.KeyUsage
	SigName_pq			string	
	Sig_pq				[]byte


	AlgName_pq			string
	PrivateKey_pq_PEM 		[]byte
	PubKey_pq			[]byte
	ParentKeySerialNum		string
	Extensions			[]pkix.Extension

	IsKeyEncap			bool
	IsIndividual			bool
	IsCA				bool
	IsRevoked 			bool

	// Additional info provided when cert was created
	// Additional 		[]byte
}
type certgob_NoPrivKey struct {
	Version 			string
	SerialNumber			string
	Country				[]string
	Organization 			[]string
	OrganizationalUnit		[]string
	Locality 			[]string
	Province 			[]string
	StreetAddress 			[]string
	PostalCode 			[]string
	CommonName 			string
	NotBefore,NotAfter 		string
	KeyUsage 			x509.KeyUsage
	SigName_pq			string	
	Sig_pq				[]byte


	AlgName_pq			string
	PubKey_pq			[]byte
	ParentKeySerialNum		string
	Extensions			[]pkix.Extension

	IsKeyEncap			bool
	IsIndividual			bool
	IsCA				bool
	IsRevoked 			bool

	// Additional info provided when cert was created
	// Additional 		[]byte
}

// hash publickey; we use it as a salt for encryption and also SubjectKeyId
func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}

func randSerial() *big.Int {
	min := big.NewInt(1)
	min.Lsh(min, 120)
	max := big.NewInt(1)
	max.Lsh(max, 130)
	for {
		serial, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(fmt.Errorf("ca: can't generate serial#: %w", err))
		}

		if serial.Cmp(min) > 0 {
			return serial
		}
	}
	panic("can't gen new CA serial")
}

func marshalCert(crt *Cert)([]byte, error){
	sn := crt.subject.CommonName
	NotBeforestring:=crt.notBefore.Format(layout)
	NotAfterstring:=crt.notAfter.Format(layout)
	SerialNumberstring:=crt.serialNumber.String()
	cg := &certgob{
		Version: 			crt.version,
		SerialNumber:			SerialNumberstring,
		Country: 			crt.subject.Country,
		Organization: 			crt.subject.Organization,
		OrganizationalUnit:		crt.subject.OrganizationalUnit,
		Locality:			crt.subject.Locality,
		Province:			crt.subject.Province,
		StreetAddress:			crt.subject.Province,
		PostalCode:			crt.subject.PostalCode,
		CommonName:			crt.subject.CommonName,
		NotBefore:			NotBeforestring,
		NotAfter: 			NotAfterstring,		
		KeyUsage: 			crt.keyUsage,
		SigName_pq:			crt.sigName_pq,
		Sig_pq:				crt.sig_pq,

		AlgName_pq:			crt.algName_pq,
		PrivateKey_pq_PEM: 		crt.privateKey_pq_PEM,
		PubKey_pq:			crt.pubKey_pq,
		ParentKeySerialNum:		crt.parentKeySerialNum,

		IsKeyEncap:			crt.isKeyEncap,
		IsIndividual:			crt.isID,
		IsCA:				crt.isCA,
		IsRevoked:			crt.isRevoked,
	}

	var b bytes.Buffer
	g := gob.NewEncoder(&b)
	err := g.Encode(cg)
	if err != nil {
		return nil, fmt.Errorf("%s: can't gob-encode cert: %s", sn, err)
	}
	return b.Bytes(), nil
}

func marshalCert_NoPrivKey(crt *Cert)([]byte, error){
	sn := crt.subject.CommonName
	NotBeforestring:=crt.notBefore.Format(layout)
	NotAfterstring:=crt.notAfter.Format(layout)
	SerialNumberstring:=crt.serialNumber.String()
	cg := &certgob_NoPrivKey{
		Version: 			crt.version,
		SerialNumber:			SerialNumberstring,
		Country: 			crt.subject.Country,
		Organization: 			crt.subject.Organization,
		OrganizationalUnit:		crt.subject.OrganizationalUnit,
		Locality:			crt.subject.Locality,
		Province:			crt.subject.Province,
		StreetAddress:			crt.subject.Province,
		PostalCode:			crt.subject.PostalCode,
		CommonName:			crt.subject.CommonName,
		NotBefore:			NotBeforestring,
		NotAfter: 			NotAfterstring,		
		KeyUsage: 			crt.keyUsage,
		SigName_pq:			crt.sigName_pq,
		Sig_pq:				crt.sig_pq,

		AlgName_pq:			crt.algName_pq,
		PubKey_pq:			crt.pubKey_pq,
		ParentKeySerialNum:		crt.parentKeySerialNum,

		IsKeyEncap:			crt.isKeyEncap,
		IsIndividual:			crt.isID,
		IsCA:				crt.isCA,
		IsRevoked:			crt.isRevoked,
	}
	var b bytes.Buffer
	g := gob.NewEncoder(&b)
	err := g.Encode(cg)
	if err != nil {
		return nil, fmt.Errorf("%s: can't gob-encode cert: %s", sn, err)
	}
	return b.Bytes(), nil
}

func unmarshalCert(crtBytes []byte)(*Cert, error){
	var cg certgob
	
	b := bytes.NewBuffer(crtBytes)
	g := gob.NewDecoder(b)
	err := g.Decode(&cg)
	if err != nil {
		return nil, fmt.Errorf("g.Decode: can't decode gob: %s", err)
	}

	sub := &pkix.Name{
		Country: 			cg.Country,
		Organization: 			cg.Organization,
		OrganizationalUnit: 		cg.OrganizationalUnit,
		Locality: 			cg.Locality,
		Province: 			cg.Province,
		StreetAddress: 			cg.StreetAddress,
		PostalCode: 			cg.PostalCode,
		// SerialNumber: 		        "CN",
		CommonName: 			cg.CommonName,
	}
	NotBefore,_:=time.Parse(layout,cg.NotBefore)
	NotAfter,_:=time.Parse(layout,cg.NotAfter)
	SerialNumber,_:=new(big.Int).SetString(cg.SerialNumber,10)
	crt := &Cert{
		version: 			cg.Version,
		serialNumber:			SerialNumber,
		issuer: 			*sub,
		subject: 			*sub,
		notBefore:			NotBefore,
		notAfter: 			NotAfter,
		keyUsage: 			cg.KeyUsage,
		sigName_pq:			cg.SigName_pq,
		sig_pq:				cg.Sig_pq,

		algName_pq:			cg.AlgName_pq,
		privateKey_pq_PEM: 		cg.PrivateKey_pq_PEM,
		pubKey_pq:			cg.PubKey_pq,
		parentKeySerialNum:		cg.ParentKeySerialNum,

		isKeyEncap:			cg.IsKeyEncap,
		isID:				cg.IsIndividual,
		isCA:				cg.IsCA,
		isRevoked:			cg.IsRevoked,
	}
	return crt, nil
}

func unmarshalCert_NoPrivKey(crtBytes []byte)(*Cert_NoPrivKey, error){
	var cg certgob
	b := bytes.NewBuffer(crtBytes)
	g := gob.NewDecoder(b)
	err := g.Decode(&cg)
	if err != nil {
		return nil, fmt.Errorf("g.Decode: can't decode gob: %s", err)
	}
	sub := &pkix.Name{
		Country: 			cg.Country,
		Organization: 			cg.Organization,
		OrganizationalUnit: 		cg.OrganizationalUnit,
		Locality: 			cg.Locality,
		Province: 			cg.Province,
		StreetAddress: 			cg.StreetAddress,
		PostalCode: 			cg.PostalCode,
		// SerialNumber: 		        "CN",
		CommonName: 			cg.CommonName,
	}
	NotBefore,_:=time.Parse(layout,cg.NotBefore)
	NotAfter,_:=time.Parse(layout,cg.NotAfter)
	SerialNumber,_:=new(big.Int).SetString(cg.SerialNumber,10)
	crt := &Cert_NoPrivKey{
		version: 			cg.Version,
		serialNumber:			SerialNumber,
		issuer: 			*sub,
		subject: 			*sub,
		notBefore:			NotBefore,
		notAfter: 			NotAfter,
		keyUsage: 			cg.KeyUsage,
		sigName_pq:			cg.SigName_pq,
		sig_pq:				cg.Sig_pq,

		algName_pq:			cg.AlgName_pq,
		pubKey_pq:			cg.PubKey_pq,
		parentKeySerialNum:		cg.ParentKeySerialNum,

		isKeyEncap:			cg.IsKeyEncap,
		isID:				cg.IsIndividual,
		isCA:				cg.IsCA,
		isRevoked:			cg.IsRevoked,
	}
	return crt, nil
}

func GenEncapUploadSessionKey(certKE *Cert_NoPrivKey)([]byte,  []byte, error){
	kemName := certKE.algName_pq
	kemer := oqs.KeyEncapsulation{}
	defer kemer.Clean() // clean up even in case of panic
	if err := kemer.Init(kemName, nil); err != nil {
		log.Fatal(err)
		return nil,  nil, err
	}
	/////////////////////////////////////////////////////////////////////
	// sessionkey is used to be encapsulated by PQ key Encapsulatioon
	/////////////////////////////////////////////////////////////////////
	fmt.Println("\nThe details of the KEM algorithm is:\n", kemer.Details().String())

	encapSessionKey, sessionKey, err := kemer.EncapSecret(certKE.pubKey_pq)
	if err != nil {
		log.Fatal(err)
		return nil,  nil, err
	}
	return encapSessionKey, sessionKey, nil
}

func DownloadDecapSessionKey(encapSessionKey []byte,  certKE *Cert, pwKE string)([]byte, error){

	/////////////////////////////////////////////////////////////////////
	// Decode the PQ private key from PEM in the given certKE
	//in order to decapsulate sessionKey
	/////////////////////////////////////////////////////////////////////
	passKE := []byte(pwKE)
	privateKey_pq_PEMBlock_KE, _ := pem.Decode(certKE.privateKey_pq_PEM)
	privateKey_pq_KE, _ := x509.DecryptPEMBlock(privateKey_pq_PEMBlock_KE, passKE)
	kemer := oqs.KeyEncapsulation{}
	defer kemer.Clean() // clean up even in case of panic

	if err := kemer.Init(certKE.algName_pq, privateKey_pq_KE); err != nil {
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

type user struct{}

func (t *user) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("begin to init chaincode....")
   return shim.Success([]byte("Success invoke and not opter!!"))
}

func (t *user) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "Verify_Cert" {
		return t.Verify_Cert(stub, args)
	}else if fn == "Encap" {
		return t.Encap(stub, args)
	}else if fn == "Decap" {
		return t.Decap(stub, args)
	}else if fn == "Query_Cert" {
		return t.Query_Cert(stub, args)
	}
	return shim.Error("Recevied unkown function invocation")
}

	/////////////////////////////////////////////////////////////////////
	// Invoke the chaincode PqCa and verify one
	// certificate (together with the related
	// cerificate chain)
	/////////////////////////////////////////////////////////////////////
func (t *user) Verify_Cert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	Serialstring2 := args[0]
	str2 := []byte(Serialstring2)
	queryArgs := [][]byte{[]byte("Verify_Cert"), str2}
	/////////////////////////////////////////////////////////////////////
	// Invoke the function Verify_Cert of the chaincode PqCa
	/////////////////////////////////////////////////////////////////////
	response := stub.InvokeChaincode("PqCa", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	fmt.Println("\nVerifyCertChain succeeds! ")
	elapsedTime := time.Since(startTime)
	fmt.Println("The Verify time is ",elapsedTime)
	return shim.Success([]byte("The Verify successfully !!!" ))
}

	/////////////////////////////////////////////////////////////////////
	// Perform the encapsulation operation
	// and upload encapsulated shared secret
	/////////////////////////////////////////////////////////////////////
func (t *user) Encap(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	Serialstring0 :=args[0]
	Serialstring1 :=args[1]
	str1 := []byte(Serialstring0)
	queryArgs := [][]byte{[]byte("Query_Cert"), str1}
	/////////////////////////////////////////////////////////////////////
	// Invoke the function Query_Cert of the chaincode PqCa to get the cert
	/////////////////////////////////////////////////////////////////////
	response := stub.InvokeChaincode("PqCa", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	startTime := time.Now()
	combinecertKEbytes := response.Payload
	certKE,_ :=unmarshalCert_NoPrivKey(combinecertKEbytes)
	/////////////////////////////////////////////////////////////////////
	// Generate encapSessionKey and sessionKey
	/////////////////////////////////////////////////////////////////////
	encapSessionKey,  sessionKey, errs := GenEncapUploadSessionKey(certKE)
	if errs != nil {
		fmt.Println("\nGenEncapUploadSessionKey fails: ", errs)
	}
	fmt.Printf("\n===== The length of encSessionKey is: %d =====\n", len(encapSessionKey))
	fmt.Printf("\n===== The length of sessionKey is: %d =====\n", len(sessionKey))
	Serialstring := Serialstring1 + Serialstring0
	fmt.Printf("\nThe Serialstring of encap is := %s\n", Serialstring)
	/////////////////////////////////////////////////////////////////////
	// Put the encapSessionKey on chain
	/////////////////////////////////////////////////////////////////////
	err := stub.PutState(Serialstring, encapSessionKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The encap time is ",elapsedTime)
	return shim.Success([]byte("encap successfully !!!---------combineserial:"+ Serialstring))
}

	/////////////////////////////////////////////////////////////////////
	// Download the encapsulated share secret
	// and perform decapsulation operation
	/////////////////////////////////////////////////////////////////////
func (t *user) Decap(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	pwKE := args[0]
	combinecert := args[1]
	encapSerialstring := args[2]
	/////////////////////////////////////////////////////////////////////
	// Get the encapSessionKey from chain
	/////////////////////////////////////////////////////////////////////
	encapSessionKey,err :=stub.GetState(encapSerialstring)
	if err != nil {
		return shim.Error(err.Error())
	}
	combinecertKEbytes,_ := base64.StdEncoding.DecodeString(combinecert)
	encertKEbytes :=bytes.Split(combinecertKEbytes, []byte(";;;"))
	passwordke := encertKEbytes[0]
	passwordKE :=string(passwordke[:])
	if passwordKE != pwKE{
		return shim.Error("Incorrect password")
	}
	certKEbytes := encertKEbytes[1]
	certKE,_ :=unmarshalCert(certKEbytes)
	fmt.Printf("\n===== The type of certKE is: %T =====\n", certKE)
	/////////////////////////////////////////////////////////////////////
	// Decap the encapSessionKey to get sessionKey
	/////////////////////////////////////////////////////////////////////
	sk, err5 := DownloadDecapSessionKey(encapSessionKey,  certKE, pwKE)
	if err5 != nil {
		fmt.Println("\nDownloadDecapSessionKey fails: ", err5)
	}
	fmt.Printf("\n===== The length of sk is: %d =====\n", len(sk))
	elapsedTime := time.Since(startTime)
	fmt.Println("The decap time is ",elapsedTime)
	return shim.Success([]byte("decap successfully !!!" ))
}
	/////////////////////////////////////////////////////////////////////
	// Invoke the chaincode PqCa and query one
	// certificate with a given serial number
	/////////////////////////////////////////////////////////////////////
func (t *user) Query_Cert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	Serialstring := args[0]
	str1 := []byte(Serialstring)
	queryArgs := [][]byte{[]byte("Query_Cert"), str1}
	/////////////////////////////////////////////////////////////////////
	// Invoke the function Query_Cert of the the chaincode PqCa
	/////////////////////////////////////////////////////////////////////
	response := stub.InvokeChaincode("PqCa", queryArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to query chaincode.got error :%s", response.Payload))
	}
	fmt.Println("\nquery successfully")
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)
	return shim.Success([]byte("The query successfully !!!" ))
}

func main() {
	layout="Mon Jan 2 15:04:05 -0700 MST 2006"
	err1 := shim.Start(new(user))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}




