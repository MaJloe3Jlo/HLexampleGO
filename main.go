package main

import (
	"log"
	"os"

	"github.com/MaJloe3Jlo/HLexampleGO/blockchain"
	webapp "github.com/MaJloe3Jlo/HLexampleGO/webapp/backend"
	"github.com/MaJloe3Jlo/HLexampleGO/webapp/backend/handlers"
)

func main() {
	// Properties for Hyperledger Fabric.
	settings := blockchain.FabricSettings{
		Orderer:       "orderer.m3.me",
		ChannelID:     "m3channel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/MaJloe3Jlo/HLexampleGO/additional/artifacts/m3.channel.tx",

		CCID:     "HLexampleGO",
		GoPath:   os.Getenv("GOPATH"),
		CCPath:   "github.com/MaJloe3Jlo/HLexampleGO/cc",
		OrgAdmin: "Admin",
		OrgName:  "org1",
		Config:   "config.yaml",

		UserName: "User1",
	}

	if err := settings.Init(); err != nil {
		log.Printf("Init was failed: %s", err)
		return
	}
	defer settings.Close()

	if err := settings.InstallCC(); err != nil {
		log.Printf("Install chaincode was failed: %s", err)
		return
	}

	app := &handlers.App{
		Hyperledger: &settings,
	}
	webapp.Server(app)

}
