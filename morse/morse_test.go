package morse

import (
	"testing"
)

const succeed = "\u2714"
const failed = "\u2717"

type test struct {
	plaintext  string
	ciphertext string
}

func TestBase64MorseEncode(t *testing.T) {
	tests := []test{
		{"hello", "._ --. ...- *** _... --. ---.. -...- "},
		{"asia", "-.-- -..- -. .__. -.-- --.-  -...- -...- "},
		{"Asia", "--.- -..- -. .__. -.-- --.-  -...- -...- "},
		{"0123567890", "-- -.. . _.__ -- __.. ..- ..--- -. __.. __. ..... -- .-  -...- -...- "},
		{"Homo", "... --. ----. _ _... .__  -...- -...- "},
		{"Lorep Ipsum", "- --. ----. _.__ --.. -..- .- __. ... -..- -... __.. _.. .-- ----- -...- "},
		{"gopher", "--.. ..--- ----. .__ ._ --. ...- _.__ "},
		{"GOLANG", ".-. ----- ----. -- --.- ..- ..... .... "},
		{"enCIPh7R", "--.. .-- ..... -.. ... ...- -... ___ -. .---- .. -...- "},
		{"Hello World", "... --. ...- *** _... --. ---.. __. ...- ..--- ----. _.__ _... --. --.- -...- "},
		{"My name is morse", "- -..- _._ __. _... __ ..-. _ --.. ... -... .__. _._. _.__ -... _ _... ...-- .--- __.. --.. --.-  -...- -...- "},
	}

	t.Log("Given the need to test encoding plaintext.")
	{
		for i, v := range tests {
			t.Logf("\tTest %d:\tWhen encoding %q into ciphertext %q", i, v.plaintext, v.ciphertext)
			{
				cipher := Base64MorseEncode(v.plaintext)
				if cipher != v.ciphertext {
					t.Fatalf("\t%s\tShould be able to encode into a correct ciphertext", failed)
				}
				t.Logf("\t%s\tShould be able to encode into a correct ciphertext", succeed)
			}
		}
	}
}

func TestDecodeMorse(t *testing.T) {
	tests := []test{
		{"._ --. ...- *** _... --. ---.. -...- ", "hello"},
		{"-.-- -..- -. .__. -.-- --.-  -...- -...- ", "asia"},
		{"--.- -..- -. .__. -.-- --.-  -...- -...- ", "Asia"},
		{"-- -.. . _.__ -- __.. ..- ..--- -. __.. __. ..... -- .-  -...- -...- ", "0123567890"},
		{"... --. ----. _ _... .__  -...- -...- ", "Homo"},
		{"- --. ----. _.__ --.. -..- .- __. ... -..- -... __.. _.. .-- ----- -...- ", "Lorep Ipsum"},
		{"--.. ..--- ----. .__ ._ --. ...- _.__ ", "gopher"},
		{".-. ----- ----. -- --.- ..- ..... .... ", "GOLANG"},
		{"--.. .-- ..... -.. ... ...- -... ___ -. .---- .. -...- ", "enCIPh7R"},
		{"... --. ...- *** _... --. ---.. __. ...- ..--- ----. _.__ _... --. --.- -...- ", "Hello World"},
		{"- -..- _._ __. _... __ ..-. _ --.. ... -... .__. _._. _.__ -... _ _... ...-- .--- __.. --.. --.-  -...- -...- ", "My name is morse"},
	}

	t.Log("Given the need to test decoding of the ciphertext.")
	{
		for i, v := range tests {
			ciphertext := v.plaintext
			plaintext := v.ciphertext
			t.Logf("\tTest %d:\tWhen decoding %q into plaintext %q", i, ciphertext, plaintext)
			{
				_, p := DecodeMorse(ciphertext)
				if plaintext != p {
					t.Fatalf("\t%s\tShould be able to decode into plaintext", failed)
				}
				t.Logf("\t%s\tShould be able to decode into plaintext", succeed)
			}
		}
	}

}
