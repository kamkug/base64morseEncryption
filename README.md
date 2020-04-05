Encryption : base64 -> morse and decryption back to the original plaintext

Author: Kamil Kugler

Package morse provides a convenient API for encryption
of some plaintext into ciphertext in the following way:
- Base64 -> morse

Where morse is actually modified in a way that it can separate between lower and upper case letters:
if the morse code represents a lower case letter:
- if composed of dots (".") only I am replacing them with stars ("*")
- if composed of hyphens ("-") and dots (".") I am replacing hyphens ("-") with underscores ("_")

Please do not forget that this is a program of a hobbist and should only be used at your own risk
if in production or for any other purpose, however I have a hope that you will enjoy it

Change "/usr/bin/go" based on the placement of your go binary file in your system
