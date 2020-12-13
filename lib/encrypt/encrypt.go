package encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
