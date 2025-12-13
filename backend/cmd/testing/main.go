package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := bcrypt.CompareHashAndPassword([]byte("$2a$14$1dBqmBQwyI9Mf1Th/2WKleQYBokC2MmgbalD0yfieKEGFWBWlYy22"), []byte("t3stPassword"))
	fmt.Println(err)
}
