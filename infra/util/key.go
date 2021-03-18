package util

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	AlphabetAndNumber = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Number            = "1234567890"
)

func CreateKey(str string, len int) string {
	var res string
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		res += string(str[randomInt.Int64()])
	}
	return res
}
