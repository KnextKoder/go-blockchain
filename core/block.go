package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/knextkoder/GO-BLOCKCHAIN/crypto"
	"github.com/knextkoder/GO-BLOCKCHAIN/types"
)

type Header struct {
	Version   uint32
	DataHash  types.Hash
	PrevBLockHash types.Hash
	Height 	  uint32
	TImeStamp int64
}

// type InformationHash struct {
// 	Value types.Hash
// 	Valid bool
// }

// type Knowledge struct {
// 	blockHash types.Hash
// 	*InformationHash
// 	timestamp int64
// 	currentValue float64 // min == 0.0 and max <= 1.0
// 	previousValue float64 // min == 0.0 and max < 1.0
// }

type Block struct {
	*Header
	Transactions []Transaction
	Validator  crypto.PublicKey
	Signature  *crypto.Signature
	
	// *Knowledge
	// Cached version of the header hash
	hash types.Hash
}

// func NewSmartBlock(h *Header, txx []Transaction, k *Knowledge) *Block {
// 	return &Block{
// 		Header: h, 
// 		Transactions: txx,
// 		Knowledge: k,
// 	}
// }

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header: h,
		Transactions: txx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil{
        return err
    }

	b.Validator = privKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}

	// for _, tx := range b.Transactions {
	// 	if err := tx.Verify(); err != nil {
	// 		return err
	// 	}
	// }

	// dataHash, err := CalculateDataHash(b.Transactions)
	// if err != nil {
	// 	return err
	// }

	// if dataHash != b.DataHash {
	// 	return fmt.Errorf("block (%s) has an invalid data hash", b.Hash(BlockHasher{}))
	// }

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {

	if b.hash.IsZero(){
		b.hash = hasher.Hash(b)
	}
	return b.hash
}


func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}