package utils

import (
	"strconv"
	"strings"
	"time"
)

func GetTmeStampInMs() int64 {
	return time.Now().UnixMilli()
}

// convert string to int64
func StringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// check if a string contains a substring
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}
