package tool

import (
	"strconv"
)

func ParseUint64FromStr(str string) (uint64, error) {
	number, parseErr := strconv.ParseUint(str, 10, 64)
	if parseErr != nil {
		return 0, parseErr
	}
	return number, nil
}

func ParseUintFromStr(str string) uint {
	number, parseErr := strconv.Atoi(str)
	if parseErr != nil {
		panic(parseErr)
	}
	return uint(number)
}