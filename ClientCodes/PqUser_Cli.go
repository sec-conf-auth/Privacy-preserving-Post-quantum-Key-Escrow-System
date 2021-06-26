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
	_"bytes"
	_"io"
	_"encoding/base64"
	_"crypto/x509"
	_"crypto/x509/pkix"

	_"crypto/sha256"
	_"math/big"
	_"crypto/rand"

	"time"
	_"encoding/pem"
	_"encoding/gob"
 
//	"encoding/hex"
//	"encoding/json"
//	"encoding/asn1"
)
	/////////////////////////////////////////////////////////////////////
	// Channel related information for the client
	/////////////////////////////////////////////////////////////////////
var (
        cc          = "PqUser"
        user        = "Admin" 
        secret      = ""
        channelName = "mychannel"
        lvl         = logging.INFO
        orgName     = "Org1MSP"
	)

	/////////////////////////////////////////////////////////////////////
	// The client codes for user to query its private database
	/////////////////////////////////////////////////////////////////////
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
	// The client codes for user to insert its private database
	/////////////////////////////////////////////////////////////////////
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
	// The client codes for user to update its private database
	/////////////////////////////////////////////////////////////////////
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
	/////////////////////////////////////////////////////////////////////
	// Choose the function of Chaincode
	/////////////////////////////////////////////////////////////////////
	if num =="Verify_Cert"{		
		newValue := args[2]
		verify_CC(client, newValue )		
	}
	if num =="Encap"{		
		MyID_KE := args[2]
		OpID_KE := args[3]
		encap_CC(client, MyID_KE, OpID_KE)		
	}
	if num =="Decap"{		
		MyPassword_KE := args[2]
		certKESerial := args[3]
		ID_Encap := args[4]
		decap_CC(client, MyPassword_KE, certKESerial, ID_Encap )		
	}		
	if num =="Query_Cert"{		
		name := args[2]
		query_CC(client, name)		
	}		
	elapsedTime := time.Since(startTime)
	fmt.Println("The query time is ",elapsedTime)	
	}

	/////////////////////////////////////////////////////////////////////
	// Here are the fuctions which invoke the corresponding chaincode APIs
	/////////////////////////////////////////////////////////////////////
	//1---Verify_Cert
func verify_CC(client *channel.Client, newValue string) {
	verifyArgs := [][]byte{[]byte(newValue)}
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Verify_Cert",
		Args:        verifyArgs,
	},channel.WithTargetEndpoints("peer0.org1.example.com"))
	if err != nil {
		fmt.Printf("Failed to verify: %+v\n", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
	}

	//2---Encap
func encap_CC(client *channel.Client, MyID_KE string, OpID_KE string) {
	encapArgs := [][]byte{[]byte(MyID_KE), []byte(OpID_KE)}
 
	response, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Encap",
		Args:        encapArgs,
	},channel.WithTargetEndpoints("peer0.org1.example.com"))
 
	if err != nil {
		fmt.Printf("Failed to encap: %+v\n", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
	}

	//3---Decap
func decap_CC(client *channel.Client, MyPassword_KE string, certKESerial string, ID_Encap string) {
	db, err := sql.Open("mysql", "root:hzaucoi@tcp(127.0.0.1:3306)/cert?charset=utf8");
	if err != nil {
	fmt.Println(err);
	}
	combinecertKEstring := Queryuser(db,certKESerial)
	decapArgs := [][]byte{[]byte(MyPassword_KE), []byte(combinecertKEstring), []byte(ID_Encap)}
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Decap",
		Args:        decapArgs,
	},channel.WithTargetEndpoints("peer0.org1.example.com"))
	if err != nil {
		fmt.Printf("Failed to encap: %+v\n", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
	}

//4---Query_Cert
func query_CC(client *channel.Client, name string){
	queryArgs := [][]byte{[]byte(name)}
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:         "Query_Cert",
		Args:        queryArgs,
	},channel.WithTargetEndpoints("peer0.org1.example.com"))
	if err != nil {
		fmt.Println("Failed to query: ", err)
	}
	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
        //return ret
	}
