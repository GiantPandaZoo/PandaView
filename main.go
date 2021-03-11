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

type RoundData struct {
	R uint64
	S *big.Int
}

func main() {
	app := &cli.App{
		Name:                 "PandaView",
		Usage:                "An offchain system for data collection",
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
			&cli.StringFlag{
				Name:  "datadir",
				Value: `data/`,
				Usage: "output data directory",
			},
			&cli.DurationFlag{
				Name:  "duration",
				Value: 2 * time.Hour,
				Usage: "data refresh duration",
			},
		},
		Action: func(c *cli.Context) error {
			contractAddress := common.HexToAddress(c.String("aggregator"))
			provider := c.String("provider")
			duration := c.Duration("duration")
			datadir := c.String("datadir")
			log.Println("Aggregator Contract:", contractAddress)
			log.Println("Provider:", provider)
			log.Println("Refresh Duration:", duration)
			log.Println("Data Directory:", datadir)

			if _, err := os.Stat(datadir); os.IsNotExist(err) {
				os.Mkdir(datadir, os.ModeDir|os.ModePerm)
			}

			updateView(provider, contractAddress, datadir)
			ticker := time.NewTicker(duration)
			for {
				select {
				case <-ticker.C:
					updateView(provider, contractAddress, datadir)
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func updateView(provider string, address common.Address, datadir string) {
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
		handlePool(pool, client, datadir)
		poolId.Add(poolId, common.Big1)
	}
}

func handlePool(address common.Address, client *ethclient.Client, datadir string) {
	pool, err := NewIOptionPool(address, client)
	if err != nil {
		log.Println("PandaView: NewAggregateUpdater failed:", err)
		return
	}

	name, err := pool.Name(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pool:", name)

	// read all options
	options, err := pool.ListOptions(nil)
	if err != nil {
		log.Fatal(err)
	}

	// handle round data from options
	for k := range options {
		handleOption(options[k], client, datadir)
	}
}

func handleOption(address common.Address, client *ethclient.Client, datadir string) {
	var rounds []RoundData
	option, err := NewIOption(address, client)
	currentRound, err := option.GetRound(nil)
	if err != nil {
		log.Fatal(err)
	}

	name, err := option.Name(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Option:", name)

	lastMonth := time.Now().Add(-24 * 30 * time.Hour).Unix()
	for r := currentRound.Uint64(); r > 0; r-- {
		expiryDate, err := option.GetRoundExpiryDate(nil, new(big.Int).SetUint64(r))
		if expiryDate.Int64() < lastMonth {
			break
		}

		accPremiumShare, err := option.GetRoundAccPremiumShare(nil, new(big.Int).SetUint64(r))
		if err != nil {
			log.Fatal(err)
		}
		rounds = append(rounds, RoundData{r, accPremiumShare})
	}

	file, err := os.Create(datadir + address.String() + ".json")
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(file)
	enc.Encode(rounds)
	if err != nil {
		log.Fatal(err)
	}
}
