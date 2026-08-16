package kaspar

import "fmt"

func encrypt_reverse(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			total := int('z' - 'a') 
			step := int(s[i] - 'a')
			result += string(total - step + 'a')
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			total := int('Z' - 'A')
			step := int(s[i] - 'A')
			result += string(total - step + 'A')
		} else {
			result += string(s[i])
		}
	}
	return result
}

func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}

func encrypt_rot13(s string) string {
	result := ""
	for _, r := range s {
		dist := (r - 'a')
		step := (dist + 13) %26
		result += string(step + 'a')
	}
	return result
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}
