package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/tamathecxder/randomail"
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

func RandomDateOfBirth() time.Time {
	year := rand.Intn(20) + 1990
	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(28) + 1
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func RandomPassword() string {
	//TODO
	return RandomString(10)
}

func RandomUsername() string {
	//TODO
	return RandomString(8)
}

func RandomEmail() string {
	return randomail.GenerateRandomEmails(1)[0]
}
