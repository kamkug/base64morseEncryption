//usr/bin/go run $0 $@; exit $?
// Author: Kamil Kugler
// package morse provides a convenient API for encryption
// of some plaintext into ciphertext in the following way:
// Base64 -> morse
//
// Where morse is actually modified in a way that it can separate between lower and upper case letters:
// if the morse code represents a lower case letter:
// - if composed of dots (".") only I am replacing them with stars ("*")
// - if composed of hyphens ("-") and dots (".") I am replacing hyphens ("-") with underscores ("_")
//
// Please do not forget that this is a program of a hobbist and should only be used at your own risk
// if in production or for any other purpose, however I have a hope that you will enjoy it
//
// Change "/usr/bin/go" based on the placement of your go binary file in your system
package morse

import (
	"fmt"
	"strconv"
)

//Base64MorseEncode provides a full plaintext->base64->morse encryption and spits out a ciphertext
func Base64MorseEncode(plaintext string) string {
	var s string
	// collecting required dictionaries for the ease of encryption
	base64 := GetBase64Dict()
	morse := GetMorseDict()
	// making sure that each of the characters will be represented by its full ubyte8 representation
	for _, v := range plaintext {
		char := fmt.Sprintf("%08b", uint8(v))
		s += char
	}

	length := len(plaintext)
	// verify if any padding is actually required to be able to keep our plaintext as multiples of 3 bytes
	var requiredPadding int
	switch {
	case length < 3:
		requiredPadding = 3 - length
	case length%3 == 0:
		requiredPadding = 0
	default:
		requiredPadding = 3 - (length % 3)
	}
	// specifying padding for keeping multiples of 3 bytes in terms of base64 encrypted plaintext
	const padding = `00000000`
	// if padding was required we are adding it at the back
	if requiredPadding != 0 {
		for i := 0; i < requiredPadding; i++ {
			s += padding
		}
	}
	lengthS := len(s)
	// actual encryption in its current form
	var ciphertext string
	// ensuring that we are moving triplet at a times
	for i := 0; i < lengthS; i += 24 {
		triplet := s[i : i+24]
		// making sure that we are working with 6 bit values only for successful base64 encryption
		for j := 18; j >= 0; j -= 6 {
			// we are capturing the whole triplet in here
			a, _ := strconv.ParseInt(triplet, 2, len(triplet))
			// making sure that we are only using next 6 bits from the sequence by:
			// - shifting everything to the right by j bits (simply multiples of 6)
			// - performing a bitwise & operation to ensure that everything else is just 0's
			lttr := int((a >> j) & 0x3F)
			// actually encrypting a single 6 bit chunk into a base64 code using our base64 dictionary
			b64 := base64[lttr]
			// encrypting our base64 character into its morse alternative using our morse dictionary
			ciphertext += morse[b64] + " "
		}
	}

	// making sure that outstanding 6 bits sequences (all zeroes) from padding (at the back) are marked as "="
	// depending on the length of our padding
	ciphertextLength := len(ciphertext)
	oneHexPadding := 3
	twoHexPadding := 6
	switch requiredPadding {
	case 1:
		ciphertext = fmt.Sprintf("%v%v ", ciphertext[:ciphertextLength-oneHexPadding], morse["="])
	case 2:
		ciphertext = fmt.Sprintf("%v %v %v ", ciphertext[:ciphertextLength-twoHexPadding], morse["="], morse["="])
	}
	// happily return our ciphertext
	return ciphertext
}

// DecodeMorse provides decoding functionality following : morse->base64->plaintext
// it returns both : base64 and plaintext translation
func DecodeMorse(encodedWord string) (string, string) {
	var morsechar string
	var b64 string
	var counter int
	// load a dictionary of "unmorsed" characters
	unmorseIt := UnmorsedDict()
	// look through each morse character inside of the ciphertext
	for _, v := range encodedWord {
		// decode each morse char into a base64 character
		switch {
		// capture a single morsechar if between words (one whitespace character)
		// will not work with more then one space in between
		case v == 32:
			// count the amount of padding
			if unmorseIt[morsechar] == "=" {
				counter++
			}
			// decrypt a morse char into a base64 code unit
			b64 += unmorseIt[morsechar]
			morsechar = ""
		// else build a morse character
		default:
			morsechar += string(v)
		}
	}
	// load up a dictionary to decode base64 to plaintext
	dict := DecodeBase64Dict()
	var s string
	var p string
	// simply go over each character from our decoded morse string
	for i := 0; i < len(b64); i++ {
		// ensure that each character is of length 6 bits and create a string for decoding
		s += fmt.Sprintf("%06b", dict[string(b64[i])])
	}
	// traverse at the rate of 8 bits
	for i := 0; i < len(s)-counter*8; i += 8 {
		// decode to ascii (plaintext) by decoding each of the 8 bit chunks into ascii
		letter, _ := strconv.ParseInt((s[i : i+8]), 2, 8)
		p += string(letter)
	}
	// return base64 and plaintext
	return b64, p
}

// GetBase64Dict used to generates a base64 dictionary for encoding
func GetBase64Dict() map[int]string {
	base64 := make(map[int]string)
	// dynamically generate all of the 64 base64 chars
	for i := 0; i < 64; i++ {
		switch {
		// generate upper case letters entries
		case i >= 0 && i < 26:
			base64[i] = string(i + 65)
		// generate lower case letters entries
		case i > 25 && i < 52:
			base64[i] = string((i - 26) + 97)
		// generate numeric entries
		case i > 51 && i < 62:
			base64[i] = string((i - 62) + 58)
		// generate remaining entries
		case i == 62:
			base64[i] = "+"
		case i == 63:
			base64[i] = "/"
		}
	}
	// return the dictionary
	return base64

}

