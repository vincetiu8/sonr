// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: error.proto

package models

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// Error Type from Enum
type ErrorMessage_Type int

const (
	// General
	ErrorMessage_MARSHAL    ErrorMessage_Type = 0
	ErrorMessage_UNMARSHAL  ErrorMessage_Type = 1
	ErrorMessage_IP_RESOLVE ErrorMessage_Type = 2
	ErrorMessage_IP_LOCATE  ErrorMessage_Type = 3
	ErrorMessage_BOOTSTRAP  ErrorMessage_Type = 4
	ErrorMessage_DEVICE_ID  ErrorMessage_Type = 60
	// Topic
	ErrorMessage_TOPIC_JOIN    ErrorMessage_Type = 5
	ErrorMessage_TOPIC_CREATE  ErrorMessage_Type = 6
	ErrorMessage_TOPIC_INVALID ErrorMessage_Type = 7
	ErrorMessage_TOPIC_RPC     ErrorMessage_Type = 8
	ErrorMessage_TOPIC_SUB     ErrorMessage_Type = 9
	ErrorMessage_TOPIC_HANDLER ErrorMessage_Type = 10
	ErrorMessage_TOPIC_LEAVE   ErrorMessage_Type = 11
	ErrorMessage_TOPIC_MESSAGE ErrorMessage_Type = 12
	ErrorMessage_TOPIC_UPDATE  ErrorMessage_Type = 13
	// User
	ErrorMessage_USER_CREATE ErrorMessage_Type = 17
	ErrorMessage_USER_SAVE   ErrorMessage_Type = 18
	ErrorMessage_USER_LOAD   ErrorMessage_Type = 19
	ErrorMessage_USER_FS     ErrorMessage_Type = 20
	ErrorMessage_USER_UPDATE ErrorMessage_Type = 21
	// Transfer
	ErrorMessage_TRANSFER_START ErrorMessage_Type = 14
	ErrorMessage_TRANSFER_CHUNK ErrorMessage_Type = 15
	ErrorMessage_TRANSFER_END   ErrorMessage_Type = 16
	ErrorMessage_INCOMING       ErrorMessage_Type = 22
	ErrorMessage_OUTGOING       ErrorMessage_Type = 23
	ErrorMessage_SESSION        ErrorMessage_Type = 24
	// Host
	ErrorMessage_HOST_DHT    ErrorMessage_Type = 25
	ErrorMessage_HOST_KEY    ErrorMessage_Type = 26
	ErrorMessage_HOST_STREAM ErrorMessage_Type = 27
	ErrorMessage_HOST_START  ErrorMessage_Type = 28
	ErrorMessage_HOST_PUBSUB ErrorMessage_Type = 29
	ErrorMessage_HOST_INFO   ErrorMessage_Type = 37
	ErrorMessage_HOST_MDNS   ErrorMessage_Type = 57
	ErrorMessage_CRYPTO_GEN  ErrorMessage_Type = 30
	// Peer
	ErrorMessage_PEER_NOT_FOUND_INVITE   ErrorMessage_Type = 31
	ErrorMessage_PEER_NOT_FOUND_REPLY    ErrorMessage_Type = 32
	ErrorMessage_PEER_NOT_FOUND_TRANSFER ErrorMessage_Type = 33
	ErrorMessage_PEER_PUBKEY_DECODE      ErrorMessage_Type = 58
	ErrorMessage_PEER_PUBKEY_UNMARSHAL   ErrorMessage_Type = 59
	ErrorMessage_URL_HTTP_GET            ErrorMessage_Type = 34
	ErrorMessage_URL_INFO_RESP           ErrorMessage_Type = 35
	ErrorMessage_FAILED_CONNECTION       ErrorMessage_Type = 36
	// Key Pair
	ErrorMessage_KEY_SET     ErrorMessage_Type = 38
	ErrorMessage_KEY_INVALID ErrorMessage_Type = 39
	ErrorMessage_KEY_ID      ErrorMessage_Type = 40
	// Memory Store
	ErrorMessage_STORE_PUT  ErrorMessage_Type = 41
	ErrorMessage_STORE_GET  ErrorMessage_Type = 42
	ErrorMessage_STORE_FIND ErrorMessage_Type = 43
	ErrorMessage_STORE_INIT ErrorMessage_Type = 44
	// Textile Client
	ErrorMessage_TEXTILE_START_CLIENT ErrorMessage_Type = 45
	ErrorMessage_TEXTILE_USER_CTX     ErrorMessage_Type = 46
	ErrorMessage_TEXTILE_TOKEN_CTX    ErrorMessage_Type = 47
	// Threads Service
	ErrorMessage_THREADS_START_NEW      ErrorMessage_Type = 48
	ErrorMessage_THREADS_START_EXISTING ErrorMessage_Type = 49
	ErrorMessage_THREADS_LIST_ALL       ErrorMessage_Type = 50
	// Mailbox Service
	ErrorMessage_MAILBOX_START_NEW           ErrorMessage_Type = 51
	ErrorMessage_MAILBOX_START_EXISTING      ErrorMessage_Type = 52
	ErrorMessage_MAILBOX_LIST_ALL            ErrorMessage_Type = 53
	ErrorMessage_MAILBOX_MESSAGE_OPEN        ErrorMessage_Type = 54
	ErrorMessage_MAILBOX_MESSAGE_SEND        ErrorMessage_Type = 55
	ErrorMessage_MAILBOX_MESSAGE_PEER_PUBKEY ErrorMessage_Type = 56
	ErrorMessage_MAILBOX_EVENT_STATE         ErrorMessage_Type = 61
	ErrorMessage_MAILBOX_MESSAGE_READ        ErrorMessage_Type = 66
	ErrorMessage_MAILBOX_MESSAGE_UNMARSHAL   ErrorMessage_Type = 67
	ErrorMessage_MAILBOX_MESSAGE_DELETE      ErrorMessage_Type = 68
	ErrorMessage_MAILBOX_ACTION_INVALID      ErrorMessage_Type = 69
	// Push Notifications
	ErrorMessage_PUSH_SINGLE          ErrorMessage_Type = 62
	ErrorMessage_PUSH_MULTIPLE        ErrorMessage_Type = 63
	ErrorMessage_PUSH_START_APP       ErrorMessage_Type = 64
	ErrorMessage_PUSH_START_MESSAGING ErrorMessage_Type = 65
)

