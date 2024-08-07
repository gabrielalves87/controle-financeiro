package utils

import (
	"math/rand"
	"fmt"
	"strings"

)

const alphabet = "abcdefghijlmnopqrstuvwxyz"


func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
	
}




