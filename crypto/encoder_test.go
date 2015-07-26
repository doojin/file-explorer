package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"crypto/aes"
)

func Test_NewEncoder_ShouldNotReturnErrorIfKeySizeIsValid(t *testing.T) {
	key1 := "1234567890123456" // 16 symbols
	key2 := "123456789012345678901234" // 24 symbols
	key3 := "12345678901234567890123456789012" // 32 symbols

	_, err1 := NewEncoder(key1)
	_, err2 := NewEncoder(key2)
	_, err3 := NewEncoder(key3)

	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
	assert.Equal(t, nil, err3)
}


func Test_NewEncoder_ShouldReturnErrorIfKeySizeInInvalid(t *testing.T) {
	key := "12345"

	_, err := NewEncoder(key)

	assert.Equal(t, aes.KeySizeError(5), err)
}

func Test_Encrypt_ShouldReturnEncryptedTextOfValidSize(t *testing.T) {
	encoder, err := NewEncoder("123456789012345678901234")
	text := "randomText"

	assert.Equal(t, nil, err)

	encryptedText, err := encoder.Encrypt(text)

	assert.Equal(t, nil, err)
	assert.Equal(t, 36, len(encryptedText))
}

func Test_Decrypt_ShouldDecryptEncryptedTextCorrectly(t *testing.T) {
	encoder, err := NewEncoder("123456789012345678901234")
	text := "randomText"

	assert.Equal(t, nil, err)

	encryptedText, err := encoder.Encrypt(text)

	assert.Equal(t, nil, err)

	decryptedText, err := encoder.Decrypt(encryptedText)

	assert.Equal(t, nil, err)
	assert.Equal(t, "randomText", decryptedText)
}

func Test_Decrypt_ShouldReturnErrorIfEncryptedTextHasWrongSize(t *testing.T) {
	encoder, err := NewEncoder("123456789012345678901234")

	assert.Equal(t, nil, err)

	decryptedText, err := encoder.Decrypt("tooShortText")

	assert.Equal(t, "", decryptedText)
	assert.Equal(t, ERR_CIPHERTEXT_TOO_SHORT, err)
}