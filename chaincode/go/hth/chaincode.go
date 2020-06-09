package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the databasestructure.
type Cad struct {
	Owner         string `json:"owner"`
	Creation_Date string `json:"creationdate"`
	Project_name  string `json:"projectname"`
}

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCad" {
		return s.queryCad(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "create" {
		return s.createCad(APIstub, args)
	} else if function == "queryAll" {
		return s.queryAllCads(APIstub)
	} else if function == "changeOwner" {
		return s.changeCadOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCad(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cadAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(cadAsBytes)
}

/*
 * The creation of a first entry through the interface with the proper type and usign the chaincode interface
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cads := []Cad{
		Cad{Owner: "Neil", Creation_Date: "06/06/19", Project_name: "Test"},
	}

	i := 0
	for i < len(cads) {
		fmt.Println("i is ", i)
		cadAsBytes, _ := json.Marshal(cads[i])
		APIstub.PutState("CAD"+strconv.Itoa(i), cadAsBytes)
		fmt.Println("Added", cads[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCad(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var cad = Cad{Owner: args[1], Creation_Date: args[2], Project_name: args[3]}

	cadAsBytes, _ := json.Marshal(cad)
	APIstub.PutState(args[0], cadAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllCads(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAD0"
	endKey := "CAD999"

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
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCads:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeCadOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cadAsBytes, _ := APIstub.GetState(args[0])
	cad := Cad{}

	json.Unmarshal(cadAsBytes, &cad)
	cad.Owner = args[1]

	cadAsBytes, _ = json.Marshal(cad)
	APIstub.PutState(args[0], cadAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
