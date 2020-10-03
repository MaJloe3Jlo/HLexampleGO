package blockchain

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

func (fs *FabricSettings) ProductQuery() (string, error) {
	var arguments []string
	arguments = append(arguments, "invoke")
	arguments = append(arguments, "query")
	arguments = append(arguments, "product")

	response, err := fs.client.Query(channel.Request{
		ChaincodeID: fs.CCID,
		Fcn:         arguments[0],
		Args:        [][]byte{[]byte(arguments[1]), []byte(arguments[2])},
	})
	if err != nil {
		return "", err
	}

	return string(response.Payload), nil
}
