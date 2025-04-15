package implementation

import (
	"errors"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type BlockchainClient struct {
	EthClient *ethclient.Client
	BtcClient *rpcclient.Client
}

func NewBlockchainClient(ethRPC, btcHost, btcUser, btcPass string) (*BlockchainClient, error) {
	ethClient, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, errors.New("failed to connect to Ethereum node")
	}
	btcClient, btcErr := rpcclient.New(&rpcclient.ConnConfig{
		Host:         btcHost,
		User:         btcUser,
		Pass:         btcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if btcErr != nil {
		log.Println("failed to connect to Bitcoin node")
		btcClient = nil
	}
	return &BlockchainClient{EthClient: ethClient, BtcClient: btcClient}, nil
}
