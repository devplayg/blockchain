package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"encoding/binary"
)

func Int64ToHex(n int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		log.Error(err)
	}
	return buf.Bytes()
}

