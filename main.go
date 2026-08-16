///usr/bin/env go run "$0" "$@"; exit $?

package main

import "fmt"
import "bufio"
import "os"
// import "strconv"
import "strings"


func main (){
	var input string
	var option1 string
	var option2 string
	var methodType string
	var result string

	fmt.Println("Welcome to the Cypher Tool!\n")

	for{
		fmt.Println("Select operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		fmt.Scanln(&input) //-> this will stop when it hits the whitespace

		fmt.Println("")

		if input == string('1') || input == string('2') {
			option1 = input
			input = ""
			break
		} else{
			fmt.Println("Please choose 1 or 2 only, my friend! ^-^\n")
		}
	}

	for{
		fmt.Println("Select cypher (1/2):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Vigenere Cipher.")
		fmt.Scanln(&input)
		fmt.Println("")

		if input == string('1') || input == string('2') || input == string('3'){
			option2 = input
			input = ""
			break
		} else{
			fmt.Println("Please choose 1, 2 or 3 only, my friend! ^-^\n")
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
			if option1 == "1"{
				message = "encrypt\n"
			} else{
				message = "Decrypt\n"
			}
			fmt.Println("Please add the message you would like to " + message)
		}else {
			if(option1 == "1"){
				if(option2 == "1"){
					result = encrypt_rot13(trimmed)
					methodType = "ROT13:"
				} else if(option2 == "2"){
					result = encrypt_reverse(trimmed)
					methodType = "Reverse:"
				} else{
					result = encrypt_Vigenere_Cipher(trimmed)
					methodType = "Vigenere Cipher:"
				}
			} else {
				if(option2 == "1"){
					result = decrypt_rot13(trimmed)
					methodType = "ROT13:"
				} else if(option2 == "2"){
					result = decrypt_reverse(trimmed)
					methodType = "Reverse:"
				} else{
					result = decrypt_Vigenere_Cipher(trimmed)
					methodType = "Vigenere Cipher:"
				}
			}
			break
		}
	}

	fmt.Println("")

	if(option1 == "1"){
		fmt.Printf("Encrypted message using %s\n", methodType )
	} else {
		fmt.Printf("Decrypted message using %s\n", methodType)
	}

	methodType = ""

	fmt.Println(result)
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	return "encrypt_rot13!"
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
	return "encrypt_reverse!"
}

// Encrypt the message with a keyword
func encrypt_Vigenere_Cipher(s string) string {
	result := ""
	keyword := "lam"
	keyIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			keyCharIndex := keyIndex % len(keyword)

			base := int(s[i] - 'a')
			shift := (base + int(keyword[keyCharIndex]-'a')) % 26

			result += string(rune(shift + 'a'))
			keyIndex++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			keyCharIndex := i % len(keyword)

			base := int(s[i] - 'A')
			shift := (base + int(keyword[keyCharIndex]-'A')) % 26

			result += string(rune(shift + 'A'))
			keyIndex++
		}else {
			result += string(s[i])
		}
	}

	return result
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	return "decrypt_rot13!"
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	return "decrypt_reverse!"
}

// Decrypt the message with a keyword
func decrypt_Vigenere_Cipher(s string) string {
	result := ""
	keyword := "lam"
	keyIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			keyCharIndex := keyIndex % len(keyword)

			base := int(s[i] - 'a')
			shift := (base - int(keyword[keyCharIndex]-'a') + 26) % 26

			result += string(rune(shift + 'a'))
			keyIndex++
		} else if s[i] >= 'A' && s[i] <= 'Z'{
			keyCharIndex := i % len(keyword)

			base := int(s[i] - 'A')
			shift := (base - int(keyword[keyCharIndex]-'A') + 26) % 26

			result += string(rune(shift + 'A'))
			keyIndex++
		} else{
			result += string(s[i])
		}
	}

	return result
}