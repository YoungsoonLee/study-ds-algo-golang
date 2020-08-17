package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type OneTimeToken struct {
	UserNo int64  `json:"user_no"`
	UserId string `json:"user_id"`
}

func main() {
	test := OneTimeToken{UserNo: 1755891, UserId: "jeongjh1105"}
	//originalText := "encrypt this golang"
	s, _ := json.Marshal(test)
	fmt.Println(string(s))

	key := []byte("4c6ThMX2MWF8a^Co")

	// encrypt value to base64
	cryptoText := encrypt(key, string(s))
	fmt.Println(cryptoText)

	// encrypt base64 crypto to original value
	text := decrypt(key, cryptoText)
	fmt.Println(text)

	var ott OneTimeToken
	err := json.Unmarshal([]byte(text), &ott)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ott.UserId)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
