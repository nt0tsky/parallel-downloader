package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func CreateFileName(url string, etag string, contentLength int64) string {
	hash := sha256.New()
	hash.Write([]byte(url + etag + strconv.FormatInt(contentLength, 10)))
	fullHash := hash.Sum(nil)

	return hex.EncodeToString(fullHash)
}
