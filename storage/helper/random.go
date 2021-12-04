package helper

import (
	"math/rand"
	"time"
)

// RandomTimeout to return a Timeout including fix and rand duration
func RandomTimeout(fixDuration time.Duration, randDuration time.Duration) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if randDuration == 0 {
		return fixDuration
	}
	d := r.Int63n(randDuration.Milliseconds())
	return time.Duration(d)*time.Millisecond + fixDuration
}
