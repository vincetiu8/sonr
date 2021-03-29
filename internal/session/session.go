package session

import (
	"bytes"
	"strings"

	"encoding/base64"
	"log"

	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	msg "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

type Session struct {
	sender   *md.Peer
	receiver *md.Peer

	incoming *incomingFile
	outgoing *outgoingFile

	callback md.NodeCallback
	filesys  *us.FileSystem
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *us.FileSystem, tc md.NodeCallback) *Session {
	o := newOutgoingFile(req, p)
	return &Session{
		sender:   p,
		receiver: req.To,
		outgoing: o,
		callback: tc,
		filesys:  fs,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *us.FileSystem, tc md.NodeCallback) *Session {
	return &Session{
		sender:   inv.From,
		receiver: p,
		callback: tc,
		filesys:  fs,
		incoming: &incomingFile{
			// Inherited Properties
			properties: inv.Card.Properties,
			payload:    inv.Payload,
			owner:      inv.From.Profile,
			preview:    inv.Card.Preview,
			call:       tc,

			// Builders
			stringsBuilder: new(strings.Builder),
			bytesBuilder:   new(bytes.Buffer),
		},
	}
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, in *incomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				s.callback.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := in.AddBuffer(i, buffer)
			if err != nil {
				s.callback.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := in.Save(s.filesys); err != nil {
					s.callback.Error(err, "HandleIncoming:Save")
				}
				break
			}
			md.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), s.incoming)
}

// ^ write file as Base64 in Msgio to Stream ^ //
func WriteToStream(writer msgio.WriteCloser, s *Session) {
	// Initialize Buffer and Encode File
	buffer := new(bytes.Buffer)
	if err := s.outgoing.encodeFile(buffer); err != nil {
		log.Fatalln(err)
	}

	// Encode Buffer to base 64
	data := buffer.Bytes()
	base := base64.StdEncoding.EncodeToString(data)
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, dat := range ChunkBase64(base) {
		// Create Block Protobuf from Chunk
		chunk := &md.Chunk64{
			Size:  int32(len(dat)),
			Data:  dat,
			Total: total,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
		md.GetState().NeedsWait()
	}

	// Call Completed Sending
	s.callback.Transmitted(s.receiver)
}

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
	chunkSize := K_B64_CHUNK
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}