package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// EncryptCBC 使用 AES-CBC 模式对数据进行加密
func EncryptCBC(data, key, iv []byte) ([]byte, error) {
	// 创建一个新的 AES 加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = padPKCS7(data, block.BlockSize())

	// 使用 AES-CBC 加密模式进行加密
	ciphertext := make([]byte, len(data))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, data)

	return ciphertext, nil
}

// DecryptCBC 使用 AES-CBC 模式对数据进行解密
func DecryptCBC(data, key, iv []byte) ([]byte, error) {
	// 创建一个新的 AES 加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用 AES-CBC 解密模式进行解密
	plaintext := make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plaintext, data)
	plaintext = unpadPKCS7(plaintext)

	return plaintext, nil
}

// EncryptCBCBase64 加密后 转为base64字符串 返回
func EncryptCBCBase64(data, key, iv []byte) (string, error) {
	res, err := EncryptCBC(data, key, iv)
	if err != nil {
		return "", err
	}
	encodedCiphertext := base64.StdEncoding.EncodeToString(res)
	return encodedCiphertext, nil
}

// DecryptCBCBase64 传入base64字符串, 获取结果
func DecryptCBCBase64(data []byte, key, iv []byte) (string, error) {
	str, _ := base64.StdEncoding.DecodeString(string(data))
	res, err := DecryptCBC(str, key, iv)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func EncryptEcb(originData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData = padPKCS7(originData, block.BlockSize())
	secretData := make([]byte, len(originData))
	blockMode := newECBEncrypter(block)
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}

func DecryptEcb(secretData, key []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := newECBDecrypter(block)
	originByte = make([]byte, len(secretData))
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return unpadPKCS7(originByte), nil
}

// padPKCS7 使用 PKCS#7 填充对数据进行填充
func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// unpadPKCS7 使用 PKCS#7 填充对数据进行去除填充
func unpadPKCS7(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// newECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// newECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func newECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
