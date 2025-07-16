package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var Rsa = &RsaHandler{}

type RsaHandler struct {
	pubkey *rsa.PublicKey  //公钥
	prikey *rsa.PrivateKey //私钥
}

// 设置公钥
func (r *RsaHandler) SetPublicKey(pub string) (err error) {
	r.pubkey, err = getPubKey([]byte(pub))
	return err
}

// 设置私钥
func (r *RsaHandler) SetPrivateKey(pri string) (err error) {
	r.prikey, err = getPriKey([]byte(pri))
	return err
}

// RsaSignWithSha256 私钥签名
func (r *RsaHandler) RsaSignWithSha256(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, r.prikey, crypto.SHA256, hashed)
	if err != nil {
		return []byte{}, err
	}

	return signature, nil
}

// RsaVerySignWithSha256 公钥验证
func (r *RsaHandler) RsaVerySignWithSha256(data, signData []byte) (bool, error) {
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(r.pubkey, crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
		return false, err
	}
	return true, nil
}

// RsaEncrypt 公钥加密
// 结果一般都需要 转成base64字符串 base64.StdEncoding.EncodeToString(res)
func (r *RsaHandler) RsaEncrypt(data []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, r.pubkey, data)
	if err != nil {
		return []byte{}, err
	}
	return ciphertext, nil
}

// RsaDecrypt 私钥解密
func (r *RsaHandler) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	data, err := rsa.DecryptPKCS1v15(rand.Reader, r.prikey, ciphertext)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
