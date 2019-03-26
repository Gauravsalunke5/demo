package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	BU = "Blockcoderz"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

/*
type student struct {
	PR_no             string `json:"PR_no"`
	First_Name        string `json:"First_Name"`
	Middle_Name       string `json:"Middle_Name "`
	Last_Name         string `json:"Last_Name"`
	College_Name      string `json:"College_Name"`
	Branch            string `json:"Branch"`
	Year_Of_Admission string `json:"Year_Of_Admission"`
	Email_Id          string `json:"Email_Id"`
	Mobile            string `json:"Mobile"`
}
*/
type cert struct {
	PR_no           string `json:"PR_no"`
	Student_Name    string `json:"Student_Name"`
	College_Name    string `json:"College_Name"`
	Seat_no         string `json:"Seat_no"`
	Examination     string `json:"Examination"`
	Year_Of_Passing string `json:"Year_Of_Passing"`
	Sub             string `json:"Sub"`
}

// ===========================
// main function starts up the chaincode in the container during instantiate
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

// ===========================
// Init initializes chaincode
func (t *SimpleChaincode) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// ========================================
// Invoke - Our entry point for Invocations
func (t *SimpleChaincode) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()
	// Handle different functions
	/*
		if function == "addStudent" { //add a Student
			return t.addStudent(APIstub, args)
		} else if function == "readStudent" { //read a Student
			return t.readStudent(APIstub, args)
		} else if function == "addCert" { //add a Certificate
			return t.addCert(APIstub, args)
		} else if function == "readCert" { //read a Certificate
			return t.readCert(APIstub, args)
		} else if function == "transferCert" { //transfer a Certificate
			return t.transferCert(APIstub, args)
		} else if function == "initLedger" {
			return t.initLedger(APIstub, args)
		} else if function == "queryAllCert" {
			return t.queryAllCert(APIstub, args)
		}
	*/

	if function == "initLedger" {
		return t.initLedger(APIstub, args)
	} else if function == "queryAllCert" {
		return t.queryAllCert(APIstub, args)
	} else if function == "addCert" { //add a Certificate
		return t.addCert(APIstub, args)
	} else if function == "readCert" { //read a Certificate
		return t.readCert(APIstub, args)
	} else if function == "transferCert" { //transfer a Certificate
		return t.transferCert(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

// ===============================================
// readcert - read a certificate from chaincode state
func (t *SimpleChaincode) readCert(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	name = args[0]
	valAsbytes, err := APIstub.GetState(name)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Student does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

/*
 * The initLedger method *
Will add test data (10 cert catches)to our network
*/
//PR_no,Student_Name,Seat_no,College_Name,Examination,Year_Of_Passing,Sub
func (t *SimpleChaincode) initLedger(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	Cert := []cert{
		cert{PR_no: "101", Student_Name: "Gaurav1", Seat_no: "1", College_Name: "PCCE", Examination: "may/june", Year_Of_Passing: "2019", Sub: "abc"},
		cert{PR_no: "102", Student_Name: "Gaurav2", Seat_no: "2", College_Name: "PCCE", Examination: "may/june", Year_Of_Passing: "2019", Sub: "abc"},
		cert{PR_no: "103", Student_Name: "Gaurav3", Seat_no: "3", College_Name: "PCCE", Examination: "may/june", Year_Of_Passing: "2019", Sub: "abc"},
		cert{PR_no: "104", Student_Name: "Gaurav4", Seat_no: "4", College_Name: "PCCE", Examination: "may/june", Year_Of_Passing: "2019", Sub: "abc"},
	}

	i := 0
	for i < len(Cert) {
		fmt.Println("i is ", i)
		valAsBytes, _ := json.Marshal(Cert[i])
		APIstub.PutState(strconv.Itoa(i+1), valAsBytes)
		fmt.Println("Added", Cert[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryAllCert(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCert:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// add certificate details
//PR_no,Student_Name,Seat_no,College_Name,Examination,Year_Of_Passing,Sub
func (t *SimpleChaincode) addCert(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1 argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2 argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3 argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4 argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5 argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6 argument must be a non-empty string")
	}

	PRno := args[0]
	CName := args[1]
	Seatno := args[2]
	examination := args[3]
	YOP := args[4]
	sub := args[5]

	// ==== Check if certificate already exists ====
	certAsBytes, err := APIstub.GetState(Seatno)
	if err != nil {
		return shim.Error("Failed to get certificate: " + err.Error())
	} else if certAsBytes != nil {
		return shim.Error("This certificate already exists: " + PRno)
	}

	// ==== Create certificate object and marshal to JSON ====
	cert := &cert{PRno, BU, CName, Seatno, examination, YOP, sub}

	certJSONasBytes, err := json.Marshal(cert)
	err = APIstub.PutState(Seatno, certJSONasBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record Cert catch: %s", Seatno))
	}

	return shim.Success(nil)
}

// ========================================================================
// transferCert - transfer ownership of cert from BlockCoderz to Student
func (t *SimpleChaincode) transferCert(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   0       1
	// "Seatno", "SName"
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	Seatno := args[0]
	SName := args[1]

	certAsBytes, err := APIstub.GetState(Seatno)
	if certAsBytes == nil {
		return shim.Error("Could not locate Cert")
	}
	certToTransfer := cert{}
	json.Unmarshal(certAsBytes, &certToTransfer) //unmarshal it aka JSON.parse()

	certToTransfer.Student_Name = SName //change the owner

	certJSONasBytes, _ := json.Marshal(certToTransfer)
	err = APIstub.PutState(Seatno, certJSONasBytes) //rewrite the certificate
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change Cert holder: %s", Seatno))
	}

	return shim.Success(nil)
}
