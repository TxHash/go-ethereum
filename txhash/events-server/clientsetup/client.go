package clientsetup

import (
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var serverMode = flag.String("serverMode","production","Set to 'testing' to enable debug access.")
var server = flag.String("host", "http://localhost:8545", "Set server host.")

var Cl *ethclient.Client

func init() {
	flag.Parse()

	//fmt.Print(*server)
	cl, err := ethclient.Dial(*server)
	if err != nil {
		log.Panic("Connection Error: ", err)
	}
	Cl = cl
}