package tf

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// CryproMD5 返回字符串小写的MD5
func CryproMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// CryproSHA1 返回字符串小写的SHA1
func CryproSHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

// CryproSHA1Bytes 返回字符串小写的SHA1
func CryproSHA1Bytes(bs []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(bs))
}
