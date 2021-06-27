package main

import (
	"github.com/nomadcoders/nomadcoin/cli"
	"github.com/nomadcoders/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
