package types

import (
	"encoding/hex"
	"fmt"
)

type ObjectId [12]byte
type Datetime int64

func (id ObjectId) Hex() string {
	var buf [24]byte
	hex.Encode(buf[:], id[:])
	return string(buf[:])
}

func (id ObjectId) String() string {
	return fmt.Sprintf("ObjectId(%q)", id.Hex())
}
