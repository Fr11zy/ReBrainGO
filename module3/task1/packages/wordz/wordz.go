package wordz

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

var Hello = "This is package wordz"
var Prefix = "Random word: "
var words = []string{"one","two","three","four", "five"}

func init() {
	fmt.Println("Function init in package wordz")
}

func Random() string {
	max := len(words)
	r,_ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return words[index]
}