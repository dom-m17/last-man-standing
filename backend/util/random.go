package util

import (
	"fmt"
	"strings"

	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int32 {
	return int32(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
