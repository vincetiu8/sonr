package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MPCCreate(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")
	_, err = w.PublicKey()
	assert.NoError(t, err, "public key creation succeeds")
}

func Test_MPCDID(t *testing.T) {
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	_, err = w.Address()
	assert.NoError(t, err, "Bech32Address successfully created")
}

func Test_MPCSignMessage(t *testing.T) {
	m := []byte("sign this message")
	w, err := GenerateWallet()
	assert.NoError(t, err, "wallet generation succeeds")

	sig, err := w.Sign(m)
	assert.NoError(t, err, "signing succeeds")
	deserializedSigVerified := sig.Verify(w.Config().PublicPoint(), m)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}