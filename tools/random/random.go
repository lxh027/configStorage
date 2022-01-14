package random

import (
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

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
