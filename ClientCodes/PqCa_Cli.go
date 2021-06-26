package main
 
import (
	"os"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	_"log"
	"bytes"
	_"io"
	"encoding/base64"
	//"encoding/base64"
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/sha256"
	"math/big"
	"crypto/rand"
	"time"
	_"encoding/pem"
	"encoding/gob"
	
 
//	"encoding/hex"
//	"encoding/json"
//	"encoding/asn1"


	
)
/////////////////////////////////////////////////////////////////////
// Channel related information for the client
/////////////////////////////////////////////////////////////////////
var (
        cc          = "PqCa"
        user        = "Admin" 
        secret      = ""
        channelName = "mychannel"
        lvl         = logging.INFO
        orgName     = "Org0MSP"
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
	isCA 				bool
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

func marshalCert_NoPrivKey(crt *Cert_NoPrivKey)([]byte, error){
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

	/////////////////////////////////////////////////////////////////////
	// The client codes for CA adminitrator (user) to query its private database
	/////////////////////////////////////////////////////////////////////
func Query(db *sql.DB,key string)string {
	rows,err:=db.Query("select * from ca where Serial= ?",key)
	if err!=nil{
		fmt.Println("db.Query error:",err)
		return "db.Query error:"
	}
 	var serial string
		var cert string
	for rows.Next() {
		error1 := rows.Scan(&serial, &cert)
		if error1 != nil {
			fmt.Println("rows.Scan error:", error1)
			return "rows.Scan error:"
		}
	}
	return cert
	}

func QueryUser(db *sql.DB,key string)string {
	rows,err:=db.Query("select * from user where Serial= ?",key)
	if err!=nil{
		fmt.Println("db.Query error:",err)
		return "db.Query error:"
	}
 	var serial string
		var cert string
	for rows.Next() {
		error1 := rows.Scan(&serial, &cert)
		if error1 != nil {
			fmt.Println("rows.Scan error:", error1)
			return "rows.Scan error:"
		}
	}
	return cert
	}

type User struct {
	Serial 	string 
	Cert 	string
}

	/////////////////////////////////////////////////////////////////////
	// The client codes for CA adminitrator (user) to insert its private database
	/////////////////////////////////////////////////////////////////////
func Insert(db *sql.DB,u1 User){
	stmt,pError:=db.Prepare("insert into ca (Serial,Cert) values (?,?)")
	defer stmt.Close()
	if pError!=nil{
		fmt.Println("db.Prepare error:",pError)
		return
	}
	res,error:=stmt.Exec(u1.Serial,u1.Cert)
	if error!=nil{
		fmt.Println("stmt.Exec error:",error)
		return
	}
	flag,_:=res.RowsAffected()
	fmt.Println("Affected lines:",flag)
 
}

func InsertUser(db *sql.DB,u1 User){
	stmt,pError:=db.Prepare("insert into user (Serial,Cert) values (?,?)")
	defer stmt.Close()
	if pError!=nil{
		fmt.Println("db.Prepare error:",pError)
		return
	}
	res,error:=stmt.Exec(u1.Serial,u1.Cert)
	if error!=nil{
		fmt.Println("stmt.Exec error:",error)
		return
	}
	flag,_:=res.RowsAffected()
	fmt.Println("Affected lines:",flag)
 
}

	/////////////////////////////////////////////////////////////////////
	// The client codes for CA adminitrator (user) to update its private database
	/////////////////////////////////////////////////////////////////////
func Update(db *sql.DB,u1 User){
	stmt,pError:=db.Prepare("update ca set Serial=?,Cert=?")
	defer stmt.Close()
	if pError!=nil{
		fmt.Println("db.Prepare error:",pError)
		return
	}
	res,error:=stmt.Exec(u1.Serial,u1.Cert)
	if error!=nil{
		fmt.Println("stmt.Exec error:",error)
		return
	}
	flag,_:=res.RowsAffected()
	fmt.Println("Affected lines:",flag)
 
}

func UpdateUser(db *sql.DB,u1 User){
	stmt,pError:=db.Prepare("update user set Serial=?,Cert=?")
	defer stmt.Close()
	if pError!=nil{
		fmt.Println("db.Prepare error:",pError)
		return
	}
	res,error:=stmt.Exec(u1.Serial,u1.Cert)
	if error!=nil{
		fmt.Println("stmt.Exec error:",error)
		return
	}
	flag,_:=res.RowsAffected()
	fmt.Println("Affected lines:",flag)
 
}

func main(){
	startTime := time.Now()
	/////////////////////////////////////////////////////////////////////
	// Read the configuration file
	/////////////////////////////////////////////////////////////////////
	c := config.FromFile("./connection-profile.yaml")
	sdk, err := fabsdk.New(c)
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		os.Exit(1)
	}
	defer sdk.Close()
	/////////////////////////////////////////////////////////////////////
	// Establish a connection with the channel
	/////////////////////////////////////////////////////////////////////
	clientChannelContext := sdk.ChannelContext(channelName, fabsdk.WithUser(user), fabsdk.WithOrg(orgName))
	if err != nil {
		fmt.Printf("Failed to create channel [%s] client: %#v", channelName, err)
		os.Exit(1)
	}
	client, err := channel.New(clientChannelContext)
	if err != nil {
		fmt.Printf("Failed to create channel [%s]:", channelName, err)
	}
	args := os.Args
	num :=args[1]
	if num =="Gen_CertCA"{		
		Sig_alg := args[2]
		Password := args[3]
		Country := args[4]
		Organization := args[5]
		OrganizationUnit := args[6]
		Locality := args[7]
		Province := args[8]
		StreetAddress := args[9]
		PostalCode := args[10]
		CommonName := args[11]
		certCA_CC(client, Sig_alg, Password, Country, Organization,OrganizationUnit, 
					Locality, Province, StreetAddress, PostalCode, CommonName)
				
	}
	if num =="Gen_CertID"{		
		Sig_alg := args[2]
		Password_CA  := args[3]
		Password_ID  := args[4]
		certCASerial := args[5]
		Country := args[6]
		Organization := args[7]
		OrganizationUnit := args[8]
		Locality := args[9]
		Province := args[10]
		StreetAddress := args[11]
		PostalCode := args[12]
		CommonName := args[13]
		certID_CC(client , Sig_alg , Password_CA , Password_ID , certCASerial , Country , Organization , OrganizationUnit , Locality , Province , StreetAddress , PostalCode , CommonName )
				
	}
	if num =="Gen_CertKE"{		
		Chg_sig := args[2]
		Password_ID  := args[3]
		Password_KE  := args[4]
		certIDSerial := args[5]
		certKE_CC(client, Chg_sig, Password_ID, Password_KE, certIDSerial)
				
	}
	if num =="Verify_Cert"{		
		newValue := args[2]
		verify_CC(client, newValue )
				
	}
	
	if num =="Revoke_Cert"{		
		Password_Cert := args[2]
		ID_Cert := args[3]
		revoke_CC(client, Password_Cert, ID_Cert )
	}		
	if num =="Query_Cert"{		
		name := args[2]
		query_CC(client, name)
				
	}		
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)	
	}

	/////////////////////////////////////////////////////////////////////
	//  Here are the fuctions which invoke the corresponding chaincode APIs
	/////////////////////////////////////////////////////////////////////
	//1---certCA
