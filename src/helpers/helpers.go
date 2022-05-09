package helpers

import (
	"math/rand"
	"strings"
)

func RandomUser() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CleanUpStringsArray(strArray []string) []string {
	cleanArray := make([]string, 0)

	for _, str := range strArray {
		cleanStr := strings.ReplaceAll(str, "\n", "")
		cleanStr = strings.TrimSpace(cleanStr)
		cleanArray = append(cleanArray, cleanStr)

	}

	return cleanArray
}

func CleanUpString(str string) string {
	cleanStr := strings.Replace(str, "\\", "", -1)
	cleanStr = strings.ReplaceAll(cleanStr, ",", " ")
	cleanStr = strings.TrimSpace(cleanStr)
	return cleanStr
}

func IsAbsoluteUrl(url string) bool {
	return strings.Contains(url, "http") || strings.Contains(url, "https")
}
