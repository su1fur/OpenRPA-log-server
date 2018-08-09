package tool

import (
	"github.com/pkg/errors"
)

func ParseStrFromInterface(i interface{}) (str string, err error) {
	if str, ok := i.(string); ok {
		return str, nil
	} else {
		err = errors.New("Not string")
		return "", err
	}
}
