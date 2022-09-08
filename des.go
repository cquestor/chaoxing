package cx

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

func desEncrypt(data, key []byte) (string, error) {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", errors.New("desEncrypt newCipher error")
	}
	bs := block.BlockSize()
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("desEncrypt need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil

}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
