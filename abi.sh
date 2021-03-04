#!/bin/sh
# IOption.abi  IOptionPool.abi  IPoolerToken.abi
 ~/go/src/github.com/ethereum/go-ethereum/cmd/abigen/abigen --abi IOption.abi --pkg main --type IOption --out IOption.go
 ~/go/src/github.com/ethereum/go-ethereum/cmd/abigen/abigen --abi IOptionPool.abi --pkg main --type IOptionPool --out IOptionPool.go
 ~/go/src/github.com/ethereum/go-ethereum/cmd/abigen/abigen --abi IPoolerToken.abi --pkg main --type IPoolerToken --out IPoolerToken.go
