package utils

import (
	"math/rand"
	"time"
)

type StrGenFunc func() string

func RandomStringGenerator() StrGenFunc {
	return func() string {
		rand.Seed(time.Now().Unix())
		length := 6

		ranStr := make([]byte, length)

		// Generating Random string
		for i := 0; i < length; i++ {
			ranStr[i] = ranStr[65+rand.Intn(25)]
		}

		// Displaying the random string
		str := string(ranStr)
		return str
	}
}