func certCA_CC(client *channel.Client, Sig_alg string, Password string, Country string, Organization string,OrganizationUnit string, Locality string, Province string, StreetAddress string, PostalCode string, CommonName string) {
	certCA_Args := [][]byte{[]byte(Sig_alg), []byte(Password), []byte(Country), []byte(Organization), 
				[]byte(OrganizationUnit), []byte(Locality),[]byte(Province), 
				[]byte(StreetAddress), []byte(PostalCode), []byte(CommonName)}
	/////////////////////////////////////////////////////////////////////
	// Invoke the Chaincode PqCa
	/////////////////////////////////////////////////////////////////////
	response, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Gen_CertCA",
		Args:        certCA_Args,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))
	if err != nil {
		fmt.Printf("Failed to certCA: %+v\n", err)
	}
	
	commsg := response.Payload
	enConBytes := bytes.Split(commsg, []byte(";;;"))
	msg := enConBytes[1]
	certCA,_ := unmarshalCert(msg)
	Serial := certCA.serialNumber
	db, err := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err != nil {
		fmt.Println(err);
	}
	encoding := base64.StdEncoding.EncodeToString(commsg)
	Serialstring := Serial.String()
	fmt.Println("certCA serial =\n",Serialstring)
	/////////////////////////////////////////////////////////////////////
	// Store the certificate with the private key in the database
	/////////////////////////////////////////////////////////////////////
	Insert(db,User{Serialstring,encoding})
	Insertuser(db,User{Serialstring,encoding})
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	}

	//2---certID
