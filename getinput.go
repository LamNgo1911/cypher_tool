package main

import "fmt"
import "os"
import "bufio"
// import "strconv"
import "strings"

func getInput() (toEncrypt bool, encoding string, message string) {
	var input string

	fmt.Println("Welcome to the Cypher Tool!\n")

	for{
		fmt.Println("Select operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		fmt.Scanln(&input) //-> this will stop when it hits the whitespace

		fmt.Println("")

		if input == string('1') || input == string('2'){
			if input == string('1'){
				toEncrypt = true
			} else{
				toEncrypt = false
			}
			input = ""
			break
		}else{
			fmt.Println("Please choose 1 or 2 only, my friend! ^-^\n")
		}
	}

	for{
		fmt.Println("Select cypher (1/3):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Vigenere Cipher.")
		fmt.Scanln(&input)
		fmt.Println("")

		if input == string('1') || input == string('2') || input == string('3'){
			if input == string('1'){
				encoding =  "ROT13"
			}else if input == string('2'){
				encoding =  "Reverse"
			} else{
				encoding =  "Vigenere Cipher"
			}

			input = ""
			break
		}else{
			fmt.Println("Please choose 1, 2 or 3 only, my friend! ^-^\n")
		}
	}

	for{
		fmt.Println("Enter the message:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		message = strings.TrimSpace(text)

		if(message == ""){
			operationType := ""
			if toEncrypt{
				operationType = "Encrypt\n"
			} else{
				operationType = "Decrypt\n"
			}
			fmt.Println("Please add the message you would like to " + operationType)
		} else{
			break
		}
	}

	return toEncrypt, encoding, message
}