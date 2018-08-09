package tool

import (
	"encoding/json"
)

func ParseStringJSONFromInterface(i interface{}) (str string, err error) {

	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	str = ByteTostring(b)
	return str, nil
}

func ParseInterfaceFromStringJSON(str string) (i interface{}, err error) {

	b := StringToBytes(str)

	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