func certID_CC(client *channel.Client, Sig_alg string, Password_CA string, Password_ID string, certCASerial string, Country string, Organization string, OrganizationUnit string, Locality string, Province string, StreetAddress string, PostalCode string, CommonName string) {
	db, err := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err != nil {
	fmt.Println(err);
	}
	/////////////////////////////////////////////////////////////////////
	// Get the certificate with the private key from the database
	// and transfer to the channel
	/////////////////////////////////////////////////////////////////////
	combinecertCAstring := Query(db,certCASerial)
	certID_Args := [][]byte{[]byte(Sig_alg), []byte(Password_CA), []byte(Password_ID), 
				[]byte(combinecertCAstring), []byte(Country), []byte(Organization),
				[]byte(OrganizationUnit), []byte(Locality), []byte(Province), 
				[]byte(StreetAddress), []byte(PostalCode), []byte(CommonName)}

	response, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:		"Gen_CertID",
		Args:		certID_Args,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))//call the peer0.org0.example.com

	if err != nil {
		fmt.Printf("Failed to certID: %+v\n", err)
	}
	commsg := response.Payload
	enConBytes := bytes.Split(commsg, []byte(";;;"))
	msg := enConBytes[1]
	certID,_ := unmarshalCert(msg)
	Serial := certID.serialNumber
	db, err1 := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err1 != nil {
		fmt.Println(err1);
	}
	encoding := base64.StdEncoding.EncodeToString(commsg)
	Serialstring := Serial.String()
	fmt.Println("certID serial =\n",Serialstring)
	Insert(db,User{Serialstring,encoding})
	Insertuser(db,User{Serialstring,encoding})
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	}

	//3---certKE
func certKE_CC(client *channel.Client, Chg_sig string, Password_ID string, Password_KE string, certIDSerial string) {
	db, err := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err != nil {
	fmt.Println(err);
	}
	combinecertIDstring := Query(db,certIDSerial)
	certKE_Args := [][]byte{[]byte(Chg_sig), []byte(Password_ID), 
				[]byte(Password_KE), []byte(combinecertIDstring)}
	response, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:		"Gen_CertKE",
		Args:		certKE_Args,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))
	if err != nil {
		fmt.Printf("Failed to certKE: %+v\n", err)
	}
	commsg := response.Payload
	enConBytes := bytes.Split(commsg, []byte(";;;"))
	msg := enConBytes[1]
	certKE,_ := unmarshalCert(msg)
	Serial := certKE.serialNumber
	db, err2 := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err2 != nil {
		fmt.Println(err2);
	}
	encoding := base64.StdEncoding.EncodeToString(commsg)
	Serialstring := Serial.String()
	fmt.Println("certKE serial =\n",Serialstring)
	Insert(db,User{Serialstring,encoding})
	Insertuser(db,User{Serialstring,encoding})
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	}

	//4---verify
func verify_CC(client *channel.Client, newValue string) {
	verifyArgs := [][]byte{[]byte(newValue)}
 
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Verify_Cert",
		Args:        verifyArgs,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))
 
	if err != nil {
	fmt.Printf("Failed to verify: %+v\n", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
	}


	//5---revoke
func revoke_CC(client *channel.Client, Password_Cert string, ID_Cert string) {
	db, err := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err != nil {
	fmt.Println(err);
	}
	combinecertKEstring := Query(db,ID_Cert)
	revokeArgs := [][]byte{[]byte(Password_Cert), []byte(combinecertKEstring)}
	response, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:		"Revoke_Cert",
		Args:		revokeArgs,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))
	if err != nil {
	fmt.Printf("Failed to revoke: %+v\n", err)
	}
	commsg := response.Payload
	enConBytes := bytes.Split(commsg, []byte(";;;"))
	msg := enConBytes[1]
	certCA,_ := unmarshalCert(msg)
	Serial := certCA.serialNumber
	encoding := base64.StdEncoding.EncodeToString(commsg)
	Serialstring := Serial.String()
	fmt.Println("the cert which has been roveked serial = %s \n",Serialstring)
	/////////////////////////////////////////////////////////////////////
	// Update database to consistent with the cert on  chain
	/////////////////////////////////////////////////////////////////////
	Update(db,User{Serialstring,encoding})
	Updateuser(db,User{Serialstring,encoding})
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	
}

	//8---query
func query_CC(client *channel.Client, name string){
	queryArgs := [][]byte{[]byte(name)}
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:		"Query_Cert",
		Args:		queryArgs,
	},channel.WithTargetEndpoints("peer0.org0.example.com"))
 
	if err != nil {
		fmt.Println("Failed to query: ", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
}
