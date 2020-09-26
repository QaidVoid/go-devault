package devault

import (
	"log"
	"time"
)

// SubscribeNewBlock subscribes to new blocks
func (rpc *RPC) SubscribeNewBlock(c chan int64) {
	b, err := rpc.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	lastBlock := b

	go func() {
		for {
			newBlock, err := rpc.GetBlockCount()
			if err != nil {
				log.Fatal(err)
			}

			if newBlock > lastBlock {
				lastBlock++
				c <- lastBlock
			}
			time.Sleep(time.Second * 15)
		}
	}()
}

// SubscribeNewTransaction subscribes new transactions
func (rpc *RPC) SubscribeNewTransaction(c chan string) {
	go func() {
		newblock := make(chan int64)
		rpc.SubscribeNewBlock(newblock)

		for height := range newblock {
			hash, err := rpc.GetBlockHash(height)
			if err != nil {
				log.Fatal(err)
			}
			block, err := rpc.GetBlock(hash)
			if err != nil {
				log.Fatal(err)
			}

			for _, tx := range block.Transactions {
				c <- tx
			}
		}
	}()
}
