package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyLen    = 32
	saltBytes = 32
	iter      = 4096
)

// aesCryp implements cipher
type aesCryp struct{}

// generate key for AES
func (a *aesCryp) keyGen(pass []byte) (key []byte, salt []byte, err error) {
	salt = make([]byte, saltBytes)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, nil, err
	}
	key = pbkdf2.Key(pass, salt, iter, keyLen, sha1.New)
	return key, salt, nil
}

// generate key with salt for AES
func (a *aesCryp) keyGenWithSalt(pass []byte, salt []byte) (key []byte) {
	return pbkdf2.Key(pass, salt, iter, keyLen, sha1.New)
}

// encrypt string to base64 crypto using AES
func (a *aesCryp) encrypt(key, text []byte) (string, error) {
	// key := []byte(keyText)
	plaintext := text

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// decrypt from base64 to decrypted string
func (a *aesCryp) decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("cipher to short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
