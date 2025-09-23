package helper

// #nosec

import (
	cryptoRand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"

	"gin-starter/common/constant"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// GenerateRandomNumber generates a random number between minimum (inclusive) and maximum (exclusive).
func GenerateRandomNumber(minimum, maximum int) int {
	// Create a new random source seeded with the current time
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404

	return r.Intn(maximum-minimum) + minimum // #nosec
}

// GenerateTrxID generates a transaction ID with a prefix and random number.
func GenerateTrxID(prefix string) string {
	// Create a new random generator seeded with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404

	// Generate random number in given range
	num := r.Intn(constant.NinetyNineHundred-constant.Hundred) + constant.TenThousand

	// Build transaction ID
	return fmt.Sprintf("%s%s/%d", prefix, time.Now().Format("20060102"), num)
}

// GenerateExternalID generate external ID. Commonly used for third party payment
func GenerateExternalID(prefix string) string {
	res := prefix + fmt.Sprint(time.Now().Unix())

	return res
}

// GenerateOTP generate OTP number for auth
func GenerateOTP(maximum int) string {
	b := make([]byte, maximum)
	n, err := io.ReadAtLeast(cryptoRand.Reader, b, maximum)
	if n != maximum {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes generate random string by bytes
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))] // #nosec
	}
	return string(b)
}
