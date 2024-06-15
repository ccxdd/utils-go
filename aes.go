package utils_go

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if unpadding > length {
		return []byte{}
	}
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[len(key)-blockSize:])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[len(key)-blockSize:])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// AesCbcZeroEncrypt 使用 AES 算法对数据进行加密，并返回加密后的字符串。
func AesCbcZeroEncrypt(data, key, iv []byte) (string, error) {
	// 创建一个新的 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 获取密码块的大小
	blockSize := block.BlockSize()

	// 如果输入数据的长度不是密码块大小的倍数，则进行填充
	plaintext := data
	plaintextLength := len(plaintext)
	if plaintextLength%blockSize != 0 {
		plaintextLength = plaintextLength + (blockSize - (plaintextLength % blockSize))
	}
	paddedText := make([]byte, plaintextLength)
	copy(paddedText, plaintext)

	// 创建一个新的字节切片用于存储加密后的数据
	cipherText := make([]byte, len(paddedText))

	// 创建一个新的 CBC 加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 对填充后的数据进行加密
	mode.CryptBlocks(cipherText, paddedText)

	// 将加密后的数据转换为 Base64 编码的字符串并返回
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// AesCbcZeroDecrypt 使用给定的密钥和初始向量对数据进行 AES 解密。
func AesCbcZeroDecrypt(data string, key string, iv string) (string, error) {
	// 将输入的字符串进行Base64解码
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	// 创建AES算法实例
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 使用密钥和初始化向量初始化Cipher对象
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))

	// 解密数据
	decryptedData := make([]byte, len(decodedData))
	blockMode.CryptBlocks(decryptedData, decodedData)

	// 去除填充
	padding := int(decryptedData[len(decryptedData)-1])
	decryptedData = decryptedData[:len(decryptedData)-padding]

	// 将解密后的字节数组转换为字符串并返回
	return string(decryptedData), nil
}
