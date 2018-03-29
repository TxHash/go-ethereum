package main

import (
	"flag"
	"fmt"
	"os"
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var serverMode = flag.String("serverMode","production","Set to 'testing' to enable debug access.")
var server = flag.String("host", "http://localhost:8545", "Set server host.")


func main() {
	flag.Parse()

	ctx := context.Background()

	Cl, err := ethclient.Dial(*server)
	if err != nil {
		log.Panic("Connection Error: ", err)
	}

	prog, err := Cl.SyncProgress(ctx)
	if err != nil {
		log.Panic("Error Fetching Sync Status: ", err)
	}

	if prog == nil {
		fmt.Fprint(os.Stdout, "Syncing complete!\n")
	} else {
		fmt.Fprint(os.Stdout, "Current Block: ", prog.CurrentBlock, "\nHighestBlock: ", prog.HighestBlock)
	}
}