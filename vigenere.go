package main

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
			keyCharIndex := keyIndex % len(keyword)

			base := int(s[i] - 'A')
			shift := (base + int(keyword[keyCharIndex]-'a')) % 26

			result += string(rune(shift + 'A'))
			keyIndex++
		}else {
			result += string(s[i])
		}
	}

	return result
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
			keyCharIndex := keyIndex % len(keyword)

			base := int(s[i] - 'A')
			shift := (base - int(keyword[keyCharIndex]-'a') + 26) % 26

			result += string(rune(shift + 'A'))
			keyIndex++
		} else{
			result += string(s[i])
		}
	}

	return result
}