package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"math"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	bys, err := DesEncrypt([]byte("Hrbr12@Tensor.*#)"), []byte("talkerss"))
	fmt.Println("hello", err)
	d1, err := DesDecrypt(bys, []byte("talkerss"))

	fmt.Println(string(d1), err)
}

type RejectReasonMap struct {
	En map[int64]string `json:"en"`
	Zh map[int64]string `json:"zh"`
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
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// GetMixedSetForString 取交集，但是但一个为空时，就返回另一个集合，而不是返回空
func GetMixedSetForString(pre, after []string) []string {
	res := make([]string, 0)
	if len(pre) == 0 {
		return after
	}
	if len(after) == 0 {
		return pre
	}
	preMap := make(map[string]int)
	for _, p := range pre {
		preMap[p] = 1
	}
	for _, p := range after {
		if preMap[p] == 1 {
			res = append(res, p)
			preMap[p]++ // 去重
		}
	}
	return res
}

// GetMixedSetForInt64 取交集，但是但一个为空时，就返回另一个集合，而不是返回空
func GetMixedSetForInt64(pre, after []int64) []int64 {
	pre = append(pre, after...)

	preMap := make(map[int64]int)
	res := make([]int64, 0)
	for _, p := range pre {
		if preMap[p] < 1 {
			res = append(res, p)
			preMap[p]++ // 去重
		}
	}
	return res
}

func MinInt(res ...int) int {
	if len(res) == 0 {
		return 0
	}
	ans := math.MaxInt64
	for _, r := range res {
		if r < ans {
			ans = r
		}
	}
	return ans
}

func DeepCopy(dst, src interface{}) error {
	eventData, err := jsoniter.Marshal(src)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(eventData, &dst)
}