var ErrorMessage_Type_name = map[int]string{
	0:  "MARSHAL",
	1:  "UNMARSHAL",
	2:  "IP_RESOLVE",
	3:  "IP_LOCATE",
	4:  "BOOTSTRAP",
	60: "DEVICE_ID",
	5:  "TOPIC_JOIN",
	6:  "TOPIC_CREATE",
	7:  "TOPIC_INVALID",
	8:  "TOPIC_RPC",
	9:  "TOPIC_SUB",
	10: "TOPIC_HANDLER",
	11: "TOPIC_LEAVE",
	12: "TOPIC_MESSAGE",
	13: "TOPIC_UPDATE",
	17: "USER_CREATE",
	18: "USER_SAVE",
	19: "USER_LOAD",
	20: "USER_FS",
	21: "USER_UPDATE",
	14: "TRANSFER_START",
	15: "TRANSFER_CHUNK",
	16: "TRANSFER_END",
	22: "INCOMING",
	23: "OUTGOING",
	24: "SESSION",
	25: "HOST_DHT",
	26: "HOST_KEY",
	27: "HOST_STREAM",
	28: "HOST_START",
	29: "HOST_PUBSUB",
	37: "HOST_INFO",
	57: "HOST_MDNS",
	30: "CRYPTO_GEN",
	31: "PEER_NOT_FOUND_INVITE",
	32: "PEER_NOT_FOUND_REPLY",
	33: "PEER_NOT_FOUND_TRANSFER",
	58: "PEER_PUBKEY_DECODE",
	59: "PEER_PUBKEY_UNMARSHAL",
	34: "URL_HTTP_GET",
	35: "URL_INFO_RESP",
	36: "FAILED_CONNECTION",
	38: "KEY_SET",
	39: "KEY_INVALID",
	40: "KEY_ID",
	41: "STORE_PUT",
	42: "STORE_GET",
	43: "STORE_FIND",
	44: "STORE_INIT",
	45: "TEXTILE_START_CLIENT",
	46: "TEXTILE_USER_CTX",
	47: "TEXTILE_TOKEN_CTX",
	48: "THREADS_START_NEW",
	49: "THREADS_START_EXISTING",
	50: "THREADS_LIST_ALL",
	51: "MAILBOX_START_NEW",
	52: "MAILBOX_START_EXISTING",
	53: "MAILBOX_LIST_ALL",
	54: "MAILBOX_MESSAGE_OPEN",
	55: "MAILBOX_MESSAGE_SEND",
	56: "MAILBOX_MESSAGE_PEER_PUBKEY",
	61: "MAILBOX_EVENT_STATE",
	66: "MAILBOX_MESSAGE_READ",
	67: "MAILBOX_MESSAGE_UNMARSHAL",
	68: "MAILBOX_MESSAGE_DELETE",
	69: "MAILBOX_ACTION_INVALID",
	62: "PUSH_SINGLE",
	63: "PUSH_MULTIPLE",
	64: "PUSH_START_APP",
	65: "PUSH_START_MESSAGING",
}
var ErrorMessage_Type_value = map[string]int{
	"MARSHAL":                     0,
	"UNMARSHAL":                   1,
	"IP_RESOLVE":                  2,
	"IP_LOCATE":                   3,
	"BOOTSTRAP":                   4,
	"DEVICE_ID":                   60,
	"TOPIC_JOIN":                  5,
	"TOPIC_CREATE":                6,
	"TOPIC_INVALID":               7,
	"TOPIC_RPC":                   8,
	"TOPIC_SUB":                   9,
	"TOPIC_HANDLER":               10,
	"TOPIC_LEAVE":                 11,
	"TOPIC_MESSAGE":               12,
	"TOPIC_UPDATE":                13,
	"USER_CREATE":                 17,
	"USER_SAVE":                   18,
	"USER_LOAD":                   19,
	"USER_FS":                     20,
	"USER_UPDATE":                 21,
	"TRANSFER_START":              14,
	"TRANSFER_CHUNK":              15,
	"TRANSFER_END":                16,
	"INCOMING":                    22,
	"OUTGOING":                    23,
	"SESSION":                     24,
	"HOST_DHT":                    25,
	"HOST_KEY":                    26,
	"HOST_STREAM":                 27,
	"HOST_START":                  28,
	"HOST_PUBSUB":                 29,
	"HOST_INFO":                   37,
	"HOST_MDNS":                   57,
	"CRYPTO_GEN":                  30,
	"PEER_NOT_FOUND_INVITE":       31,
	"PEER_NOT_FOUND_REPLY":        32,
	"PEER_NOT_FOUND_TRANSFER":     33,
	"PEER_PUBKEY_DECODE":          58,
	"PEER_PUBKEY_UNMARSHAL":       59,
	"URL_HTTP_GET":                34,
	"URL_INFO_RESP":               35,
	"FAILED_CONNECTION":           36,
	"KEY_SET":                     38,
	"KEY_INVALID":                 39,
	"KEY_ID":                      40,
	"STORE_PUT":                   41,
	"STORE_GET":                   42,
	"STORE_FIND":                  43,
	"STORE_INIT":                  44,
	"TEXTILE_START_CLIENT":        45,
	"TEXTILE_USER_CTX":            46,
	"TEXTILE_TOKEN_CTX":           47,
	"THREADS_START_NEW":           48,
	"THREADS_START_EXISTING":      49,
	"THREADS_LIST_ALL":            50,
	"MAILBOX_START_NEW":           51,
	"MAILBOX_START_EXISTING":      52,
	"MAILBOX_LIST_ALL":            53,
	"MAILBOX_MESSAGE_OPEN":        54,
	"MAILBOX_MESSAGE_SEND":        55,
	"MAILBOX_MESSAGE_PEER_PUBKEY": 56,
	"MAILBOX_EVENT_STATE":         61,
	"MAILBOX_MESSAGE_READ":        66,
	"MAILBOX_MESSAGE_UNMARSHAL":   67,
	"MAILBOX_MESSAGE_DELETE":      68,
	"MAILBOX_ACTION_INVALID":      69,
	"PUSH_SINGLE":                 62,
	"PUSH_MULTIPLE":               63,
	"PUSH_START_APP":              64,
	"PUSH_START_MESSAGING":        65,
}

