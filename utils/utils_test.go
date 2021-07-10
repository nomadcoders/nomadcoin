package utils

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"
	s := struct{ Test string }{Test: "test"}
	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)
		if x != hash {
			t.Errorf("Expected %s, got %s", hash, x)
		}
	})
	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

func ExampleHash() {
	s := struct{ Test string }{Test: "test"}
	x := Hash(s)
	fmt.Println(x)
	// Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("ToBytes should return a slice of bytes got %s", k)
	}
}
