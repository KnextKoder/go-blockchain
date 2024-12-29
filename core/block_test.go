package core

import (
	// "fmt"
	"testing"
	"time"

	"github.com/knextkoder/GO-BLOCKCHAIN/crypto"
	"github.com/knextkoder/GO-BLOCKCHAIN/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	Header := &Header{
		Version: 1,
		PrevBLockHash: types.RandomHash(),
		Height: height,
		TImeStamp: time.Now().UnixNano(),
	}

	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(Header, []Transaction{tx})
}

func TestSignBlock (t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

}
func TestVerifyBlock (t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}