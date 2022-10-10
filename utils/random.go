package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alp = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alp)
	for i := 0; i < n; i++ {
		c := alp[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return randomInt(0, 1000)
}

func RandomCurrency() string  {
	currencies := []string{"EUR", "USD", "IDR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

