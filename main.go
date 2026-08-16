package main

import "fmt"
// import "strconv"


func main (){
	toEncrypt, encoding, message := getInput() 

	result := ""

	if(toEncrypt){
		if(encoding == "ROT13"){
			result = encrypt_rot13(message)
		} else if(encoding == "Reverse"){
			result = encrypt_reverse(message)
		}else{
			result = encrypt_Vigenere_Cipher(message)
		}
	}else{
		if(encoding == "ROT13"){
			result = decrypt_rot13(message)
		} else if(encoding == "Reverse"){
			result = decrypt_reverse(message)
		}else{
			result = decrypt_Vigenere_Cipher(message)
		}
	}	

	if(toEncrypt){
		fmt.Printf("\nEncrypted message using %s\n", encoding)
	}else{
		fmt.Printf("\nDecrypted message using %s\n", encoding)
	}

	fmt.Println(result)
}