//usr/bin/go run $0 $@; exit $?

package morse

import (
	"fmt"
	"strconv"
)

func Base64MorseEncode(plaintext string) string {
	const padding = `00000000`
	var s string

	base64 := GetBase64Dict()
	morse := GetMorseDict()

	for _, v := range plaintext {
		if a := fmt.Sprintf("%b", uint8(v)); len(a) < 8 {
			padding := 8 - len(a)
			for i := 0; i < padding; i++ {
				s += "0"
			}
			s += a
		}
	}

	length := len(plaintext)

	var requiredPadding int
	switch {
	case length < 3:
		requiredPadding = 3 - length
	case length%3 == 0:
		requiredPadding = 0
	default:
		requiredPadding = 3 - (length % 3)
	}

	if requiredPadding != 0 {
		for i := 0; i < requiredPadding; i++ {
			s += padding
		}
	}
	lengthS := len(s)
	var ciphertext string
	for i := 0; i < lengthS; i += 24 {
		triplet := s[i : i+24]
		for j := 18; j >= 0; j -= 6 {
			a, _ := strconv.ParseInt(triplet, 2, len(triplet))
			lttr := int((a >> j) & 0x3F)
			b64 := base64[lttr]
			//fmt.Printf(b64)
			ciphertext += morse[b64] + " "
		}
	}
	//fmt.Println()
	ciphertextLength := len(ciphertext)
	//fmt.Println(ciphertext)
	switch requiredPadding {
	case 1:
		//fmt.Println("check1:", ciphertext[:ciphertextLength-3])
		ciphertext = fmt.Sprintf("%v%v ", ciphertext[:ciphertextLength-3], morse["="])
	case 2:
		//fmt.Println("check2:", ciphertext[:ciphertextLength-6])
		ciphertext = fmt.Sprintf("%v %v %v ", ciphertext[:ciphertextLength-6], morse["="], morse["="])
	}
	//fmt.Println(ciphertext)
	return ciphertext
}

func DecodeMorse(encodedWord string) (string, string) {
	var word1 string
	var decoded string
	var counter int

	unmorseIt := unmorsedDict()
	for _, v := range encodedWord {
		empty := ' '
		if v != empty {
			word1 += string(v)
		}
		switch {
		case v == 32:
			if unmorseIt[word1] == "=" {
				counter++
			}
			decoded += unmorseIt[word1]
			//fmt.Println(word1, "decoded to:", decoded)
			word1 = ""
		case v == 61:
			counter++
			//fmt.Println("oki")
			decoded += "="
			word1 = ""
		}
	}

	dict := DecodeBase64Dict()
	var s string
	var p string
	for i := 0; i < len(decoded); i++ {

		s += fmt.Sprintf("%06b", dict[string(decoded[i])])
	}
	//fmt.Println(len(s), counter)
	for i := 0; i < len(s)-counter*8; i += 8 {
		letter, _ := strconv.ParseInt((s[i : i+8]), 2, 8)
		p += string(letter)
	}

	return decoded, p
}

func GetBase64Dict() map[int]string {
	base64 := make(map[int]string)
	for i := 0; i < 64; i++ {
		switch {
		case i >= 0 && i < 26:
			base64[i] = string(i + 65)
		case i > 25 && i < 52:
			base64[i] = string((i - 26) + 97)
		case i > 51 && i < 62:
			base64[i] = string((i - 62) + 58)

		case i == 62:
			base64[i] = "+"
		case i == 63:
			base64[i] = "/"
		}
	}
	return base64

}

func DecodeBase64Dict() map[string]int {
	decodeB64 := make(map[string]int)
	var respectiveValue int
	for i := 65; i < 91; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}

	for i := 97; i < 123; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}

	for i := 48; i < 58; i++ {
		decodeB64[string(i)] = respectiveValue
		respectiveValue++
	}
	decodeB64["+"] = 62
	decodeB64["/"] = 63
	//fmt.Println(decodeB64)
	return decodeB64
	//s := fmt.Sprintf("%06b", decodeB64["A"])
	//fmt.Println(s, "in base64 binary representation")
}

func GetMorseDict() map[string]string {
	morse := make(map[string]string)

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
	return morse
}

func unmorsedDict() map[string]string {

	unmorsed := make(map[string]string)

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
	return unmorsed
}
