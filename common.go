package types

import "encoding/hex"

type ObjectId [12]byte

func (id ObjectId) Hex() string {
	var buf [24]byte
	hex.Encode(buf[:], id[:])
	return string(buf[:])
}

type Datetime int64
