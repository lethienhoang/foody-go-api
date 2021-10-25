package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type Crypto struct {
	FakeID string
	ID     string
	DbType int
	Key []byte
}

func NewCrypto() Crypto {
	key := []byte("key id")
	return Crypto{
		Key: key,
	}
}

// encrypt string to base64 crypto using AES
func(e *Crypto) Encrypt(text string, dbType int) Crypto {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(e.Key)
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
	return Crypto{DbType: dbType, FakeID: base64.URLEncoding.EncodeToString(ciphertext), ID: text}
}

// decrypt from base64 to decrypted string
func(e *Crypto) Decrypt(cryptoText string, dbType int) Crypto {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(e.Key)
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

	return Crypto{DbType: dbType, FakeID: cryptoText, ID: fmt.Sprintf("%s", ciphertext)}
}
