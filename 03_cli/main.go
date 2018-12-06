package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()
	//for _, b := range blockchain.Blocks {
		//fmt.Printf("%d - %s (%s) (%3.1fs)\n", b.Timestamp, b.Data, b.Hash, b.dur.Seconds())
		//spew.Dump(b)
	//}

	bc.AddBlock("hello world! what's your name?")
	bc.AddBlock("my name is won")

	cli := CLI{bc}

	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		//fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	//
	//
	//b := Block{
	//	Hash: []byte("ab0123"),
	//	Nonce: 3,
	//	Timestamp: time.Now().Unix(),
	//}
	//
	////b.se
	//
	//data := b.Serialize()
	////spew.Dump(b)
	////spew.Dump(data)
	//
	//b2 := DeserializeBlock(data)
	//spew.Dump(b2)

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
