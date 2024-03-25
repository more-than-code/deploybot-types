package types

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

type ObjectId [12]byte

// ErrInvalidHex indicates that a hex string cannot be converted to an ObjectID.
var ErrInvalidHex = errors.New("the provided hex string is not a valid ObjectID")

var NilObjectId ObjectId

func (id ObjectId) Hex() string {
	var buf [24]byte
	hex.Encode(buf[:], id[:])
	return string(buf[:])
}

// ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a
// valid ObjectID.
func ObjectIDFromHex(s string) (ObjectId, error) {
	if len(s) != 24 {
		return NilObjectId, ErrInvalidHex
	}

	var oid [12]byte
	_, err := hex.Decode(oid[:], []byte(s))
	if err != nil {
		return NilObjectId, err
	}

	return oid, nil
}

// MarshalText returns the ObjectID as UTF-8-encoded text. Implementing this allows us to use ObjectID
// as a map key when marshalling JSON. See https://pkg.go.dev/encoding#TextMarshaler
func (id ObjectId) MarshalText() ([]byte, error) {
	return []byte(id.Hex()), nil
}

// UnmarshalText populates the byte slice with the ObjectID. Implementing this allows us to use ObjectID
// as a map key when unmarshalling JSON. See https://pkg.go.dev/encoding#TextUnmarshaler
func (id *ObjectId) UnmarshalText(b []byte) error {
	oid, err := ObjectIDFromHex(string(b))
	if err != nil {
		return err
	}
	*id = oid
	return nil
}

// MarshalJSON returns the ObjectID as a string
func (id ObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}

// UnmarshalJSON populates the byte slice with the ObjectID. If the byte slice is 24 bytes long, it
// will be populated with the hex representation of the ObjectID. If the byte slice is twelve bytes
// long, it will be populated with the BSON representation of the ObjectID. This method also accepts empty strings and
// decodes them as NilObjectID. For any other inputs, an error will be returned.
func (id *ObjectId) UnmarshalJSON(b []byte) error {
	// Ignore "null" to keep parity with the standard library. Decoding a JSON null into a non-pointer ObjectID field
	// will leave the field unchanged. For pointer values, encoding/json will set the pointer to nil and will not
	// enter the UnmarshalJSON hook.
	if string(b) == "null" {
		return nil
	}

	var err error
	switch len(b) {
	case 12:
		copy(id[:], b)
	default:
		// Extended JSON
		var res interface{}
		err := json.Unmarshal(b, &res)
		if err != nil {
			return err
		}
		str, ok := res.(string)
		if !ok {
			m, ok := res.(map[string]interface{})
			if !ok {
				return errors.New("not an extended JSON ObjectID")
			}
			oid, ok := m["$oid"]
			if !ok {
				return errors.New("not an extended JSON ObjectID")
			}
			str, ok = oid.(string)
			if !ok {
				return errors.New("not an extended JSON ObjectID")
			}
		}

		// An empty string is not a valid ObjectID, but we treat it as a special value that decodes as NilObjectID.
		if len(str) == 0 {
			copy(id[:], NilObjectId[:])
			return nil
		}

		if len(str) != 24 {
			return fmt.Errorf("cannot unmarshal into an ObjectID, the length must be 24 but it is %d", len(str))
		}

		_, err = hex.Decode(id[:], []byte(str))
		if err != nil {
			return err
		}
	}

	return err
}

type Datetime int64
