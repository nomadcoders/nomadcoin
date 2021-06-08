package main

import (
	"github.com/nomadcoders/nomadcoin/blockchain"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
