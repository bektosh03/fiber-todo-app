package etc

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode(length int, charset ...string) string {
	var (
		chars      = charset[0]
		seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
		code       = make([]byte, length)
	)
	if len(charset) == 0 {
		chars = defaultCharset
	}
	for i := range code {
		code[i] = chars[seededRand.Intn(len(chars))]
	}
	return string(code)
}