func (x ErrorMessage_Type) String() string {
	return ErrorMessage_Type_name[int(x)]
}

type ErrorMessage_Severity int

const (
	ErrorMessage_LOG      ErrorMessage_Severity = 0
	ErrorMessage_WARNING  ErrorMessage_Severity = 1
	ErrorMessage_CRITICAL ErrorMessage_Severity = 2
	ErrorMessage_FATAL    ErrorMessage_Severity = 3
)

var ErrorMessage_Severity_name = map[int]string{
	0: "LOG",
	1: "WARNING",
	2: "CRITICAL",
	3: "FATAL",
}
var ErrorMessage_Severity_value = map[string]int{
	"LOG":      0,
	"WARNING":  1,
	"CRITICAL": 2,
	"FATAL":    3,
}

func (x ErrorMessage_Severity) String() string {
	return ErrorMessage_Severity_name[int(x)]
}

// Error Message returned from Core Library
type ErrorMessage struct {
	Type     ErrorMessage_Type
	Severity ErrorMessage_Severity
	Message  string
	Data     string
	Error    string
}

// GetType gets the Type of the ErrorMessage.
func (m *ErrorMessage) GetType() (x ErrorMessage_Type) {
	if m == nil {
		return x
	}
	return m.Type
}

// GetSeverity gets the Severity of the ErrorMessage.
func (m *ErrorMessage) GetSeverity() (x ErrorMessage_Severity) {
	if m == nil {
		return x
	}
	return m.Severity
}

