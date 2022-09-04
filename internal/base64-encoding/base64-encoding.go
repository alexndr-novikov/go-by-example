package base64_encoding

import (
	b64 "encoding/base64"
	"fmt"
)

func RunSample() {
	fmt.Println("Base64 encoding package output:")
	data := "Lorem ipsum 123 !?$*&()'-=@~"
	encoded := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded: ", encoded)
	fmt.Println("")
	decoded, _ := b64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded bytes: ", decoded)
	fmt.Println("Decoded string: ", string(decoded))

	fmt.Println("Url encode")
	urlEncoded := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Urlencoded: ", urlEncoded)
	urlDecoded, _ := b64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println("Url decoded bytes: ", urlDecoded)
	fmt.Println("Url decoded sting: ", string(urlDecoded))

	fmt.Println("")
}
