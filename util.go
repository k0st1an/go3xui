package go3xui

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func RandomHexString(length int) string {
	bytes := make([]byte, (length+1)/2)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	hexStr := hex.EncodeToString(bytes)
	return hexStr[:length]
}

func GenerateShortIds() []string {
	lengths := []int{4, 8, 2, 6, 10, 8, 14, 16}
	result := make([]string, len(lengths))
	for i, l := range lengths {
		result[i] = RandomHexString(l)
	}
	return result
}
