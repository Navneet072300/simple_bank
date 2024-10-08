package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
const alphabte = "abcdefghijklmnopqrstuvwxyz"
const pass = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$&*+=?<>"


func RandomInt(min, max int64) int64 {
	return min + rnd.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabte[rnd.Intn(len(alphabte))])
		}
		return sb.String()
}
func RandomPassword(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabte[rnd.Intn(len(pass))])
		}
		return sb.String()
}

func RandomOwner() string{
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies :=  []string{USD, IND, EUR}
	return currencies[rnd.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}