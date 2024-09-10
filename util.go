package testingutil

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"strings"
)

// [1, maxValue) random int
func randomInt(maxValue, step int) int {
	var buf [4]byte
	rnd.Read(buf[:])
	buf[0] &= 0x7f

	if step <= 1 {
		return 1 + int(binary.BigEndian.Uint32(buf[:]))%(maxValue-1)
	}

	const minValue = 1
	minValueStep := ((minValue-1)/step + 1) * step
	if minValueStep == 0 {
		minValueStep = step
	}

	maxValueStep := ((maxValue - 1) / step) * step

	count := (maxValueStep-minValueStep)/step + 1

	n := int(binary.BigEndian.Uint32(buf[:])) % count

	return minValueStep + n*step
}

func h(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	s = strings.TrimPrefix(s, "0x")
	for _, c := range s {
		if '0' <= c && c <= '9' {
			sb.WriteRune(c)
		} else if 'a' <= c && c <= 'f' {
			sb.WriteRune(c)
		} else if 'A' <= c && c <= 'F' {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}

// hex to *big.Int
func h2i(s string) *big.Int {
	s = h(s)
	if s == "" {
		return new(big.Int)
	}
	result, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic(s)
	}
	return result
}

// hex to byte
func h2b(s string) []byte {
	s = h(s)
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(s)
	}
	return b
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
