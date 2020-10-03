package main

import (
	"encoding/json"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ChaincodeService struct {
}

// Init first information of object.
func (cs *ChaincodeService) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Print("Init chaincode")

	// Get function and arguments.
	function, _ := stub.GetFunctionAndParameters()

	// Check for init function.
	if function != "init" {
		return shim.Error("Function is not init")
	}

	// Marshall json data.
	bytes, err := json.Marshal(Product{
		ID:       1,
		Name:     "Sofa",
		Price:    200,
		Quantity: 1,
	})
	if err != nil {
		shim.Error(err.Error())
	}

	// Put information in ledger by key/value.
	if err := stub.PutState("product", bytes); err != nil {
		shim.Error(err.Error())
	}

	// All success.
	log.Print("Init chaincode was successful")
	return shim.Success(nil)
}

func (cs *ChaincodeService) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("Invoke chaincode")

	function, arguments := stub.GetFunctionAndParameters()

	if function != "invoke" {
		return shim.Error("Function is not invoke or query")
	}

	if len(arguments) < 1 {
		return shim.Error("Quantity of arguments are wrong")
	}

	if arguments[0] == "query" {
		return cs.query(stub, arguments)
	}

	if arguments[0] == "invoke" {
		return cs.invoke(stub, arguments)
	}

	return shim.Error(function)
}

func main() {
	// Starting chaincode and make it ready.
	err := shim.Start(new(ChaincodeService))
	if err != nil {
		log.Fatalf("Cannot start chaincode service: %s", err)
	}
}
