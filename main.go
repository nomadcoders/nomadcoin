package main

import (
	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
