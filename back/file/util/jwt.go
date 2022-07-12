package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		SecretKey: secretKey,
	}
}

const DownloadFileKey = "MDa8NSNQRcaZZnZO"

func (j *JWT) GenJWTToken(sub string, expire time.Duration) (string, error) {
	clm := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + int64(expire.Seconds()),
		Subject:   sub,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) DecodeJwtToken(token string) (*jwt.StandardClaims, error) {
	clm := &jwt.StandardClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, clm, j.keyFunc)
	if err != nil {
		return nil, err
	}

	clm2, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("not StandardClaims")
	}

	return clm2, nil
}

func (j *JWT) ValidateToken(token string) bool {
	jwtToken, err := j.DecodeJwtToken(token)
	if err != nil {
		return false
	}
	return jwtToken.Valid() == nil
}

func (j *JWT) keyFunc(token *jwt.Token) (i interface{}, err error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		return []byte(j.SecretKey), nil
	} else {
		return nil, fmt.Errorf("expect token signed with HMAC but got %v", token.Header["alg"])
	}
}

func AesEncryptCBC(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()                  // 获取秘钥块的长度
	origData = PKCS5Padding(origData, blockSize)    // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key) // 加密模式
	encrypted := make([]byte, len(origData))        // 创建数组
	blockMode.CryptBlocks(encrypted, origData)      // 加密
	return encrypted, nil
}

func AesDecryptCBC(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted := make([]byte, len(encrypted))                   // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	return PKCS5UnPadding(decrypted)                            // 去除补全码
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	return PKCS5UnPadding(origData)
}

func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length-unpadding < 0 {
		return nil, fmt.Errorf("illeagl data")
	}
	return origData[:(length - unpadding)], nil
}
