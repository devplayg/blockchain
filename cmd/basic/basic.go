package main

import (
	"github.com/devplayg/blockchain/objs"
)

func main() {


	//newblockch

	blockchain := objs.NewBlockchain()
	blockchain.AddBlock("send $$$ to won")

	//for _, b := range blockchain.Blocks {
		//fmt.Printf("time=%d, hash=%x.., data=%s, prev_block_hash=%x\n",b.Timestamp, string(b.Hash[:9]), b.Data, string(b.PrevBlockHash))
	//}
}