// DecodeBase64Dict generates a base64 char to its numeric representation dictionary, used for decoding
func DecodeBase64Dict() map[string]int {
	decodeB64 := make(map[string]int)
	var respectiveValue int
	// generate upper case letters entries
	for i := 65; i < 91; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}
	// generate lower case letters entries
	for i := 97; i < 123; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}
	// generate numeric entries
	for i := 48; i < 58; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}
	// add the other two chars
	decodeB64["+"] = 62
	decodeB64["/"] = 63
	// return our dictionary
	return decodeB64
}

// GetMorseDict generates our base64 char to morse dictionary
func GetMorseDict() map[string]string {
	morse := make(map[string]string)
	// not too complicated
	morse["A"] = ".-"
	morse["a"] = "._"
	morse["B"] = "-..."
	morse["b"] = "_..."
	morse["C"] = "-.-."
	morse["c"] = "_._."
	morse["D"] = "-.."
	morse["d"] = "_.."
	morse["E"] = "."
	morse["e"] = "*"
	morse["F"] = "..-."
	morse["f"] = ".._."
	morse["G"] = "--."
	morse["g"] = "__."
	morse["H"] = "...."
	morse["h"] = "****"
	morse["I"] = ".."
	morse["i"] = "**"
	morse["J"] = ".---"
	morse["j"] = ".___"
	morse["K"] = "-.-"
	morse["k"] = "_._"
	morse["L"] = ".-.."
	morse["l"] = "._.."
	morse["M"] = "--"
	morse["m"] = "__"
	morse["N"] = "-."
	morse["n"] = "_."
	morse["O"] = "---"
	morse["o"] = "___"
	morse["P"] = ".--."
	morse["p"] = ".__."
	morse["Q"] = "--.-"
	morse["q"] = "__._"
	morse["R"] = ".-."
	morse["r"] = "._."
	morse["S"] = "..."
	morse["s"] = "***"
	morse["T"] = "-"
	morse["t"] = "_"
	morse["U"] = "..-"
	morse["u"] = ".._"
	morse["V"] = "...-"
	morse["v"] = "..._"
	morse["W"] = ".--"
	morse["w"] = ".__"
	morse["X"] = "-..-"
	morse["x"] = "_.._"
	morse["Y"] = "-.--"
	morse["y"] = "_.__"
	morse["Z"] = "--.."
	morse["z"] = "__.."
	morse["0"] = "-----"
	morse["1"] = ".----"
	morse["2"] = "..---"
	morse["3"] = "...--"
	morse["4"] = "....-"
	morse["5"] = "....."
	morse["6"] = "-...."
	morse["7"] = "--..."
	morse["8"] = "---.."
	morse["9"] = "----."
	morse["="] = "-...-"
	morse["/"] = "-..-."
	morse["+"] = ".-.-."
	// return our dictionary for encoding
	return morse
}

// UnmorseDict generates our morse code to base64 dictionary for decryption
func UnmorsedDict() map[string]string {

	unmorsed := make(map[string]string)
	// not too complicated
	unmorsed[".-"] = "A"
	unmorsed["._"] = "a"
	unmorsed["-..."] = "B"
	unmorsed["_..."] = "b"
	unmorsed["-.-."] = "C"
	unmorsed["_._."] = "c"
	unmorsed["-.."] = "D"
	unmorsed["_.."] = "d"
	unmorsed["."] = "E"
	unmorsed["*"] = "e"
	unmorsed["..-."] = "F"
	unmorsed[".._."] = "f"
	unmorsed["--."] = "G"
	unmorsed["__."] = "g"
	unmorsed["...."] = "H"
	unmorsed["****"] = "h"
	unmorsed[".."] = "I"
	unmorsed["**"] = "i"
	unmorsed[".---"] = "J"
	unmorsed[".___"] = "j"
	unmorsed["-.-"] = "K"
	unmorsed["_._"] = "k"
	unmorsed[".-.."] = "L"
	unmorsed["._.."] = "l"
	unmorsed["--"] = "M"
	unmorsed["__"] = "m"
	unmorsed["-."] = "N"
	unmorsed["_."] = "n"
	unmorsed["---"] = "O"
	unmorsed["___"] = "o"
	unmorsed[".--."] = "P"
	unmorsed[".__."] = "p"
	unmorsed["--.-"] = "Q"
	unmorsed["__._"] = "q"
	unmorsed[".-."] = "R"
	unmorsed["._."] = "r"
	unmorsed["..."] = "S"
	unmorsed["***"] = "s"
	unmorsed["-"] = "T"
	unmorsed["_"] = "t"
	unmorsed["..-"] = "U"
	unmorsed[".._"] = "u"
	unmorsed["...-"] = "V"
	unmorsed["..._"] = "v"
	unmorsed[".--"] = "W"
	unmorsed[".__"] = "w"
	unmorsed["-..-"] = "X"
	unmorsed["_.._"] = "x"
	unmorsed["-.--"] = "Y"
	unmorsed["_.__"] = "y"
	unmorsed["--.."] = "Z"
	unmorsed["__.."] = "z"
	unmorsed["-----"] = "0"
	unmorsed[".----"] = "1"
	unmorsed["..---"] = "2"
	unmorsed["...--"] = "3"
	unmorsed["....-"] = "4"
	unmorsed["....."] = "5"
	unmorsed["-...."] = "6"
	unmorsed["--..."] = "7"
	unmorsed["---.."] = "8"
	unmorsed["----."] = "9"
	unmorsed["-...-"] = "="
	unmorsed["-..-."] = "/"
	unmorsed[".-.-."] = "+"
	// return our dictionary
	return unmorsed
}
