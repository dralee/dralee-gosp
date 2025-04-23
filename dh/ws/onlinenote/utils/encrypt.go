/*
加密

2025.4.23 by dralee
*/
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

const (
	EncryptKey = "DraleeOnlineNote#2025abcd12378fg"
)

// AES加密，返回base64
func AesEncrypt(plainText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// AES解密
func AesDecrypt(cipherTextBase64, key []byte) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(string(cipherTextBase64))
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText) < aes.BlockSize {
		return nil, err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}
