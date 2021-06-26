// signature Go example
package main

import (
	"fmt"
	"log"
	"bytes"
	_"io"
	_"encoding/base64"
	//"encoding/base64"
	"crypto/x509"
	"crypto/x509/pkix"
	

	"crypto/sha256"
	"math/big"
	"crypto/rand"
	_"crypto/elliptic"
	"time"
	"encoding/pem"
	"encoding/gob"
	"encoding/base64" 
//	"encoding/hex"
//	"encoding/json"
//	"encoding/asn1"

/*	
	"net"
	"runtime"
*/
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
//	"ecies"
)

var rootCASerial *big.Int
var clk clock
var newserial *big.Int
var CAserial bool
// default system clock
type sysClock struct{}
// time keeper - should return the current time in UTC
type clock interface {
	Now() time.Time
}

func newSysClock() clock {
	c := &sysClock{}
	return c
}

func (c *sysClock) Now() time.Time {
	return time.Now().UTC()
}
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

	//////////////////////////////////////////////////////////////////
	// create the CA
	//////////////////////////////////////////////////////////////////
	// NOTE: 
	// Before creating the CA, the validity check of CA should be performed.
	// One should check whether the CA certificate is revoked or not.
	//////////////////////////////////////////////////////////////////
func CreateCA(clk clock, sub *pkix.Name, pw string, sigName_pq string) (*Cert, error) {

	rootCASerial = randSerial()
	ParentSerialNum := rootCASerial.String()
	fmt.Printf("\nThe rootCASerial := %d\n", rootCASerial)
	//////////////////////////////////////////////////////////////////////////////
	// Generate a pair of PQ signing key 
	// and export its private key in PEM as privKey_pq_PEM
	//////////////////////////////////////////////////////////////////////////////
	signer := oqs.Signature{}
	defer signer.Clean() // clean up even in case of panic

	if err := signer.Init(sigName_pq, nil); err != nil {
		log.Fatal(err)
	}
	pubKey_pq, err := signer.GenerateKeyPair()
	privKey_pq := signer.ExportSecretKey()
	var privKey_pq_PEMBlock *pem.Block
	pass := []byte(pw)
	privKey_pq_PEMBlock, err = x509.EncryptPEMBlock(rand.Reader, "PQ SIGNING PRIVATE KEY", privKey_pq, pass, x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	privKey_pq_PEM := pem.EncodeToMemory(privKey_pq_PEMBlock)
	//////////////////////////////////////////////////////////////////////////////
	// Prepare the Time, Issuer and Subject Name
	//////////////////////////////////////////////////////////////////////////////
	now := clk.Now()
	//////////////////////////////////////////////////////////////////////////////
	// Create the certificate request template
	//////////////////////////////////////////////////////////////////////////////
	var sig_pq []byte
	ck := &Cert{
		version: 			"1.3",
		serialNumber:			rootCASerial,
		issuer: 			*sub,
		subject:			*sub,
		notBefore:			now,
		notAfter:		 	now.Add(365*24*60*60*1000*1000*1000),
		keyUsage:			x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		sigName_pq:			sigName_pq,
		sig_pq:				[]byte(""),

		algName_pq:			sigName_pq,
		privateKey_pq_PEM: 		privKey_pq_PEM,
		pubKey_pq:			pubKey_pq,
		parentKeySerialNum:		ParentSerialNum,

		isKeyEncap:			false,
		isID:				false,
		isCA:				true,
		isRevoked:			false,
	}
	//////////////////////////////////////////////////////////////////////////////
	// Self-sign the certificate info using the PQ signing key
	//////////////////////////////////////////////////////////////////////////////	
	certBytes, _ := marshalCert_NoPrivKey(ck)
	sig_pq, _ = signer.Sign(certBytes)
	ck.sig_pq = sig_pq
	///////////////////////////////////////////////////////////////
	// Return the generated certificate
	///////////////////////////////////////////////////////////////
	return ck, nil
}

func NewSerial(caSerial *big.Int) (*big.Int) {
	new := big.NewInt(1)
	caSerial.Add(caSerial, new)
	new.Set(caSerial)
	return new
}

	//////////////////////////////////////////////////////////////////
	// Generate Certificate for Individual or Company
	//////////////////////////////////////////////////////////////////
	// NOTE: 
	// Before the certificate creation, the CA validity check should be performed.
	// One should check whether the CA certificate is revoked or not.
	//////////////////////////////////////////////////////////////////
func GenerateIDCert(clk clock, certCA *Cert, pwCA string, sub *pkix.Name, pwID string, sigName_pq string,ParentSerialNum string) (*Cert, error) {
	serial := NewSerial(rootCASerial)
	fmt.Printf("\nThe new serial := %d\n", serial)
	//////////////////////////////////////////////////////////////////////////////
	// Generate a pair of pq signing key for an ID
	// and export its private key in PEM as privKey_pq_PEM
	//////////////////////////////////////////////////////////////////////////////
	signer := oqs.Signature{}
	defer signer.Clean() // clean up even in case of panic
	if err := signer.Init(sigName_pq, nil); err != nil {
		log.Fatal(err)
	}

	pubKey_pq, err := signer.GenerateKeyPair()
	privKey_pq := signer.ExportSecretKey()
	//////////////////////////////////////////////////////////////////////////////
	// Encrypt the newly-generated PQ private key
	//////////////////////////////////////////////////////////////////////////////
	pass := []byte(pwID)
	var privKey_pq_PEMBlock *pem.Block
	privKey_pq_PEMBlock, err = x509.EncryptPEMBlock(rand.Reader, "PQ SIGNING PRIVATE KEY", privKey_pq, pass, x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	privKey_pq_PEM := pem.EncodeToMemory(privKey_pq_PEMBlock)

	//////////////////////////////////////////////////////////////////////////////
	// Decrypt the PQ private key  based on the CA
	// in order to sign the newly-generated certificate
	//////////////////////////////////////////////////////////////////////////////
	passCA := []byte(pwCA)
	privKey_pq_PEMBlock_CA, _ := pem.Decode(certCA.privateKey_pq_PEM)
	privKey_pq_CA, _ := x509.DecryptPEMBlock(privKey_pq_PEMBlock_CA, passCA)
	//////////////////////////////////////////////////////////////////////////////
	// Prepare the Time, Issuer and Subject Name
	//////////////////////////////////////////////////////////////////////////////
	now := clk.Now()
	issuer := certCA.issuer
	var sig_pq []byte
	ck := &Cert{
		version: 			"1.3",
		serialNumber:			serial,
		issuer: 			issuer,
		subject:			*sub,
		notBefore:			now,
		notAfter:		 	now.Add(365*24*60*60*1000*1000*1000),
		keyUsage:			x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		sigName_pq:			certCA.algName_pq,		// string
		sig_pq:				[]byte(""),				// []byte

		algName_pq:			sigName_pq,
		privateKey_pq_PEM:		privKey_pq_PEM,
		pubKey_pq:			pubKey_pq,		// []byte
		parentKeySerialNum:		ParentSerialNum,

		isKeyEncap:			false,
		isID:				true,
		isCA:				false,
		isRevoked:			false,
	}
	////////////////////////////////////////////////////////////////
	// Initialize the PQ signer based on the CA PQ signing secret key
	////////////////////////////////////////////////////////////////
	signer_CA := oqs.Signature{}
	defer signer_CA.Clean() // clean up even in case of panic
	if err := signer_CA.Init(certCA.algName_pq, privKey_pq_CA); err != nil {
		log.Fatal(err)
	}
	certBytes, _ := marshalCert_NoPrivKey(ck)
	////////////////////////////////////////////////////////////////
	// Generate the PQ signature "sig_pq" for the Cert structure "ck"
	////////////////////////////////////////////////////////////////
	sig_pq, _ = signer_CA.Sign(certBytes)
	// Store the generated PQ signature in the Cert structure
	ck.sig_pq = sig_pq
	////////////////////////////////////////////////////////////////
	// Signature validation for sanity check
	////////////////////////////////////////////////////////////////
	return ck, nil
}

	//////////////////////////////////////////////////////////////////////////////////
	// Based on ID certificate, generate Key Encapsulation 
	// (key pair) Certificate for the same ID
	//////////////////////////////////////////////////////////////////////////////////
	// NOTE: Before the certificate creation, 
	// the CA and ID certificates validity check should be performed.
	// One should check whether the CA and ID certificates are revoked or not.
	//////////////////////////////////////////////////////////////////////////////////
func GenerateKeyEncapCert(clk clock, certID *Cert, pwID string, pwKeyEncap string, kemName_pq string,ParentSerialNum string) (*Cert, error) {
	serial := NewSerial(rootCASerial)
	fmt.Printf("\nThe new serial := %d\n", serial)
	//////////////////////////////////////////////////////////////////////////////
	// Generate a pair of pq signing key for a KE
	// and export its private key in PEM as privateKey_pq_PEM
	//////////////////////////////////////////////////////////////////////////////
	kemer := oqs.KeyEncapsulation{}
	defer kemer.Clean() // clean up even in case of panic
	if err := kemer.Init(kemName_pq, nil); err != nil {
		log.Fatal(err)
	}
	pubKey_pq, err := kemer.GenerateKeyPair()
	privateKey_pq := kemer.ExportSecretKey()
	//////////////////////////////////////////////////////////////////////////////
	// Encrypt the newly-generated PQ private key
	//////////////////////////////////////////////////////////////////////////////
	pass := []byte(pwKeyEncap)
	var privateKey_pq_PEMBlock *pem.Block
	privateKey_pq_PEMBlock, err = x509.EncryptPEMBlock(rand.Reader, "PQ SIGNING PRIVATE KEY", privateKey_pq, pass, x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	privateKey_pq_PEM := pem.EncodeToMemory(privateKey_pq_PEMBlock)
	//////////////////////////////////////////////////////////////////////////////
	// Decrypt the PQ private key based the given ID
	// in order to sign the newly-generated Key Encapsulation certificate
	//////////////////////////////////////////////////////////////////////////////
	passID := []byte(pwID)
	privateKey_pq_PEMBlock_ID, _ := pem.Decode(certID.privateKey_pq_PEM)
	privateKey_pq_ID, _ := x509.DecryptPEMBlock(privateKey_pq_PEMBlock_ID, passID)
	//////////////////////////////////////////////////////////////////////////////
	// Prepare the Time, Issuer and Subject Name
	//////////////////////////////////////////////////////////////////////////////
	now := clk.Now()
	issuer := certID.issuer
	//////////////////////////////////////////////////////////////////////////////
	// Create the certificate request template
	//////////////////////////////////////////////////////////////////////////////
	var sig_pq []byte
	ck := &Cert{
		version: 			"1.3",
		serialNumber:			serial,
		issuer: 			issuer,
		subject:			issuer,
		notBefore:			now,
		notAfter:		 	now.Add(365*24*60*60*1000*1000*1000),
		keyUsage:			x509.KeyUsageKeyEncipherment,
		sigName_pq:			certID.algName_pq,		// string
		sig_pq:				[]byte(""),				// []byte

		algName_pq:			kemName_pq,
		privateKey_pq_PEM:		privateKey_pq_PEM,
		pubKey_pq:			pubKey_pq,			// []byte
		parentKeySerialNum:		ParentSerialNum,

		isKeyEncap:			true,
		isID:				false,
		isCA:				false,
		isRevoked:			false,
	}
	////////////////////////////////////////////////////////////////
	// Initialize the PQ signer based on the ID PQ signing secret key
	////////////////////////////////////////////////////////////////
	signer_ID := oqs.Signature{}
	defer signer_ID.Clean() // clean up even in case of panic
	sigName_pq_ID := certID.algName_pq
	if err := signer_ID.Init(sigName_pq_ID, privateKey_pq_ID); err != nil {
		log.Fatal(err)
	}
	certBytes, _ := marshalCert_NoPrivKey(ck)
	////////////////////////////////////////////////////////////////
	// Generate the PQ signature "sig_pq" for the Cert structure "ck"
	////////////////////////////////////////////////////////////////
	sig_pq, _ = signer_ID.Sign(certBytes)
	// Store the generated PQ signature in the Cert structure
	ck.sig_pq = sig_pq
	return ck, nil
}

func VerifyCertChain(certCA *Cert_NoPrivKey, certID *Cert_NoPrivKey, certKE *Cert_NoPrivKey)(bool, string){
	startTime1 := time.Now()
	/////////////////////////////////////////////////////////////////////
	// Check whether all the certs in the chain have been revoked or not
	/////////////////////////////////////////////////////////////////////
	if certCA.isRevoked == true {
		return false, "The CA cert has been revoked"
	}

	if certID.isRevoked == true {
		return false, "The ID cert has been revoked"
	}

	if certKE.isRevoked == true {
		return false, "The KE cert has been revoked"
	}
        
	/////////////////////////////////////////////////////////////////////
	// Verify the signature signed by the CA 
	/////////////////////////////////////////////////////////////////////
	fmt.Println("\nVerification on signature signed by CA succeeds")
	/////////////////////////////////////////////////////////////////////
	// Verify the PQ signature signed by the CA 
	/////////////////////////////////////////////////////////////////////
	pubKey_pq_CA := certCA.pubKey_pq
	algName_pq_CA := certCA.algName_pq
	sig_pq_ID := certID.sig_pq
	ck_ID := &Cert{
		version: 			certID.version,
		serialNumber:			certID.serialNumber,
		issuer: 			certID.issuer,
		subject:			certID.subject,
		notBefore:			certID.notBefore,
		notAfter:		 	certID.notAfter,
		keyUsage:			certID.keyUsage,
		sigName_pq:			certID.algName_pq,		// string
		sig_pq:				[]byte(""),				// []byte

		algName_pq:			certID.algName_pq,		// string
		pubKey_pq:			certID.pubKey_pq,	// []byte
		parentKeySerialNum:		certID.parentKeySerialNum,

		isKeyEncap:			certID.isKeyEncap,
		isID:				certID.isID,
		isCA:				certID.isCA,
		isRevoked: 			certID.isRevoked,
	}
	ck_ID_bytes, err := marshalCert(ck_ID)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	verifierCA := oqs.Signature{}
	defer verifierCA.Clean() // clean up even in case of panic

	if err := verifierCA.Init(algName_pq_CA, nil); err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	isValid, err := verifierCA.Verify(ck_ID_bytes, sig_pq_ID, pubKey_pq_CA)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	fmt.Println("\nVerification on PQ signature signed by CA succeeds", isValid)
	/////////////////////////////////////////////////////////////////////
	// Verify the signature signed by the ID 
	/////////////////////////////////////////////////////////////////////
	fmt.Println("\nVerification on signature signed by the ID succeeds")
	elapsedTime1 := time.Since(startTime1)
	fmt.Println("The certCA and certID Verify time is ",elapsedTime1)
	startTime2 := time.Now()
	/////////////////////////////////////////////////////////////////////
	// Verify the PQ signature signed by the CA 
	/////////////////////////////////////////////////////////////////////
	pubKey_pq_ID := certID.pubKey_pq
	algName_pq_ID := certID.algName_pq
	sig_pq_KE := certKE.sig_pq
	ck_KE := &Cert{
		version: 			certKE.version,
		serialNumber:			certKE.serialNumber,
		issuer: 			certKE.issuer,
		subject:			certKE.subject,
		notBefore:			certKE.notBefore,
		notAfter:		 	certKE.notAfter,
		keyUsage:			certKE.keyUsage,
		sigName_pq:			certKE.algName_pq,		// string
		sig_pq:				[]byte(""),		

		algName_pq:			certKE.algName_pq,		// string
		pubKey_pq:			certKE.pubKey_pq,		// []byte
		parentKeySerialNum:		certKE.parentKeySerialNum,

		isKeyEncap:			certKE.isKeyEncap,
		isID:				certKE.isID,
		isCA:				certKE.isCA,
		isRevoked: 			certKE.isRevoked,
	}
	ck_KE_bytes, err := marshalCert(ck_KE)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	verifierID := oqs.Signature{}
	defer verifierID.Clean() // clean up even in case of panic

	if err := verifierID.Init(algName_pq_ID, nil); err != nil {
		log.Fatal(err)
		return false, err.Error()
	}
	isValid2, err2 := verifierID.Verify(ck_KE_bytes, sig_pq_KE, pubKey_pq_ID)
	if err2 != nil {
		log.Fatal(err2)
		return false, err2.Error()
	}
	fmt.Println("\nVerification on PQ signature signed by the ID succeeds", isValid2)
	elapsedTime2 := time.Since(startTime2)
	fmt.Println("The certKE and certID Verify time is ",elapsedTime2)
	return true, "Verification on the cert chain succeeds"
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

func RevokeCert(cert *Cert)(*Cert,string){
	if cert.isCA != false{
		CAserial =false
	}
	if cert.isRevoked != false{
		return cert,"the cert has been revoked"
	}
	cert.isRevoked = true
	return cert,"right"
}
	/////////////////////////////////////////////////////////////////////
	// Verify the certKE and its parent certs whether they are revoked
	/////////////////////////////////////////////////////////////////////
func VerifyCertChain_ThreeLayers(certCA *Cert_NoPrivKey, certID *Cert_NoPrivKey, certKE *Cert_NoPrivKey)(bool, string){
	if certCA.isRevoked == true {
		return false, "The CA cert has been revoked"
	}
	if certID.isRevoked == true {
		return false, "The ID cert has been revoked"
	}
	if certKE.isRevoked == true {
		return false, "The KE cert has been revoked"
	}
	return true, "no cert has been revoked"
}

	/////////////////////////////////////////////////////////////////////
	// Verify the certID and its parent cert whether they are revoked
	/////////////////////////////////////////////////////////////////////
func VerifyCertChain_TwoLayers(certCA *Cert_NoPrivKey, certID *Cert_NoPrivKey)(bool, string){
	if certCA.isRevoked == true {
		return false, "The CA cert has been revoked"
	}

	if certID.isRevoked == true {
		return false, "The ID cert has been revoked"
	}
	return true, "no cert has been revoked"
}

	/////////////////////////////////////////////////////////////////////
	// Verify the certCA  whether it is revoked
	/////////////////////////////////////////////////////////////////////
func VerifyCertChain_OneLayer(certCA *Cert_NoPrivKey)(bool, string){
	if certCA.isRevoked == true {
		return false, "The CA cert has been revoked"
	}
	return true, "no cert has been revoked"
}

	/////////////////////////////////////////////////////////////////////
	// Verify the identity of the user
	/////////////////////////////////////////////////////////////////////
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

type ca struct{}

func (t *ca) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("begin to init chaincode....")

   return shim.Success([]byte("Success invoke and not opter!!"))
}

func (t *ca) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
		if fn == "Gen_CertCA" {
		return t.Gen_CertCA(stub, args)
	}else if fn == "Gen_CertID" {
		return t.Gen_CertID(stub, args)
	}else if fn == "Gen_CertKE" {
		return t.Gen_CertKE(stub, args)
	}else if fn == "Verify_Cert" { // Given CertKE, CertID-->CertKE,
		return t.Verify_Cert(stub, args)
	}else if fn == "Revoke_Cert" {
		return t.Revoke_Cert(stub, args)
	}else if fn == "Query_Cert" {
		return t.Query_Cert(stub, args)
	}
	return shim.Error("Recevied unkown function invocation")
}

	/////////////////////////////////////////////////////////////////////
	// Generate the key pair and certificate (self-signed) for the CA
	/////////////////////////////////////////////////////////////////////
func (t *ca) Gen_CertCA(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	clk = newSysClock()
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	// verify identity of the CA
	creator := GetCreator(stub)
	if creator == "" {
		return shim.Error("The operator of Getcreator is failed!")
	}
	if (creator != "Admin@org0.example.com") {
		return shim.Error("The escrower doesn't have authority, so it can not modify the information!---creator:" + creator)
	}
	/////////////////////////////////////////////////////////////////////
	// Verify that the CA is unique
	// and read relevant information
	/////////////////////////////////////////////////////////////////////
	if CAserial != false{
	return shim.Error("The certCA has been registered!")
	}
	sigName_pq := args[0]
	pwCA := args[1]
	str0  := []string{args[2]}
	str1  := []string{args[3]}
	str2  := []string{args[4]}
	str3  := []string{args[5]}
	str4  := []string{args[6]}
	str5  := []string{args[7]}
	str6  := []string{args[8]}
	str7 := args[9]
	subCA := &pkix.Name{
		Country: 			str0,
		Organization: 			str1,
		OrganizationalUnit: 		str2,
		Locality: 			str3,
		Province: 			str4,
		StreetAddress: 			str5,
		PostalCode: 			str6,
		// SerialNumber: 		"CN",
		CommonName: 			str7,
	}
	/////////////////////////////////////////////////////////////////////
	// Generate certCA 
	/////////////////////////////////////////////////////////////////////		
	certCA, _ := CreateCA(clk, subCA, pwCA, sigName_pq)
	fmt.Printf("\n===== The type of certCA is: %T =====\n", certCA)
	certCAbytes,_ := marshalCert(certCA)
	certCAbytes_NoPrivKey,_ := marshalCert_NoPrivKey(certCA)
	rootCASerialstring :=rootCASerial.String()
	fmt.Printf("\nThe rootCASerial of certCA which putstate is := %s\n", rootCASerialstring)
	pwCAbytes :=[]byte(pwCA)
	combinecertCAbytes :=bytes.Join([][]byte{pwCAbytes, certCAbytes}, []byte(";;;"))
	/////////////////////////////////////////////////////////////////////
	// Put the certCA on chain 
	/////////////////////////////////////////////////////////////////////
	err := stub.PutState(rootCASerialstring,certCAbytes_NoPrivKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("certCA size:",len(certCAbytes_NoPrivKey))
	elapsedTime := time.Since(startTime)
	fmt.Println("The certCA time is ",elapsedTime)
	CAserial = true
	return shim.Success(combinecertCAbytes)
}

	/////////////////////////////////////////////////////////////////////
	// Generate the key pair and ID certificate (signed by the CA) for each user
	/////////////////////////////////////////////////////////////////////
func (t *ca) Gen_CertID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 12 {
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}
	// verify identity of the CA
	creator := GetCreator(stub)
	if creator == "" {
		return shim.Error("The operator of Getcreator is failed!")
	}
	if (creator != "Admin@org0.example.com") {
		return shim.Error("The escrower doesn't have authority, so it can not modify the information!---creator:" + creator)
	}
	sigName_pq := args[0]
	pwCA := args[1]
	pwID := args[2]
	combinecertCAstring := args[3]
	str0  := []string{args[4]}
	str1  := []string{args[5]}
	str2  := []string{args[6]}
	str3  := []string{args[7]}
	str4  := []string{args[8]}
	str5  := []string{args[9]}
	str6  := []string{args[10]}
	str7 := args[11]
	subID := &pkix.Name{
		Country: 			str0,
		Organization: 			str1,
		OrganizationalUnit: 		str2,
		Locality: 			str3,
		Province: 			str4,
		StreetAddress: 			str5,
		PostalCode: 			str6,
		CommonName: 			str7,
	}
	/////////////////////////////////////////////////////////////////////
	// Decode the certCAbytes
	/////////////////////////////////////////////////////////////////////
	combinecertCAbytes,err := base64.StdEncoding.DecodeString(combinecertCAstring)
	if err!=nil{
		fmt.Println(err)
	}
	enConBytes := bytes.Split(combinecertCAbytes, []byte(";;;"))
	CApw := enConBytes[0]
	pwCAsecond := string(CApw[:])
	/////////////////////////////////////////////////////////////////////
	// Verify password
	/////////////////////////////////////////////////////////////////////
	if pwCAsecond != pwCA{
		return shim.Error("Incorrect password")
	}
	fmt.Printf("correct password")
	certCAbytes := enConBytes[1]
	certCA,_ :=unmarshalCert(certCAbytes)
	rootCASerial2 := certCA.serialNumber
	rootCASerialstring := rootCASerial2.String()
	/////////////////////////////////////////////////////////////////////
	// Generate certID 
	/////////////////////////////////////////////////////////////////////
	certID, _ := GenerateIDCert(clk, certCA, pwCA, subID, pwID, sigName_pq,rootCASerialstring)
	fmt.Printf("\n===== The type of certID is: %T =====\n", certID)
	newserial := rootCASerial
	fmt.Printf("\nThe serial of certID which putsate is := %d\n", newserial)
	serialstring :=newserial.String()
	certIDbytes,_ := marshalCert(certID)
	certIDbytes_NoPrivKey,_ := marshalCert_NoPrivKey(certID)
	pwIDbytes :=[]byte(pwID)
	combinecertIDbytes :=bytes.Join([][]byte{pwIDbytes, certIDbytes}, []byte(";;;"))
	/////////////////////////////////////////////////////////////////////
	// Put the certID on chain 
	/////////////////////////////////////////////////////////////////////
	err = stub.PutState(serialstring,certIDbytes_NoPrivKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("certID size:",len(certIDbytes_NoPrivKey))
	elapsedTime := time.Since(startTime)
	fmt.Println("The certIDtime is ",elapsedTime)
	return shim.Success(combinecertIDbytes)
}

	/////////////////////////////////////////////////////////////////////
	// Generate the KEM key pair and KE certifate for each user
	/////////////////////////////////////////////////////////////////////
func (t *ca) Gen_CertKE(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	// verify identity of the CA
	creator := GetCreator(stub)
	if creator == "" {
		return shim.Error("The operator of Getcreator is failed!")
	}
	if (creator != "Admin@org0.example.com") {
		return shim.Error("The escrower doesn't have authority, so it can not modify the information!---creator:" + creator)
	}
	kemName_pq := args[0]
	pwID := args[1]
	pwKE := args[2]
	combinecertIDstring := args[3]
	/////////////////////////////////////////////////////////////////////
	// Decode the certIDbytes
	/////////////////////////////////////////////////////////////////////
	combinecertIDbytes,_ := base64.StdEncoding.DecodeString(combinecertIDstring)
	enConBytes := bytes.Split(combinecertIDbytes, []byte(";;;"))
	convert := enConBytes[0]
	pwIDsecond := string(convert[:])
	/////////////////////////////////////////////////////////////////////
	// Verify password
	/////////////////////////////////////////////////////////////////////
	if pwIDsecond != pwID{
		return shim.Error("Incorrect password")
	}
	fmt.Printf("PwIDcorrect password")
	certIDbytes := enConBytes[1]
	certID,_ :=unmarshalCert(certIDbytes)
	certIDSerialstring := certID.serialNumber
	Serialstring := certIDSerialstring.String()
	fmt.Printf("\n===== The type of certID is: %T =====\n", certID)
	/////////////////////////////////////////////////////////////////////
	// Generate certKE 
	/////////////////////////////////////////////////////////////////////
	certKE, _ := GenerateKeyEncapCert(clk, certID, pwID, pwKE, kemName_pq,Serialstring)
	fmt.Printf("\n===== The type of certKE is: %T =====\n", certKE)
	newserial2 := rootCASerial
	fmt.Printf("\nThe serial of certKE which putsate is := %d\n", newserial2)
	serialstring2 :=newserial2.String()
	certKEbytes,_ := marshalCert(certKE)
	certKEbytes_NoPrivKey,_ := marshalCert_NoPrivKey(certKE)
	pwKEbytes :=[]byte(pwKE)
	combinecertKEbytes :=bytes.Join([][]byte{pwKEbytes,certKEbytes}, []byte(";;;"))
	/////////////////////////////////////////////////////////////////////
	// Put the certKE on chain 
	/////////////////////////////////////////////////////////////////////
	err := stub.PutState(serialstring2 ,certKEbytes_NoPrivKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("certKE size:",len(certKEbytes_NoPrivKey))
	elapsedTime := time.Since(startTime)
	fmt.Println("The certKE time is ",elapsedTime)
		return shim.Success(combinecertKEbytes)
}

	/////////////////////////////////////////////////////////////////////
	// Verify one certificate and its parent
	// certificates in the certificate chain
	/////////////////////////////////////////////////////////////////////
func (t *ca) Verify_Cert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Serialstring2 := args[0]
	/////////////////////////////////////////////////////////////////////
	// Get the certKE from pubilc ledger
	/////////////////////////////////////////////////////////////////////
	certKEbytes, err1 := stub.GetState(Serialstring2)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certKE,_ :=unmarshalCert_NoPrivKey(certKEbytes)
	fmt.Printf("\n===== The type of certKE is: %T =====\n", certKE)
	Serialstring1 := certKE.parentKeySerialNum
	/////////////////////////////////////////////////////////////////////
	// Get the certID from pubilc ledger
	/////////////////////////////////////////////////////////////////////
	certIDbytes, err1 := stub.GetState(Serialstring1)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certID,_ :=unmarshalCert_NoPrivKey(certIDbytes)
	fmt.Printf("\n===== The type of certID is: %T =====\n", certID)
	Serialstring0 := certID.parentKeySerialNum
	/////////////////////////////////////////////////////////////////////
	// Get the certCA from pubilc ledger 
	/////////////////////////////////////////////////////////////////////
	certCAbytes, err1 := stub.GetState(Serialstring0)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certCA,_ :=unmarshalCert_NoPrivKey(certCAbytes)
	fmt.Printf("\n===== The type of certID is: %T =====\n", certCA)
	isVerified, wrong := VerifyCertChain(certCA, certID, certKE)
	if isVerified != true {		
		return shim.Error(wrong)
		//panic(err)
	}
	fmt.Println("\nVerifyCertChain succeeds! true")
	elapsedTime := time.Since(startTime)
	fmt.Println("The all Verify time is ",elapsedTime)
	return shim.Success([]byte("The Verify successfully !!!" ))
}

	/////////////////////////////////////////////////////////////////////
	// Revoke one certificate based on a given seial number
	/////////////////////////////////////////////////////////////////////
func (t *ca) Revoke_Cert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// verify identity of the CA
	creator := GetCreator(stub)
	if creator == "" {
		return shim.Error("The operator of Getcreator is failed!")
	}
	if (creator != "Admin@org0.example.com") {
		return shim.Error("The escrower doesn't have authority, so it can not modify the information!---creator:" + creator)
	}
	pwcert := args[0]
	combinecertIDstring :=args[1]
	combinecertbytes,_ := base64.StdEncoding.DecodeString(combinecertIDstring)
	encertbytes := bytes.Split(combinecertbytes, []byte(";;;"))
	passwordstring := encertbytes[0]
	password :=string(passwordstring[:])
	if password != pwcert{
		return shim.Error("Incorrect password")
	}
	fmt.Println(" password correct")
	certbytes := encertbytes[1]
	cert,_ :=unmarshalCert(certbytes)
	/////////////////////////////////////////////////////////////////////
	// Revoke the cert
	/////////////////////////////////////////////////////////////////////
	revokedcert,reason := RevokeCert(cert)
	if reason != "right"{
		return shim.Error(reason)
	}
	certSerial := revokedcert.serialNumber
	Serialstring:=certSerial.String()
	certbytes,_ = marshalCert(revokedcert)
	certBytes_NoPrivKey,_ := marshalCert_NoPrivKey(revokedcert)
	pwbytes := []byte(pwcert)
	revokedcombinecertbytes := bytes.Join([][]byte{pwbytes,certbytes}, []byte(";;;"))
	/////////////////////////////////////////////////////////////////////
	// Put the revoked cert on chain 
	/////////////////////////////////////////////////////////////////////
	err := stub.PutState(Serialstring, certBytes_NoPrivKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The revoke time is ",elapsedTime)
	return shim.Success(revokedcombinecertbytes)
}

	/////////////////////////////////////////////////////////////////////
	// Query one certificate based on a given serial number
	/////////////////////////////////////////////////////////////////////
func (t *ca) Query_Cert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startTime := time.Now()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Serialstring2 := args[0]
	combinecertbytes, err1 := stub.GetState(Serialstring2)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	cert,_ :=unmarshalCert_NoPrivKey(combinecertbytes)
	/////////////////////////////////////////////////////////////////////
	// Verify the cert and parent certs whether are revoked
	/////////////////////////////////////////////////////////////////////
	if cert.isKeyEncap == true {
	certKE := cert
	fmt.Printf("\n===== The type of certKE is: %T =====\n", certKE)
	Serialstring1 := certKE.parentKeySerialNum
	certIDbytes, err1 := stub.GetState(Serialstring1)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certID,_ :=unmarshalCert_NoPrivKey(certIDbytes)
	fmt.Printf("\n===== The type of certID is: %T =====\n", certID)
	Serialstring0 := certID.parentKeySerialNum
	certCAbytes, err1 := stub.GetState(Serialstring0)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certCA,_ :=unmarshalCert_NoPrivKey(certCAbytes)
	fmt.Printf("\n===== The type of certCA is: %T =====\n", certCA)
	isVerified, wrong := VerifyCertChain_ThreeLayers(certCA, certID, certKE)
	if isVerified != true {		
		return shim.Error(wrong)
		//panic(err)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)
	return shim.Success(combinecertbytes)
	}
	if cert.isID == true{
	certID := cert
	fmt.Printf("\n===== The type of certID is: %T =====\n", certID) 
	Serialstring0 := certID.parentKeySerialNum
	certCAbytes, err1 := stub.GetState(Serialstring0)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	certCA,_ :=unmarshalCert_NoPrivKey(certCAbytes)
	fmt.Printf("\n===== The type of certCA is: %T =====\n", certCA)
	isVerified, wrong := VerifyCertChain_TwoLayers(certCA, certID)
	if isVerified != true {		
		return shim.Error(wrong)
		//panic(err)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)
	return shim.Success(combinecertbytes)
	}
	if cert.isCA == true{
	certCA := cert
	isVerified, wrong := VerifyCertChain_OneLayer(certCA)
	if isVerified != true {		
		return shim.Error(wrong)
		//panic(err)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)
	return shim.Success(combinecertbytes)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)
	return shim.Error("query failed")
}

func main() {
	layout="Mon Jan 2 15:04:05 -0700 MST 2006"
	err1 := shim.Start(new(ca))
	if err1 != nil {
		fmt.Printf("error starting simple chaincode:%s \n", err1)
	}
}


