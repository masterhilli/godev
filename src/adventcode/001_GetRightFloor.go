package getrightfloor

import (
	"strings"
)

func IsOpeningBracket(bracket byte) bool {
	if bracket == byte('(') {
		return true
	} else {
		return false
	}
}

func RetrieveFloor(brackets string) int {
	retVal := strings.Count(brackets, "(")
	cnt := strings.Count(brackets, ")")
	return retVal - cnt
}

func RetrievePosOfFirstBasement(brackets string) int {
	var bytes []byte = []byte(brackets)
	var retVal int = 0
	var i int = 0

	for i = range bytes {
		if IsOpeningBracket(bytes[i]) {
			retVal++
		} else {
			retVal--
		}

		if retVal < 0 {
			return i + 1
		}
	}
	return i + 1
}
