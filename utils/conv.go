package utils

import "strconv"

func StringToUint(s string) uint {
	n, _ := strconv.Atoi(s)
	return uint(n)
}
