package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomPhoneNumber() string {
	b := make([]string, 10)
	for i := range b {
		b[i] = strconv.Itoa(rand.Intn(10))
	}
	res := strings.Join(b, "")
	res = fmt.Sprint("+44", res)

	return res
}

// TODO: Random username, password, email, d.o.b
