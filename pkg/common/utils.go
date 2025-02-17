package common

import (
	"errors"
	"strconv"
)


type UtilsT struct{}

var Utils UtilsT



func (UtilsT) ToInteger(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("invalid integer: " + str)
	}
	return num, nil
}