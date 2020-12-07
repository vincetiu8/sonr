package transfer

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Data []byte
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Data []byte
}

// Service Struct
type AuthService struct {
	// Current Data
	onInvite  OnProtobuf
	peerConn  *PeerConnection
	respCh    chan *md.AuthMessage
	inviteMsg *md.AuthMessage
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)

	// Send Callback
	as.onInvite(args.Data)

	// Received Message
	receivedMessage := md.AuthMessage{}
	err := proto.Unmarshal(args.Data, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message
	as.inviteMsg = &receivedMessage

	select {
	// Received Auth Channel Message
	case m := <-as.respCh:
		log.Println("User has replied")

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			log.Println(err)
		}

		// Set Message data and call done
		reply.Data = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ^ Send SendInvite to a Peer ^ //
func (pc *PeerConnection) SendInvite(h host.Host, id peer.ID, msgBytes []byte) {
	// Initialize Data
	rpcClient := gorpc.NewClient(h, protocol.ID("/sonr/rpc/auth"))
	var reply AuthReply
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "Invited", args, &reply, done)

	// Initiate Call on transfer
	call := <-done
	if call.Error != nil {
		// Send Error
		onError(err, "sendInvite")
		log.Panicln(err)
	}

	// Send Callback and Reset
	pc.respondedCall(reply.Data)

	// Received Message
	responseMessage := md.AuthMessage{}
	err = proto.Unmarshal(reply.Data, &responseMessage)
	if err != nil {
		// Send Error
		onError(err, "Unmarshal")
		log.Panicln(err)
	}

	// Check Response for Accept
	if responseMessage.Event == md.AuthMessage_ACCEPT {
		// Begin Transfer
		pc.SendFile(h, id, responseMessage.From)
	}
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) SendResponse(decision bool, peer *md.Peer) {
	// @ Check Decision
	if decision {
		// Initialize Transfer
		currMsg := pc.auth.inviteMsg
		pc.transfer = pc.NewTransfer(currMsg.Metadata, currMsg.From)

		// Create Accept Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_ACCEPT,
		}

		// Send to Channel
		pc.auth.respCh <- respMsg

	} else {
		// Create Decline Response
		respMsg := &md.AuthMessage{
			From:  peer,
			Event: md.AuthMessage_DECLINE,
		}

		// Send to Channel
		pc.auth.respCh <- respMsg
	}
}