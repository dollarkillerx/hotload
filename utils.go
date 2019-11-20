package hotload

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// 计算指定文件的hash值
func md5File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
