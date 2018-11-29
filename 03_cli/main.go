package main

import "github.com/davecgh/go-spew/spew"

func main() {

	blockchain := NewBlockchain()
	for _, b := range blockchain.Blocks {
		//fmt.Printf("%d - %s (%s)\n", b.Timestamp, b.Data, b.Hash)
		spew.Dump(b)
	}

	//v := uint32(500)
	//buf := make([]byte, 4)
	//binary.BigEndian.PutUint32(buf, v)
	//spew.Dump(buf)
	//
	//x := binary.BigEndian.Uint32(buf)
	//spew.Dump(x )
	//
	//binary.Write()

	//spew.Dump(Int64ToHex(1))
	//spew.Dump(Int64ToHex(10000000000000000000000000000))
}

//
//package main
//
//import (
//	"fmt"
//	"io"
//	"log"
//	"strings"
//)
//
//func main() {
//	r := strings.NewReader("some io.Reader stream to be read\n")
//
//	buf := make([]byte, 4)
//	if _, err := io.ReadFull(r, buf); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s\n", buf)
//
//	// minimal read size bigger than io.Reader stream
//	longBuf := make([]byte, 64)
//	if _, err := io.ReadFull(r, longBuf); err != nil {
//		fmt.Println("error:", err)
//	}
//
//}
//
