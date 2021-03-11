package main

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "PandaView",
		Usage:                "A view to collection pool data",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "aggregator",
				Value: "0x46887B5c77B11D262ac17740ED134CEcD39F0DE8",
				Usage: "AggregateUpdater contract address",
			},
			&cli.StringFlag{
				Name:  "provider",
				Value: "https://bsc-dataseed2.defibit.io/",
				Usage: "RPC service address",
			},
		},
		Action: func(c *cli.Context) error {
			contractAddress := common.HexToAddress(c.String("aggregator"))
			provider := c.String("provider")
			log.Println("Aggregator Contract:", contractAddress)
			log.Println("Provider:", provider)

			updateView(provider, contractAddress)
			ticker := time.NewTicker(1 * time.Minute)
			for {
				select {
				case <-ticker.C:
					updateView(provider, contractAddress)
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func updateView(provider string, address common.Address) {
	// create connection
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Printf("PandaView: connection to  %s failed, reason: %s", provider, err)
		return
	}
	defer client.Close()

	instance, err := NewAggregateUpdater(address, client)
	if err != nil {
		log.Println("PandaView: NewAggregateUpdater failed:", err)
		return
	}

	// query pools
	poolId := new(big.Int)
	for {
		pool, err := instance.Pools(nil, poolId)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Handle Pool", pool)
		handlePool(pool, client)
		poolId.Add(poolId, common.Big1)
	}
}

func handlePool(address common.Address, client *ethclient.Client) {
	instance, err := NewIOptionPool(address, client)
	if err != nil {
		log.Println("PandaView: NewAggregateUpdater failed:", err)
		return
	}

	// read all options
	options, err := instance.ListOptions(nil)
	if err != nil {
		log.Fatal(err)
	}

	// handle round data from options
	for k := range options {
		log.Println("Handle Option", options[k])
		handleOption(options[k], client)
	}
}

type RoundData struct {
	R uint64
	S *big.Int
}

const roundLimit = 8640

func handleOption(address common.Address, client *ethclient.Client) {
	var rounds []RoundData
	option, err := NewIOption(address, client)
	currentRound, err := option.GetRound(nil)
	if err != nil {
		log.Fatal(err)
	}

	var roundCount uint16
	for r := currentRound.Uint64(); r > 0; r-- {
		accPremiumShare, err := option.GetRoundAccPremiumShare(nil, new(big.Int).SetUint64(r))
		if err != nil {
			log.Fatal(err)
		}
		rounds = append(rounds, RoundData{r, accPremiumShare})
		roundCount++
		if roundCount >= roundLimit {
			break
		}
	}

	bts, err := json.Marshal(rounds)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bts))
}
