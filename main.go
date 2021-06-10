package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficutly := 6
	target := strings.Repeat("0", difficutly)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}
	}
}
