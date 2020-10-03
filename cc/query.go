package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (cs *ChaincodeService) query(stub shim.ChaincodeStubInterface, arguments []string) pb.Response {
	if len(arguments) < 2 {
		shim.Error("Arguments are missed")
	}

	if arguments[1] == "product" {
		state, err := stub.GetState("product")
		if err != nil {
			return shim.Error("Cannot get state of product")
		}

		return shim.Success(state)
	}

	return shim.Error("Query is failed. Check the parameters.")
}
