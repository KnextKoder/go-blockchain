Block DS = {
    Header: {
        version: uint32,
        datahash: [32]uint8,
        previousBlockHash: [32]uint8,
        height: uint32,
        timestamp: int64,
    },
    Transaction: {
        data: []byte,
        publicKey: *ecdsa.PublicKey
        signature: *big.Int
    }[],
    Validator: crypto.PublicKey,
    Signature: *crypto.Signature,
    Knowledge : {
        blockHash: [32]uint8
        informationHash?: [32]uint8,
        timestamp?: int64,
        currentValue: float64, // min == 0.0 and max <= 1.0
        previousValue: float64 // min == 0.0 and max < 1.0
    }

    Hash: [32]uint8,
}

