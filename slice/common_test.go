package slice

import "strings"

func isPositive(i int) (bool, error) {
	return i >= 0, nil
}

func isNegative(i float64) (bool, error) {
	return i < 0, nil
}

func isUppercase(s string) (bool, error) {
	return s == strings.ToUpper(s), nil
}
