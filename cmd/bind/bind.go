package bind

import (
	"log"

	md "github.com/sonr-io/core/pkg/models"
	sn "github.com/sonr-io/core/pkg/node"
	"google.golang.org/protobuf/proto"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host has Bootstrapped
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnRemoteStart(data []byte) // User started remote
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// * Struct: Reference for Binded Proxy Node * //
type MobileNode struct {
	ID              string
	DeviceID        string
	UserID          uint32
	node            *sn.Node
	hasStarted      bool
	hasBootstrapped bool
	hostOpts        *sn.HostOptions
}

// @ Create New Mobile Node
func NewNode(reqBytes []byte, call Callback) *MobileNode {
	// Unmarshal Request
	req := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Fatalln(err)
	}

	hostOpts, err := sn.NewHostOpts(&req)
	if err != nil {
		log.Println(err)
	}

	// Create New Sonr Client
	node := sn.NewNode(&req, call)

	// Return Mobile Node
	return &MobileNode{
		node:            node,
		hasStarted:      false,
		hasBootstrapped: false,
		hostOpts:        hostOpts,
	}
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host
func (mn *MobileNode) Connect() {
	// Start Node
	result := mn.node.Start(mn.hostOpts)
	if result {
		// Set Peer Info
		peer := mn.node.Peer()
		mn.ID = peer.Id.Peer
		mn.DeviceID = peer.Id.Device
		mn.UserID = peer.Id.User

		// Set Started
		mn.hasStarted = true

		// Bootstrap to Peers
		strapResult := mn.node.Bootstrap(mn.hostOpts)
		if strapResult {
			mn.hasBootstrapped = true
		} else {
			log.Println("Failed to bootstrap node")
		}
	} else {
		log.Println("Failed to start host")
	}
}

// @ Return URL Metadata, Helper Method
func GetURLMetadata(url string) []byte {
	// Get Link Data
	data, err := md.GetPageInfoFromUrl(url)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}

	// Marshal
	bytes, err := proto.Marshal(data)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}
	return bytes
}

// **--------------** //
// ** Node Actions ** //
// **--------------** //
// @ Update proximity/direction and Notify Lobby
func (mn *MobileNode) Update(facing float64, heading float64) {
	if mn.IsReady() {
		mn.node.Update(facing, heading)
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.IsReady() {
		mn.node.Message(msg, to)
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.IsReady() {
		// Initialize from Request
		req := &md.InviteRequest{}
		err := proto.Unmarshal(reqBytes, req)
		if err != nil {
			log.Println(err)
		}
		mn.node.Invite(req)
	}
}

// @ Create a New Group
func (mn *MobileNode) CreateGroup() string {
	if mn.IsReady() {
		w := mn.node.CreateGroup()
		return w
	}
	return ""
}

// @ Join Existing Group
func (mn *MobileNode) JoinGroup(data string) {
	if mn.IsReady() {
		mn.node.JoinGroup(data)
	}
}

// @ Respond to an Invitation
func (mn *MobileNode) Respond(decision bool) {
	if mn.IsReady() {
		mn.node.Respond(decision)
	}
}

// @ Join Existing Group
func (mn *MobileNode) JoinRemote(data string) {
	if mn.IsReady() {
		mn.node.JoinRemote(data)
	}
}

// ** User Actions ** //
// @ Info returns ALL Peer Data as Bytes
func (mn *MobileNode) Info() []byte {
	if mn.IsReady() {
		info := mn.node.Info()
		return info
	}
	return nil
}

// @ Link with a QR Code
func (mn *MobileNode) LinkDevice(json string) {
	if mn.IsReady() {
		mn.node.LinkDevice(json)
	}
}

// @ Updates Current Contact Card
func (mn *MobileNode) SetContact(conBytes []byte) {
	if mn.IsReady() {
		// Unmarshal Data
		newContact := &md.Contact{}
		err := proto.Unmarshal(conBytes, newContact)
		if err != nil {
			log.Println(err)
		}
		mn.node.SetContact(newContact)
	}
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
// @ Checks for is Ready
func (mn *MobileNode) IsReady() bool {
	return mn.hasBootstrapped && mn.hasStarted
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Pause() {
	mn.node.Pause()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Resume() {
	mn.node.Resume()
}

// @ Close Ends All Network Communication
func (mn *MobileNode) Stop() {
	mn.node.Stop()
}