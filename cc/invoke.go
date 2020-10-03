package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (cs *ChaincodeService) invoke(stub shim.ChaincodeStubInterface, arguments []string) pb.Response {
	if len(arguments) < 2 {
		shim.Error("Arguments are missed")
	}

	if arguments[1] == "product" && len(arguments) == 3 {
		if err := stub.PutState("product", []byte(arguments[2])); err != nil {
			shim.Error("Error to put state in ledger")
		}

		if err := stub.SetEvent("notificationOfInvoke", []byte("some notification message")); err != nil {
			shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}
