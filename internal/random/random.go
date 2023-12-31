package random

import (
	"math/rand"
	"time"

	"github.com/haquenafeem/shrinkie/internal/consts"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890=#$")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomStringDefualt() string {
	return RandomString(consts.DEFAULT_RANDOM_STRING_LENGTH)
}
