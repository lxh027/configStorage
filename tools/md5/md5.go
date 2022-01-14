package md5

import (
	"configStorage/tools/random"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func GetMd5(str string) string {
	base64String := base64.StdEncoding.EncodeToString([]byte(str))
	hash := md5.Sum([]byte(base64String))
	hashBase64 := fmt.Sprintf("%x", hash)
	return hashBase64
}

func GetRandomMd5() string {
	return GetMd5(random.RandString(16))
}
