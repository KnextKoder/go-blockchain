package core

import "io"

// import (
// 	"encoding/gob"
// 	"io"
// )

//
// For now we GOB encoding is used for fast bootstrapping of the project
// in a later phase I'm considering using Protobuffers as default encoding / decoding.
//

// type Encoder[T any] interface {
// 	Encode(T) error
// }

// type Decoder[T any] interface {
// 	Decode(T) error
// }

type Encoder[T any] interface {
	Encode(io.Writer, T) error
}

type Decoder[T any] interface {
	Decode(io.Reader, T) error
}