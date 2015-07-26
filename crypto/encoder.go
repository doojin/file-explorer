// crypto package provides encoding / decoding strings functionality
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"encoding/base64"
	"errors"
	"strings"
)

var ERR_CIPHERTEXT_TOO_SHORT = errors.New("Ciphertext is too short")

// Contains methods for encrypting / decrypting string data
type Encoder struct {
	Cipher cipher.Block
}

// NewEncoder returns a new instance of Encoder
func NewEncoder(key string) (encoder Encoder, err error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encoder = Encoder{Cipher: cipher}
	return
}

// Encrypt accepts plaintext and returns ciphertext
func (encoder *Encoder) Encrypt(str string) (result string, err error) {
	cipherText := make([]byte, aes.BlockSize+len(str))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	cfb := cipher.NewCFBEncrypter(encoder.Cipher, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(str))
	result = base64.StdEncoding.EncodeToString(cipherText)
	result = strings.Replace(result, "/", "$", -1)
	return
}

// Decrypt accepts encrypted ciphertext and returns decrypted plaintext
func (encoder *Encoder) Decrypt(str string) (result string, err error) {
	str = strings.Replace(str, "$", "/", -1)
	text, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	if len(text) < aes.BlockSize {
		err = ERR_CIPHERTEXT_TOO_SHORT
		return
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(encoder.Cipher, iv)
	cfb.XORKeyStream(text, text)
	result = string(text)
	return
}