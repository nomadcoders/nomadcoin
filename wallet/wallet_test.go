package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "307702010104209ab1bf01ef131b9b40201f80f73e8a322fa1bc93c0f0face07b9b983e5ea2620a00a06082a8648ce3d030107a14403420004a58855a07e004d7579faedee60806d589e4155a4c1c9dda6b2ced1ab8263592f697e73a031d42975837ddde456fa901d28c9c26bfaee2a5ba82686773b162626"
	testPayload string = "00d432c1446b8e1a6c2b35c5fb69ba41b12d2ca69ef27bb7b438fdb7ce9903b6"
	testSig     string = "e4b18b1ba428f7a4ede60664980ed785ac6204beb837bee77919edff5cdcc71460afb6670d8936655c2416f3eb5cb6455d95e4aa4f5350c8b8c40b63d19ca82c"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"04d432c1446b8e1a6c2b35c5fb69ba41b12d2ca69ef27bb7b438fdb7ce9903b6", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
