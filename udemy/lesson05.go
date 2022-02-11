package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "hello world"
	b := getSHA256Binary(s)
	h := hex.EncodeToString(b)

	fmt.Println(s, b, h)
}

func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
