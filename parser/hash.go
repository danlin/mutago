package parser

import (
	"io"
	"crypto/md5"
	"fmt"
)

func hash(reader io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	var result []byte
	return fmt.Sprintf("%x", hash.Sum(result)) , nil
}