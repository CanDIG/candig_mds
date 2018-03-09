package dummydata

import (
	"math/rand"
	"strconv"
)

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func makeRandomString() string {
	num := rand.Intn(50)
	value := randSeq(num)
	return value
}

func makeRandomDate() string {
	date := strconv.Itoa(rand.Intn(118)+1900) + "-" + strconv.Itoa(rand.Intn(11)+1) + "-" + strconv.Itoa(rand.Intn(27)+1)
	return date
}

func makeRandomFloat() float32 {
	num := rand.Float32()
	num += float32(rand.Intn(600))
	return num
}

func makeRandomBool() bool {
	num := rand.Intn(1)
	if num == 1 {
		return true
	}
	return false
}
