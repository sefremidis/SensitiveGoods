package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Product struct {
	Name   string `json:"name"`
	Size  string `json:"size"`
	UnitPrice string `json:"UnitPrice"`
	MeasurementUnit string `json:"MeasurementUnit"`
	Manufacturer  string `json:"Manufacturer"`
}

type Container struct {
	Forwarder   string `json:"Forwarder"`
}

type Storage struct {
	ProductID   string `json:"ProductID"`
	ContainerID  string `json:"ContainerID"`
	Timestamp string `json:"Timestamp"`
}

type SensorMeasurement struct {
	Timestamp   string `json:"Timestamp"`
	SensorID  string `json:"SensorID"`
	Value string `json:"Value"`
}

type Sensor struct {
	ContainerID   string `json:"ContainerID"`
	SensorType  string `json:"SensorType"`
	SamplingRate string `json:"SamplingRate"`
}

type Delivery struct {
	Sender   string `json:"Sender"`
	Receiver  string `json:"Receiver"`
	Timestamp string `json:"Timestamp"`
}

type User struct {
	Name 	string `json:"Name"`
	Balance string `json:"Balance"`
}

type Transaction struct {
	PayUser 	string `json:"PayUser"`
	ReceiveUser string `json:"ReceiveUser"`
	Amount      string `json:"Amount"`
	Timestamp	string `json:"Timestamp"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "query" {
		return s.query(APIstub, args[0])
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createProduct" {
		return s.createProduct(APIstub, args)
	} else if function == "createContainer" {
		return s.createContainer(APIstub, args)
	} else if function == "createSensor" {
		return s.createSensor(APIstub, args)
	} else if function == "createSensorMeasurement" {
		return s.createSensorMeasurement(APIstub, args)
	} else if function == "createDelivery" {
		return s.createDelivery(APIstub, args)
	} else if function == "createStorage" {
		return s.createStorage(APIstub, args)
	} else if function == "createUser" {
		return s.createUser(APIstub, args)
	} else if function == "queryRange" {
		return s.queryRange(APIstub, args)
	} else if function == "evaluateProduct" {
		return s.evaluateProduct(APIstub, args[0])
	} else if function == "makePayment" {
		return s.makePayment(APIstub, args)
	} else if function == "queryTransactionsByUser" {
		return s.queryTransactionsByUser(APIstub, args[0])
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	products := []Product{
		Product{Name: "wine", Size: "1000", UnitPrice: "10", MeasurementUnit: "bottle", Manufacturer: "Cabernet"},
		Product{Name: "coffee", Size: "10000", UnitPrice: "1", MeasurementUnit: "kilo", Manufacturer: "Nescafe"},
	}
	i := 0
	for i < len(products) {
		productAsBytes, _ := json.Marshal(products[i])
		APIstub.PutState("PRODUCT"+strconv.Itoa(i), productAsBytes)
		i = i + 1
	}

	containers := []Container{
		Container{Forwarder: "COSCO"},
		Container{Forwarder: "MSC"},
	}
	i = 0
	for i < len(containers) {
		containerAsBytes, _ := json.Marshal(containers[i])
		APIstub.PutState("CONTAINER"+strconv.Itoa(i), containerAsBytes)
		i = i + 1
	}

	storings := []Storage{
		Storage{ProductID: "PRODUCT0", ContainerID: "CONTAINER0", Timestamp: "22/7/2019 14:16"},
		Storage{ProductID: "PRODUCT1", ContainerID: "CONTAINER1", Timestamp: "22/7/2019 14:17"},
	}
	i = 0
	for i < len(storings) {
		storingAsBytes, _ := json.Marshal(storings[i])
		APIstub.PutState("STORAGE"+strconv.Itoa(i), storingAsBytes)
		i = i + 1
	}

	sensors := []Sensor{
		Sensor{ContainerID: "CONTAINER0", SensorType: "Temperature", SamplingRate: "100"},
		Sensor{ContainerID: "CONTAINER0", SensorType: "Luminosity", SamplingRate: "300"},
	}
	i = 0
	for i < len(sensors) {
		sensorAsBytes, _ := json.Marshal(sensors[i])
		APIstub.PutState("SENSOR"+strconv.Itoa(i), sensorAsBytes)
		i = i + 1
	}

	sensorMeasurements := []SensorMeasurement{
		SensorMeasurement{Timestamp: "22/7/2019 14:20", SensorID: "SENSOR0", Value: "27"},
		SensorMeasurement{Timestamp: "22/7/2019 14:21", SensorID: "SENSOR1", Value: "100"},
	}
	i = 0
	for i < len(sensorMeasurements) {
		sensorMeasurementAsBytes, _ := json.Marshal(sensorMeasurements[i])
		APIstub.PutState("SENSORMEASUREMENT"+strconv.Itoa(i), sensorMeasurementAsBytes)
		i = i + 1
	}

	deliveries := []Delivery{
		Delivery{Sender: "Cabernet", Receiver: "COSCO", Timestamp: "22/7/2019 14:22"},
		Delivery{Sender: "Nescafe", Receiver: "MSC", Timestamp: "22/7/2019 14:23"},
	}
	i = 0
	for i < len(deliveries) {
		deliveryAsBytes, _ := json.Marshal(deliveries[i])
		APIstub.PutState("DELIVERY"+strconv.Itoa(i), deliveryAsBytes)
		i = i + 1
	}

	users := []User{
		User{Name: "Nescafe", Balance: "1000000"},
		User{Name: "Grigoris", Balance: "1000000"},
	}
	i = 0
	for i < len(users) {
		userAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState("USER"+strconv.Itoa(i), userAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}
///////////////////////////////// QUERY ALL ITEMS //////////////////////////////////////////////

func (s *SmartContract) queryRange(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	resultsIterator, err := APIstub.GetStateByRange(args[0], args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",\n")
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

	return shim.Success(buffer.Bytes())
}


///////////////////////////////////////////////// QUERY ITEM /////////////////////////////////////////////////

func (s *SmartContract) query(APIstub shim.ChaincodeStubInterface, arg string) sc.Response {
	asBytes, _ := APIstub.GetState(arg)
	return shim.Success(asBytes)
}

///////////////////////////////////////// CREATE ITEM ////////////////////////////////////////////////////

func (s *SmartContract) createProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var product = Product{Name: args[0], Size: args[1], UnitPrice: args[2], MeasurementUnit: args[3], Manufacturer: args[4]}

	productAsBytes, _ := json.Marshal(product)

	resultsIterator, err := APIstub.GetStateByRange("PRODUCT0", "PRODUCT9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("PRODUCT"+strconv.Itoa(counter), productAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var container = Container{Forwarder: args[0]}

	containerAsBytes, _ := json.Marshal(container)

	resultsIterator, err := APIstub.GetStateByRange("CONTAINER0", "CONTAINER9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("CONTAINER"+strconv.Itoa(counter), containerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createStorage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var storage = Storage{ProductID: args[0], ContainerID: args[1], Timestamp: args[2]}

	storageAsBytes, _ := json.Marshal(storage)

	resultsIterator, err := APIstub.GetStateByRange("STORAGE0", "STORAGE9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("STORAGE"+strconv.Itoa(counter), storageAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createSensor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var sensor = Sensor{ContainerID: args[0], SensorType: args[1], SamplingRate: args[2]}

	sensorAsBytes, _ := json.Marshal(sensor)

	resultsIterator, err := APIstub.GetStateByRange("SENSOR0", "SENSOR9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("SENSOR"+strconv.Itoa(counter), sensorAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createSensorMeasurement(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var sensorMeasurement = SensorMeasurement{Timestamp: args[0], SensorID: args[1], Value: args[2]}

	sensorMeasurementAsBytes, _ := json.Marshal(sensorMeasurement)

	resultsIterator, err := APIstub.GetStateByRange("SENSORMEASUREMENT0", "SENSORMEASUREMENT9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("SENSORMEASUREMENT"+strconv.Itoa(counter), sensorMeasurementAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createDelivery(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var delivery = Delivery{Sender: args[0], Receiver: args[1], Timestamp: args[2]}

	deliveryAsBytes, _ := json.Marshal(delivery)

	resultsIterator, err := APIstub.GetStateByRange("DELIVERY0", "DELIVERY9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("DELIVERY"+strconv.Itoa(counter), deliveryAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var user = User{Name: args[0], Balance: args[1]}

	userAsBytes, _ := json.Marshal(user)

	resultsIterator, err := APIstub.GetStateByRange("USER0", "USER9")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}

	APIstub.PutState("USER"+strconv.Itoa(counter), userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createTransaction(APIstub shim.ChaincodeStubInterface, args []string) {
	var transaction = Transaction{PayUser: args[0], ReceiveUser: args[1], Amount: args[2]}
	//times:=time.Now()
	transaction.Timestamp="times.String()"

	transactionAsBytes, _ :=json.Marshal(transaction)
	resultsIterator, _ := APIstub.GetStateByRange("TRANSACTION0", "TRANSACTION9")
	defer resultsIterator.Close()
	counter:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		if queryResponse != nil {
			fmt.Printf("Useless code")
		}
		counter=counter+1;
	}
	APIstub.PutState("TRANSACTION"+strconv.Itoa(counter), transactionAsBytes)
}

/////////////////////////////////// OTHER FUNCTIONS //////////////////////////////////////////

func (s *SmartContract) querySensorsByContainer(APIstub shim.ChaincodeStubInterface, arg string) bytes.Buffer {
	startKey := "SENSOR0"
	endKey := "SENSOR9"

	resultsIterator, _ := APIstub.GetStateByRange(startKey, endKey)
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	sensor := Sensor{}
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		json.Unmarshal(queryResponse.Value, &sensor)
		if sensor.ContainerID==arg {
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString(",")
		}
	}
	return buffer
}

func (s *SmartContract) findSensorType(APIstub shim.ChaincodeStubInterface, arg string) string {
	asBytes , _ := APIstub.GetState(arg)
	sensor := Sensor{}
	json.Unmarshal(asBytes, &sensor)

	return sensor.SensorType
}

func (s *SmartContract) querySensorMeasurementsBySensor(APIstub shim.ChaincodeStubInterface, arg string) bytes.Buffer {
	startKey := "SENSORMEASUREMENT0"
	endKey := "SENSORMEASUREMENT9"

	resultsIterator, _ := APIstub.GetStateByRange(startKey, endKey)
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	sensorMeasurement := SensorMeasurement{}
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		json.Unmarshal(queryResponse.Value, &sensorMeasurement)
			if sensorMeasurement.SensorID==arg {
				buffer.WriteString(string(queryResponse.Value))
			}
	}
	return buffer
}

func (s *SmartContract) findProductsToContainer(APIstub shim.ChaincodeStubInterface, arg string) bytes.Buffer {
	resultsIterator, _ := APIstub.GetStateByRange("STORAGE0", "STORAGE9")
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	storage := Storage{}
	i:=0
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		json.Unmarshal(queryResponse.Value, &storage)
		if storage.ContainerID==arg {
			result, _ := APIstub.GetState(storage.ProductID)
			product := Product{}
			json.Unmarshal(result, &product)
			buffer.WriteString(product.Name)
			buffer.WriteString(",")
			i=i+1
		}
	}
	return buffer
}

func (s *SmartContract) evaluateProduct(APIstub shim.ChaincodeStubInterface, arg string) sc.Response {
	sensorsAsBytes := s.querySensorsByContainer(APIstub, arg)
	sensors:=strings.Split(sensorsAsBytes.String(), ",")
	i:=0
	var results bytes.Buffer
	var buffer bytes.Buffer

	productsToSplit := s.findProductsToContainer(APIstub, arg)
	products := strings.Split(productsToSplit.String(), ",")

	percentage:=0
	for i<len(sensors)-1 {
		sensorType := s.findSensorType(APIstub, sensors[i])
		results = s.querySensorMeasurementsBySensor(APIstub, sensors[i])
		measureSplit := strings.Split(results.String(), "\"")
		j:= 11
		k:=0
		var measures [10]int
		for j < len(measureSplit) {
			measures[k], _ =strconv.Atoi(measureSplit[j])
			j=j+12
			k=k+1
		}

		l:=0
		for l<len(products)-1 {
			switch sensorType {
			case "Temperature":
				switch products[l] {
				case "wine":
					m:=0
					for m<k {
						if measures[m] < 20 || measures[m] > 30 {
							buffer.WriteString("Measurement "+strconv.Itoa(m)+" exceeds temperature limit.\n")
							percentage=percentage+1
						}
						m=m+1
					}
					buffer.WriteString("After measuring temperature values: "+strconv.Itoa(percentage)+" percent will be reducted from final payment.\n")
				default:
					buffer.WriteString("Not "+sensorType+" requirements for "+products[l]+"\n")
				}
			case "Luminosity":
				switch products[l] {
				case "wine":
					n:=0
					for n<k {
						if measures[n] < 80 || measures[n] > 100 {
							buffer.WriteString("Measurement "+strconv.Itoa(n)+" exceeds luminosity limit.\n")
							percentage=percentage+1
						}
						n=n+1
					}
					buffer.WriteString("After measuring luminosity values: "+strconv.Itoa(percentage)+" percent will be reducted from final payment.\n")
				default:
					buffer.WriteString("Not "+sensorType+" requirements for "+products[l]+"\n")
				}
			default:
				buffer.WriteString("Sensor type "+sensorType+" cannot be evaluated.\n")

			}
			l=l+1
		}
		i=i+1
	}

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) makePayment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	var buffer bytes.Buffer
	payUser, _ := APIstub.GetState(args[0])
	receiveUser, _ := APIstub.GetState(args[1])
	user1 := User{}
	user2 := User{}

	json.Unmarshal(payUser, &user1)
	json.Unmarshal(receiveUser, &user2)

	payBalance, _ := strconv.Atoi(user1.Balance)
	amount, _ := strconv.Atoi(args[2])
	if amount<=payBalance {
		payBalance=payBalance-amount
		user1.Balance=strconv.Itoa(payBalance)
		receiveBalance, _ :=strconv.Atoi(user2.Balance)
		receiveBalance=receiveBalance+amount
		user2.Balance=strconv.Itoa(receiveBalance)

		payUser, _ = json.Marshal(user1)
		APIstub.PutState(args[0], payUser)
		receiveUser, _ = json.Marshal(user2)
		APIstub.PutState(args[1], receiveUser)
		
		s.createTransaction(APIstub, args)

		buffer.WriteString("Transaction succesfully submitted")
	} else{
		buffer.WriteString("Transaction cannot be implemented, not enough money in User")
	}

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryTransactionsByUser(APIstub shim.ChaincodeStubInterface, arg string) sc.Response {
	resultsIterator, _ := APIstub.GetStateByRange("TRANSACTION0", "TRANSACTION9")
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	transaction := Transaction{}
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		json.Unmarshal(queryResponse.Value, &transaction)
		if transaction.PayUser==arg || transaction.ReceiveUser==arg {
			buffer.WriteString(queryResponse.Key+":")
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString(",")
		}
	}
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

