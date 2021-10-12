package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the databasestructure.
type Meta struct {
	Creator      string `json:"creator"`
	Contributor  string `json:"contributor"`
	Publisher    string `json:"publisher"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Language     string `json:"language"`
	Format       string `json:"format"`
	Subject      string `json:"subject"`
	Description  string `json:"descrition"`
	Relation     string `json:"relation"`
	Source       string `json:"source"`
	Type         string `json:"type"`
	Coverage     string `json:"coverage"`
	Rights       string `json:"rights"`
	Measurements string `json:"measurements"`
	Images       string `json:"images"`
	Resolution   string `json:"resolution"`
	Version      string `json:"version"`
}

//Function to check if metasbyte is zero
func allZero(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

//Function to check if an element is contained
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
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

func (s *SmartContract) querySingle(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	metaAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(metaAsBytes)
}

//The chaincode needs to be initialize with a random entry but it must be complete
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	metas := []Meta{
		Meta{Creator: "Pavel Vazquez", Contributor: "HTH", Publisher: "Pavel Vazquez", Title: "Test", Date: "01/01/20", Language: "en", Format: ".csv", Subject: "Biomaterials", Description: "Dataset", Relation: "None", Source: "None", Type: "Dataset", Coverage: "EU", Rights: "Creative Commons Attribution 4.0 International License", Measurements: "3", Images: "5", Resolution: "100", Version: "1"},
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

//Validation area
func (s *SmartContract) create(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var ver int
	if len(args) != 18 {
		return shim.Error("Incorrect number of metadata entries")
	}

	for i := 1; i < 99; i++ {
		ver = i
		metaAsBytes, _ := APIstub.GetState(args[0] + strconv.Itoa(ver))
		if allZero(metaAsBytes) {
			break
		}
	}

	var meta = Meta{Creator: args[1], Contributor: args[2], Publisher: args[3], Title: args[4], Date: args[5], Language: args[6], Format: args[7], Subject: args[8], Description: args[9], Relation: args[10], Source: args[11], Type: args[12], Coverage: args[13], Rights: args[14], Measurements: args[15], Images: args[16], Resolution: args[17], Version: strconv.Itoa(ver)}
	metaAsBytes, _ := json.Marshal(meta)
	APIstub.PutState(args[0]+strconv.Itoa(ver), metaAsBytes)

	//Date validation
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if re.MatchString(args[5]) == false {
		return shim.Error("Invalid date format. dd/mm/yyyy")
	}

	//Language validation
	lang := []string{"aa", "ab", "ae", "af", "ak", "am", "an", "ar", "as", "av", "ay", "az", "ba", "be", "bg", "bh", "bi", "bm", "bn", "bo", "br", "bs", "ca", "ce", "ch", "co", "cr", "cs", "cu", "cv", "cy", "da", "de", "dv", "dz", "ee", "el", "en", "eo", "es", "et", "eu", "fa", "ff", "fi", "fj", "fo", "fr", "fy", "ga", "gd", "gl", "gn", "gu", "gv", "ha", "he", "hi", "ho", "hr", "ht", "hu", "hy", "hz", "ia", "id", "ie", "ig", "ii", "ik", "io", "is", "it", "iu", "ja", "jv", "ka", "kg", "ki", "kj", "kk", "kl", "km", "kn", "ko", "kr", "ks", "ku", "kv", "kw", "ky", "la", "lb", "lg", "li", "ln", "lo", "lt", "lu", "lv", "mg", "mh", "mi", "mk", "ml", "mn", "mr", "ms", "mt", "my", "na", "nb", "nd", "ne", "ng", "nl", "nn", "no", "nr", "nv", "ny", "oc", "oj", "om", "or", "os", "pa", "pi", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "rw", "sa", "sc", "sd", "se", "sg", "si", "sk", "sl", "sm", "sn", "so", "sq", "sr", "ss", "st", "su", "sv", "sw", "ta", "te", "tg", "th", "ti", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ty", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "wo", "xh", "yi", "yo", "za", "zh", "zu"}

	if contains(lang, args[6]) == false {
		return shim.Error("Invalid language format. Not an ISO 639-1 standard (e.g EN)")
	}

	//Country validation
	country := []string{"EU", "AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"}
	if contains(country, args[13]) == false {
		return shim.Error("Invalid coverage format. Not an ISO 639-1 standard ISO-3166-1 (e.g NO)")
	}

	//Type validation
	types := []string{"collection", "dataset", "event", "image", "interactiveresource", "movingimage", "physicalobject", "service", "software", "sound", "stillimage", "text"}
	if contains(types, args[12]) == false {
		return shim.Error("Invalid type format. Not a DCMI Type Vocabulary (e.g Dataset)")
	}
	return shim.Success(nil)
}

func (s *SmartContract) queryAll(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := ""
	endKey := ""

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

//Query a single project
func (s *SmartContract) query(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	startKey := args[0] + "1"
	endKey := args[0] + "99"

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
