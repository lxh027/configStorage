package random

import (
	"math/rand"
	"time"
)

func Timeout(fixDuration time.Duration, randDuration time.Duration) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if randDuration == 0 {
		return fixDuration
	}
	d := r.Int63n(randDuration.Milliseconds())
	return time.Duration(d)*time.Millisecond + fixDuration
}

func ID(idRange int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(idRange - 1)
}
