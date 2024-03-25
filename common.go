package types

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type ObjectId [12]byte
type Datetime int64

var NilObjectID ObjectId
var ErrInvalidHex = errors.New("the provided hex string is not a valid ObjectId")

func (id ObjectId) Hex() string {
	var buf [24]byte
	hex.Encode(buf[:], id[:])
	return string(buf[:])
}

func (id ObjectId) String() string {
	return fmt.Sprintf("ObjectId(%q)", id.Hex())
}

func ObjectIdFromHex(s string) (ObjectId, error) {
	if len(s) != 24 {
		return NilObjectID, ErrInvalidHex
	}

	var oid [12]byte
	_, err := hex.Decode(oid[:], []byte(s))
	if err != nil {
		return NilObjectID, err
	}

	return oid, nil
}
