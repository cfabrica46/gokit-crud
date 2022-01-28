package handler

import "crypto/sha256"

func newSha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
