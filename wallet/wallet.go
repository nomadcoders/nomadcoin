package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/nomadcoders/nomadcoin/utils"
)

func Start() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	utils.HandleErr(err)

	message := "i love you"

	hashedMessage := utils.Hash(message)

	hashAsBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	utils.HandleErr(err)

	fmt.Printf("R:%d\nS:%d", r, s)

}
