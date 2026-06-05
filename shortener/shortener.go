package shortener

import (
	"crypto/sha256"
	"encoding/hex"
)

func ShortUrl(url string, size uint) *string {
	if size > sha256.Size && size != 0 {
		return nil
	}
	hash := sha256.Sum256([]byte(url))
	hash_hex := hex.EncodeToString(hash[:])
	slice := hash_hex[:size]
	return &slice
}
