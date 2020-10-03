package blockchain

import (
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
)

// Struct for hyperledger fabric settings.
type FabricSettings struct {
	Config        string
	OrgID         string
	Orderer       string
	ChannelID     string
	CCID          string
	init          bool
	ChannelConfig string
	GoPath        string
	CCPath        string
	OrgAdmin      string
	OrgName       string
	UserName      string
	client        *channel.Client
	admin         *resmgmt.Client
	sdk           *fabsdk.FabricSDK
	event         *event.Client
}

func (fs *FabricSettings) Init() error {
	// Check for already inited.
	if fs.init {
		return errors.New("already initialized")
	}

	// Load configuration from file.
	sdk, err := fabsdk.New(config.FromFile(fs.Config))
	if err != nil {
		return errors.WithMessage(err, " cannot create SDK")
	}
	fs.sdk = sdk
	log.Print("SDK was created")

	// Create resources for managment client.
	clientResourceManagerCtx := fs.sdk.Context(fabsdk.WithUser(fs.OrgAdmin), fabsdk.WithOrg(fs.OrgName))
	if err != nil {
		return errors.WithMessage(err, " failed to load admin settings")
	}
	resources, err := resmgmt.New(clientResourceManagerCtx)
	if err != nil {
		return errors.WithMessage(err, " failed to create channel manager from admin settings")
	}
	fs.admin = resources
	log.Println("Resources for managment was created")

	// Creating MSP clients.
	mspC, err := mspclient.New(sdk.Context(), mspclient.WithOrg(fs.OrgName))
	if err != nil {
		return errors.WithMessage(err, " failed to create MSP")
	}
	identity, err := mspC.GetSigningIdentity(fs.OrgAdmin)
	if err != nil {
		return errors.WithMessage(err, " failed to get sign for admin")
	}

	// Save channel.
	request := resmgmt.SaveChannelRequest{
		ChannelID:         fs.ChannelID,
		ChannelConfigPath: fs.ChannelConfig,
		SigningIdentities: []msp.SigningIdentity{identity}}
	trID, err := fs.admin.SaveChannel(request, resmgmt.WithOrdererEndpoint(fs.Orderer))
	if err != nil || trID.TransactionID == "" {
		return errors.WithMessage(err, " channel was not saved")
	}
	log.Println("Channel was created")

	// Join channel.
	if err = fs.admin.JoinChannel(fs.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(fs.Orderer)); err != nil {
		return errors.WithMessage(err, " cannot join channel")
	}
	log.Println("Channel was joined")

	log.Println("Init was successful")
	fs.init = true
	return nil
}

func (fs *FabricSettings) InstallCC() error {
	// Create chaincode package.
	ccPkg, err := packager.NewCCPackage(fs.CCPath, fs.GoPath)
	if err != nil {
		return errors.WithMessage(err, " cannot create chaincode package")
	}
	log.Println("Chaincode package was created")

	// Install example chaincode.
	installRequest := resmgmt.InstallCCRequest{
		Name:    fs.CCID,
		Path:    fs.CCPath,
		Version: "1",
		Package: ccPkg,
	}
	if _, err := fs.admin.InstallCC(installRequest, resmgmt.WithRetry(retry.DefaultResMgmtOpts)); err != nil {
		return errors.WithMessage(err, " cannot install chaincode")
	}

	// Install policy.
	policy := cauthdsl.SignedByAnyMember([]string{"org1.m3.me"})

	response, err := fs.admin.InstantiateCC(fs.ChannelID, resmgmt.InstantiateCCRequest{Name: fs.CCID, Path: fs.GoPath, Version: "1", Args: [][]byte{[]byte("init")}, Policy: policy})
	if err != nil || response.TransactionID == "" {
		return errors.WithMessage(err, " cannot to instantiate chaincode")
	}
	log.Println("Chaincode was instantiated")

	// Create client channel context.
	clientCtx := fs.sdk.ChannelContext(fs.ChannelID, fabsdk.WithUser(fs.UserName))
	fs.client, err = channel.New(clientCtx)
	if err != nil {
		return errors.WithMessage(err, " cannot create channel client")
	}
	log.Println("Client channel was created")

	// Make an access to events.
	fs.event, err = event.New(clientCtx)
	if err != nil {
		return errors.WithMessage(err, " cannot create new event")
	}
	log.Println("Client created")

	log.Println("Install chaincode was successful")
	return nil
}

func (fs *FabricSettings) Close() {
	fs.sdk.Close()
}
