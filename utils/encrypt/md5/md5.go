package md5Encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密32位小写字母
func Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
