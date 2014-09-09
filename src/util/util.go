package util

import (
	"crypto/md5"
)

func GetMD5(str string) string {
	res := md5.Sum([]byte(str))
	return string(res[:])
}
