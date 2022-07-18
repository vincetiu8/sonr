package motor

import (
	"encoding/json"
	"errors"

	mtr "github.com/sonr-io/sonr/internal/motor"
	apiv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var instance *mtr.MotorNode

func Init(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req apiv1.InitializeRequest
	if err := json.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	// Check if public key provided
	if req.DeviceKeyprintPub == nil {
		// Create Motor instance
		instance = mtr.EmptyMotor(req.DeviceId)

		// Return Initialization Response
		resp := apiv1.InitializeResponse{
			Success: true,
		}
		return json.Marshal(resp)
	}
	return nil, errors.New("Loading existing account not implemented")
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	if res, err := instance.CreateAccount(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
  if instance == nil {
    return nil, errWalletNotExists
  }

  if res, err := instance.Login(buf); err == nil {
    return json.Marshal(res)
  } else {
    return nil, err
  }
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	addr, err := instance.Wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// Balance returns the balance of the wallet.
func Balance() int {
	return int(instance.Balance())
}

// func Connect() error {
// 	if instance == nil {
// 		return errWalletNotExists
// 	}
// 	h, err := host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
// 	if err != nil {
// 		return err
// 	}
// 	instance.host = h
// 	return nil
// }

// DidDoc returns the DID document as JSON
func DidDoc() string {
	if instance == nil {
		return ""
	}
	buf, err := instance.DIDDocument.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}
