package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type MysqlAES struct {
	key []byte
}

func NewMysqlAES(key string) *MysqlAES {
	return &MysqlAES{key: []byte(key)}
}

func fillZero(length int) []byte {
	return bytes.Repeat([]byte{0}, length)
}

func initialIV(secret []byte, length int) []byte {
	iv := fillZero(length)
	for index, v := range secret {
		iv[index%length] ^= v
	}
	return iv
}

func newBlock(secret []byte) (cipher.Block, error) {
	return aes.NewCipher(initialIV(secret, aes.BlockSize))
}

func (maes *MysqlAES) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, nil
	}

	block, err := newBlock(maes.key)
	if err != nil {
		return nil, err
	}
	paddingCount := len(plaintext) % block.BlockSize()
	paddingCount = block.BlockSize() - paddingCount
	plaintext = append(plaintext, bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)...)

	totalLength := len(plaintext)

	cipherBytes := make([]byte, totalLength, totalLength)
	for i := 0; i < totalLength/block.BlockSize(); i++ {
		startIndex := i * block.BlockSize()
		endIndex := startIndex + block.BlockSize()
		block.Encrypt(cipherBytes[startIndex:endIndex], plaintext[startIndex:endIndex])
	}
	return cipherBytes, nil
}

func (maes *MysqlAES) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, nil
	}

	block, err := newBlock(maes.key)
	if err != nil {
		return nil, err
	}

	totalLength := len(ciphertext)

	plaintext := make([]byte, totalLength, totalLength)
	for i := 0; i < totalLength/block.BlockSize(); i++ {
		startIndex := i * block.BlockSize()
		endIndex := startIndex + block.BlockSize()
		block.Decrypt(plaintext[startIndex:endIndex], ciphertext[startIndex:endIndex])
	}

	paddingSize := plaintext[len(plaintext)-1]
	plaintext = plaintext[0 : len(plaintext)-int(paddingSize)]

	return plaintext, nil
}
