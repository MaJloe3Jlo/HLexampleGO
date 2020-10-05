package main

import (
	"encoding/json"

	"github.com/MaJloe3Jlo/HLexampleGO/cc/model"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (cs *ChaincodeService) queryAll(stub shim.ChaincodeStubInterface, arguments []string) pb.Response {
	if len(arguments) < 2 {
		shim.Error("Arguments are missed")
	}

	if arguments[1] == "product" {
		state, err := stub.GetHistoryForKey("product")
		if err != nil {
			shim.Error("State is not getted")
		}
		defer state.Close()
		result := []model.Product{}

		for state.HasNext() {
			kv, err := state.Next()
			if err != nil {
				shim.Error("Has no next state")
			}
			var product model.Product
			if err := json.Unmarshal(kv.Value, &product); err != nil {
				shim.Error("Cannot unmarshal value")
			}
			result = append(result, product)
		}

		var bytes []byte

		bytes, err = json.Marshal(result)
		if err != nil {
			return shim.Error("Cannot make an payload of result")
		}

		return shim.Success(bytes)
	}

	return shim.Error("Query is failed. Check the parameters.")
}
