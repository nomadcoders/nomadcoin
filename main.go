package main

import (
	"github.com/nomadcoders/nomadcoin/blockchain"
	"github.com/nomadcoders/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
