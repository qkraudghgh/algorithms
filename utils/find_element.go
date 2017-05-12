package utils

import "errors"

func FindInt(list []int, element int) (int, error) {
	for index, value := range list {
		if value == element {
			return index, nil
		}
	}
	return -1, errors.New("Did not find Element.")
}

func FindStrByByteInMap(str string, dic map[byte]string) bool {
	for _, value := range dic {
		if str == value {
			return true
		}
	}
	return false
}
