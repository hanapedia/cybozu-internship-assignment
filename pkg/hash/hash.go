package hash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/hanapedia/cybozu-internship-assignment/pkg/models"
)

// Hash hashes the lines
func Hash(job models.Job) string {
	h := sha256.New()
	h.Write([]byte(job.Line))
	return hex.EncodeToString(h.Sum(nil))
}

// NoHash is only for testing to see if the order is preserved.
func NoHash(job models.Job) string {
	return job.Line
}
