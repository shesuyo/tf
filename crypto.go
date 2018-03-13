package wtf

import (
	"crypto/md5"
	"fmt"
)

// CryproMD5 返回字符串小写的MD5
func CryproMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}


