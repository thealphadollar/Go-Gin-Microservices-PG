package helpers

import (
	"math/rand"
	"strconv"
)

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
