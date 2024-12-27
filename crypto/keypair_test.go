package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	// address := pubKey.Address()

	msg := []byte("hello world!")
	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeyPair_SIgn_Verify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world!")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeyPair_Fail_SIgn_Verify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world!")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	otherprivKey := GeneratePrivateKey()
	otherpubKey := otherprivKey.PublicKey()
	otehrmsg := []byte("xxxxxxxxxxxx")

	assert.False(t, sig.Verify(otherpubKey, msg))
	assert.False(t, sig.Verify(pubKey, otehrmsg))
}
