package utils

import "math/rand"

func GetRandomString(num int) string {
	const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := make([]byte, num)
	for i := range str {
		str[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(str)
}
