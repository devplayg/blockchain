package main

import "fmt"

func main() {

	//newblockch

	blockchain := NewBlockchain()
	blockchain.AddBlock("send $$$ to won")

	for _, b := range blockchain.Blocks {
		fmt.Printf("time=%d, hash=%x.., data=%s, prev_block_hash=%x\n", b.Timestamp, string(b.Hash[:9]), b.Data, string(b.PrevBlockHash))
	}
}
