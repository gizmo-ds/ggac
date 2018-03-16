package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strings"
	"time"
)

func get_google_auth(now time.Time, key string) int {
	return func(secret string, value int64) int {
		key, err := base32.StdEncoding.DecodeString(secret)
		if err != nil {
			return -1
		}

		hash := hmac.New(sha1.New, key)
		err = binary.Write(hash, binary.BigEndian, value)
		if err != nil {
			return -1
		}
		h := hash.Sum(nil)

		offset := h[19] & 0x0f

		truncated := binary.BigEndian.Uint32(h[offset : offset+4])

		truncated &= 0x7fffffff
		code := truncated % 1000000

		return int(code)
	}(
		strings.ToUpper(key),
		int64(now.UTC().Unix()/30),
	)
}
