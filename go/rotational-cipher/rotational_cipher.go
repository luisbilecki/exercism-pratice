package rotationalcipher

import (
	"unicode"
)

const AlphabetSize = 26
const UpperCaseA = 65
const LowerCaseA = 97

func RotationalCipher(plain string, shiftKey int) string {
	cipher := ""
	for _, char := range plain {
		if unicode.IsUpper(char) {
			cipher += string((char+rune(shiftKey)-rune(UpperCaseA))%AlphabetSize + UpperCaseA)
		} else if unicode.IsLower(char) {
			cipher += string((char+rune(shiftKey)-rune(LowerCaseA))%AlphabetSize + LowerCaseA)
		} else {
			cipher += string(char)
		}
	}
	return cipher
}
