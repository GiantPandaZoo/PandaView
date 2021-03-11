package main

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RoundData
type RoundData struct {
	R uint64
	S *big.Int
}

// extractPools
func extractPools(provider string, address common.Address, datadir string) {
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
		extractPool(pool, client, datadir)
		poolId.Add(poolId, common.Big1)
	}
}

// extract data from pool
func extractPool(address common.Address, client *ethclient.Client, datadir string) {
	pool, err := NewIOptionPool(address, client)
	if err != nil {
		log.Println("PandaView: NewIOptionPool failed:", err)
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
		extractOption(options[k], client, datadir)
	}
}

// extract data from option
func extractOption(address common.Address, client *ethclient.Client, datadir string) {
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
