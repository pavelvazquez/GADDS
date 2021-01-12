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
type Meta struct {
	Contributor string `json:"contributor"`
	Coverage    string `json:"coverage"`
	Creator     string `json:"creator"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Format      string `json:"format"`
	//Identifier  string `json:"identifier"`
	Language  string `json:"language"`
	Publisher string `json:"publisher"`
	Relation  string `json:"relation"`
	Rights    string `json:"rights"`
	Source    string `json:"source"`
	Title     string `json:"title"`
	Type      string `json:"type"`
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
	if function == "query" {
		return s.query(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "create" {
		return s.create(APIstub, args)
	} else if function == "queryAll" {
		return s.queryAll(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) query(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	metaAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(metaAsBytes)
}

//The chaincode need to be initialize with a random entry but it must be complete
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	metas := []Meta{
		Meta{Contributor: "Pavel", Coverage: "EU", Creator: "Pavel", Date: "01/01/20", Description: "Test", Format: "none", Language: "EN", Publisher: "Pavel", Relation: "none", Rights: "none", Source: "none", Title: "Test", Type: "none"},
	}

	i := 0
	for i < len(metas) {
		fmt.Println("i is ", i)
		metaAsBytes, _ := json.Marshal(metas[i])
		APIstub.PutState("META"+strconv.Itoa(i), metaAsBytes)
		fmt.Println("Added", metas[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) create(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	var meta = Meta{Contributor: args[1], Coverage: args[2], Creator: args[3], Date: args[4], Description: args[5], Format: args[6], Language: args[7], Publisher: args[8], Relation: args[9], Rights: args[10], Source: args[11], Title: args[12], Type: args[13]}

	metaAsBytes, _ := json.Marshal(meta)
	APIstub.PutState(args[0], metaAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAll(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "META0"
	endKey := "META999"

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

	fmt.Printf("- queryAll:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
