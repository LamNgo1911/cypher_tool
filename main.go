package main

import (
	"fmt"
)

func encrypt_rot13(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			distance := r - 'a'
			base := (distance + 13) % 26
			result += string(base + 'a')
		} else if r >= 'A' && r <= 'Z' {
			distance := r - 'A'
			base := (distance + 13) % 26
			result += string(base + 'A')
		} else {
			result += string(r)
		}
	}
	return result
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

func main() {
	encrypted := encrypt_rot13("HeLlO\n")
	fmt.Print(encrypted)
	decrypted := decrypt_rot13(encrypted)
	fmt.Print(decrypted)
}