// GetMessage gets the Message of the ErrorMessage.
func (m *ErrorMessage) GetMessage() (x string) {
	if m == nil {
		return x
	}
	return m.Message
}

// GetData gets the Data of the ErrorMessage.
func (m *ErrorMessage) GetData() (x string) {
	if m == nil {
		return x
	}
	return m.Data
}

// GetError gets the Error of the ErrorMessage.
func (m *ErrorMessage) GetError() (x string) {
	if m == nil {
		return x
	}
	return m.Error
}

// MarshalToWriter marshals ErrorMessage to the provided writer.
func (m *ErrorMessage) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if int(m.Type) != 0 {
		writer.WriteEnum(1, int(m.Type))
	}

	if int(m.Severity) != 0 {
		writer.WriteEnum(2, int(m.Severity))
	}

	if len(m.Message) > 0 {
		writer.WriteString(3, m.Message)
	}

	if len(m.Data) > 0 {
		writer.WriteString(4, m.Data)
	}

	if len(m.Error) > 0 {
		writer.WriteString(5, m.Error)
	}

	return
}

// Marshal marshals ErrorMessage to a slice of bytes.
func (m *ErrorMessage) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a ErrorMessage from the provided reader.
func (m *ErrorMessage) UnmarshalFromReader(reader jspb.Reader) *ErrorMessage {
	for reader.Next() {
		if m == nil {
			m = &ErrorMessage{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Type = ErrorMessage_Type(reader.ReadEnum())
		case 2:
			m.Severity = ErrorMessage_Severity(reader.ReadEnum())
		case 3:
			m.Message = reader.ReadString()
		case 4:
			m.Data = reader.ReadString()
		case 5:
			m.Error = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a ErrorMessage from a slice of bytes.
func (m *ErrorMessage) Unmarshal(rawBytes []byte) (*ErrorMessage, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}