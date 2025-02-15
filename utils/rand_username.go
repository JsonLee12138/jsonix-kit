package utils

import (
	"math/rand"
	"strconv"
)

var (
	symbol    = []string{"_"}
	lowercase = genLettercase(false)
	uppercase = genLettercase(true)
	symbolLen = len(symbol)
)

func RandUsername(prefix string, len int, useSymbol bool) string {
	res := ""
	switch prefix {
	case "UPPER":
		res = randLetter(true)
	case "LOWER":
		res = randLetter(false)
	case "WORD":
		res = randLetterAll()
	}
	remainLen := len - 1
	if useSymbol {
		res = res + randSymbol()
		remainLen -= 1
	}
	for i := 0; i < remainLen; i++ {
		res = res + randLetterOrNumber()
	}
	return res
}

func genLettercase(isUpper bool) []string {
	res := make([]string, 26)
	if isUpper {
		for i := 0; i < 26; i++ {
			res[i] = string(rune('A' + i))
		}
	} else {
		for i := 0; i < 26; i++ {
			res[i] = string(rune('a' + i))
		}
	}
	return res
}

type RandIntProps struct {
	min int
	max int
}

func RandInt(props *RandIntProps) int {
	return rand.Intn(props.max-props.min+1) + props.min
}

func randLetter(isUpper bool) string {
	index := RandInt(&RandIntProps{min: 0, max: 25})
	if isUpper {
		return uppercase[index]
	}
	return lowercase[index]
}

func randSymbol() string {
	return symbol[RandInt(&RandIntProps{max: symbolLen - 1, min: 0})]
}

func randLetterOrNumber() string {
	_type := RandInt(&RandIntProps{min: 0, max: 2})
	switch _type {
	case 0:
		return randLetter(true)
	case 1:
		return randLetter(false)
	case 2:
		return strconv.Itoa(RandInt(&RandIntProps{min: 0, max: 9}))
	}
	return ""
}

func randLetterAll() string {
	index := RandInt(&RandIntProps{min: 0, max: 51})
	if index > 25 {
		return uppercase[index%26]
	}
	return lowercase[index]
}
