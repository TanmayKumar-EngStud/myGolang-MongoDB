package main

import "crypto/sha512"

func HashPassword(password string) string {
	temp := sha512.Sum512([]byte(password))
	return string(temp[:])
}