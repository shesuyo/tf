package tf

import (
	"encoding/gob"
	"os"
)

// GobEncode gob_encode
func GobEncode(f string, v interface{}) error {
	fh, err := os.Create(f)
	defer fh.Close()
	if err != nil {
		return err
	}
	g := gob.NewEncoder(fh)
	return g.Encode(v)
}

// GobDecode gob_decode
func GobDecode(f string, v interface{}) error {
	fh, err := os.Open(f)
	defer fh.Close()
	if err != nil {
		return err
	}
	g := gob.NewDecoder(fh)
	return g.Decode(v)
}
