package utils

import (
	"strings"
)

func GetQueryString(url string) string {
	if idx := strings.LastIndex(url, "?"); idx > -1 {
		idx += 1
		return string(url[idx:])
	}

	return ""
}
