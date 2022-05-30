package utils

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max, seed int) int {
	rand.Seed(time.Now().UnixNano() + int64(seed))
	return rand.Intn(max-min+1) + min
}
