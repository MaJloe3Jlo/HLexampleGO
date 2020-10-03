package blockchain

import (
	"log"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (fs *FabricSettings) ProductInvoke(value string) (string, error) {

	arguments := []string{"invoke", "invoke", "product", value}
	eventID := "notificationOfInvoke"

	// Description of invoke.
	transitData := make(map[string][]byte)
	transitData["result"] = []byte("Data of invoke")

	regEvent, notify, err := fs.event.RegisterChaincodeEvent(fs.CCID, eventID)
	if err != nil {
		return "", err
	}
	defer fs.event.Unregister(regEvent)

	response, err := fs.client.Execute(channel.Request{ChaincodeID: fs.CCID, Fcn: arguments[0], Args: [][]byte{[]byte(arguments[1]), []byte(arguments[2]), []byte(arguments[3])}, TransientMap: transitData})
	if err != nil {
		return "", err
	}

	select {
	case event := <-notify:
		log.Printf("Recieved event: %s", event)
	case <-time.After(time.Second * 30):
		return "", err
	}

	return string(response.TransactionID), nil
}
