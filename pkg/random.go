package pkg

import "math/rand"

var letters = []rune("abcde")

func RandomGenerate(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))] // なんだこれは？
	}
	return string(b)
}
