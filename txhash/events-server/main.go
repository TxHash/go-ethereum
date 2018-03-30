package main

import (

)
import (
	"github.com/TxHash/go-ethereum/txhash/events-server/clientsetup"
	//"os"
	"log"
	//"math/big"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"os"
)

func main() {
	ctx := context.Background()
	headerChan := make(chan *types.Header)

	subscription, err := clientsetup.Cl.SubscribeNewHead(ctx, headerChan)
	if err != nil {
		log.Panic("Header Subscription Fail: ", err)
	}
	defer subscription.Unsubscribe()

	for i := 1; i < 10; i++ {
		header := <-headerChan
		fmt.Print(header, "\n")
		num := header.Number
		block, err := clientsetup.Cl.BlockByNumber(ctx, num)
		if err != nil {
			log.Panic("Block Fetch Error: ", err)
		}

		txs := block.Transactions()

		for _, t := range txs {
			receipt, err := clientsetup.Cl.TransactionReceipt(ctx, t.Hash())
			if err != nil {
				log.Panic("Receipt Error: ", err)
			}
			for _, lg := range receipt.Logs {
				b, err := lg.MarshalJSON()
				if err != nil {
					log.Panic("JSON Marshalling Error: ", err)
				}
				_, err = os.Stdout.Write(b)
				if err != nil {
					log.Panic("Error Writing Output: ", err)
				}
				fmt.Print("\n")
			}
		}
	}
}