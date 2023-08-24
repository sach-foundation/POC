package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate Random int between max and min
func RandomInt(min, max int64) int64 {
	// return min + rand.Int63n(max-min+1)
	return min + rand.Int63n(max-min+1)

}

//Random string generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//return random owner
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currency := []string{"INR", "EUR", "USD", "CAD"}
	n := len(currency)
	return currency[rand.Intn(n)]

}
