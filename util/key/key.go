package key

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	all    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	number = "1234567890"
)

func CreateToken() string {
	return create(all, 64)
}

func CreateSmsCode() string {
	return create(number, 6)
}

func create(str string, len int) string {
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
