package utils

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
)

// GetServerID returns a random hex number string that serves
// as an identifier for the server that the request is being
// handled by
func GetServerID() string {
	if num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64)); err == nil {
		return strconv.FormatUint(num.Uint64(), 16)
	}

	return ""
}
