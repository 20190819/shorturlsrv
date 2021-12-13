package utils

import "math/rand"

var defaultWords = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type str struct {
}

func (s str) Random(n int) string {
	tmp := make([]rune, n)

	for i := range tmp {
		key := rand.Intn(int(len(defaultWords)))
		tmp[i] = defaultWords[key]
	}
	return string(tmp)
}

var Str = new(str)
