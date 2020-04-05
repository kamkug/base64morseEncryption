package main

import (
	"fmt"
	"morse_program/morse"
	"os"
)

func main() {

	ciphertext := morse.Base64MorseEncode(os.Args[1])
	fmt.Println("Ciphertext:", ciphertext)
	base64, plaintext := morse.DecodeMorse(ciphertext)
	fmt.Println("Base64:", base64)
	fmt.Println("Plaintext:", plaintext)
	//fmt.Printf("\"%v\", \"%v\";\n", plaintext, ciphertext)

}
