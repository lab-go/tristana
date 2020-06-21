package codec

import "encoding/json"

func JsonDecode(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}
