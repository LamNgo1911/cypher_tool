///usr/bin/env go run "$0" "$@"; exit $?

package main

import "fmt"
import "bufio"
import "os"
// import "strconv"
import "strings"


func main (){
	var input string
	var option1 int
	var option2 rune
	var result string

	fmt.Println("Welcome to the Cypher Tool!\n")

	for{
		fmt.Println("Select operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		fmt.Scanln(&input) //-> this will stop when it hits the whitespace

		fmt.Println("")

		if input == string('1') || input == string('2') {
			option1 = 1
			break
		} else{
			fmt.Println("Please choose 1 or 2 only, my friend! ^-^\n")
		}
	}

	for{
		fmt.Println("Select cypher (1/2):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Scanln(&input)
		fmt.Println("")

		if input == string('1') || input == string('2') {
			option2 = 2
			break
		} else{
			fmt.Println("Please choose 1 or 2 only, my friend! ^-^\n")
		}
	}

	for{
		fmt.Println("Enter the message:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		trimmed := strings.TrimSpace(text)
		
		if trimmed == ""{
			message := ""
			if option1 == 1{
				message = "encrypt\n"
			} else{
				message = "Decrypt\n"
			}
			fmt.Println("Please add the message you would like to " + message)
		}else {
			if(option1 == 1){
			if(option2 == 2){
				result = encrypt_rot13(trimmed)
			}else{
				result = encrypt_reverse(trimmed)
			}
			} else{
				if(option2 == 2){
					result = decrypt_rot13(trimmed)
				}else{
					result = decrypt_reverse(trimmed)
				}
			}
			break
		}
	}

	fmt.Println("")

	if(option1 == 1){
		fmt.Println("Encrypted message using reverse:")
	} else {
		fmt.Println("Decrypted message using reverse:")
	}

	fmt.Println(result)
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	return "Bingo!"
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
	return "Bingo!"
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	return "Bingo!"
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	return "Bingo!"
}