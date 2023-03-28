package utils

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	tmp := make([]byte, n)
	for i := range tmp {
		tmp[i] = letters[rand.Intn(len(letters))]
	}
	return string(tmp)
}
