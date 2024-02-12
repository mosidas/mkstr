package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func RandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+'}{'[]\\?:\";',./=-"
	var result string
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result += string(letters[num.Int64()])
	}
	return result, nil
}

func main() {
	var (
		length = flag.Int("l", 10, "length")
	)

	flag.Parse()

	str, err := RandomString(*length)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
