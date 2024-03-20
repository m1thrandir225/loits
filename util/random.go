package util

import (
	"fmt"
	db "m1thrandir225/loits/db/sqlc"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomMagicianName - generate a random magician name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomElement() db.Element {
	elements := []db.Element{
		db.ElementEarth,
		db.ElementFire,
		db.ElementWind,
		db.ElementWater,
		db.ElementMetal,
	}
	n := len(elements)
	return elements[rand.Intn(n)]
}

func RandomMagicRating() string {
	ratings := []string{
		"S", "A", "B", "C", "D", "F",
	}

	n := len(ratings)

	return ratings[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomDate() time.Time {
	min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.Local).Unix()
	max := time.Date(2020, 12, 0, 0, 0, 0, 0, time.Local).Unix()

	sec := RandomInt(min, max)
	return time.Unix(sec, 0)
}
