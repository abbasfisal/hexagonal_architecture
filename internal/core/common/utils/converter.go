package utils

import "strconv"
import "github.com/google/uuid"

func ConvertUInt64ToString(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func GetUUID() string {
	return uuid.New().String()
}
