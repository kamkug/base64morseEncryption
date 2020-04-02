package morse

import (
	"testing"
)

const succeed = "\u2714"
const failed = "\u2717"

// TestDownload validates the http Get function can download content.
/*func TestDownload(t *testing.T) {
	url := "https://www.ardanlabs.com/blog/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code : %d", failed, statusCode, resp.StatusCode)
			}
		}
	}
}*/

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
