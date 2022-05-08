package utils

import "strconv"

func ParseFloat(str string) float64 {
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return f
	}
	return 0
}

func ParseUint(str string) uint64 {
	if u, err := strconv.ParseUint(str, 10, 64); err == nil {
		return u
	}
	return 0
}
