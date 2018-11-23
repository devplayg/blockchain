package main

import (
	"bytes"
	"encoding/binary"
)

func Int64ToHex(n int64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, n)
	return buf.Bytes(), err
}
