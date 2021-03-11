package main

import (
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

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

			extractPools(provider, contractAddress, datadir)
			ticker := time.NewTicker(duration)
			for {
				select {
				case <-ticker.C:
					extractPools(provider, contractAddress, datadir)
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
