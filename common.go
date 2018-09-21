package hash

import (
	"strings"
	"math"
)

type Hash struct {
	base string
}

func New(base string) Hash {
	return Hash{base}
}

func (m Hash) GetId(url string) int {
	result := 0.0
	i := len([]rune(url))
	k := len([]rune(m.base))
	val := strrev(url)
	for i > 0 {
		i --
		runeVal := string(val[i])
		stringIndex := strings.Index(m.base, runeVal)
		powValue := math.Pow(float64(k), float64(i))
		result = result + float64(stringIndex)*powValue
	}

	return int(result)
}

func (m Hash) GetString(integer int) string {
	out := ""
	length := len([]rune(m.base))
	for integer > length-1 {
		id := integer % length
		out = string(m.base[id]) + out
		integer = int(math.Floor(float64(integer / length)))
	}
	return string(m.base[integer]) + out
}

func strrev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
