package payload

import (
	"encoding/hex"
	"encoding/json"
)

type Metadata struct {
	IV  []byte
	Key []byte
}

type jsonMetada struct {
	IV  string
	Key string
}

func (this *Metadata) MarshalJSON() ([]byte, error) {
	ret := &jsonMetada{
		IV:  hex.EncodeToString(this.IV),
		Key: hex.EncodeToString(this.Key),
	}
	return json.Marshal(ret)
}

func (this *Metadata) UnmarshalJSON(data []byte) (err error) {
	ret := &jsonMetada{}
	if err = json.Unmarshal(data, ret); err != nil {
		return err
	}
	if this.IV, err = hex.DecodeString(ret.IV); err != nil {
		return
	} else if this.Key, err = hex.DecodeString(ret.Key); err != nil {
		return
	}
	return
}
