package base36

import (
	"errors"
	"math/big"
	"strings"
)

var (
	UpperAlphabet      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerAlphabet      = "0123456789abcdefghijklmnopqrstuvwxyz"
	errICAPEncoding    = errors.New("invalid ICAP encoding")
	errICAPChecksum    = errors.New("invalid ICAP checksum")
)

var (
	Big1  = big.NewInt(1)
	Big0  = big.NewInt(0)
	Big36 = big.NewInt(36)
	Big97 = big.NewInt(97)
	Big98 = big.NewInt(98)
)

// Base36Encode
func Base36Encode(data []byte, alphabet string) string {
	i := new(big.Int).SetBytes(data)
	var chars []rune
	x := new(big.Int)
	for {
		x.Mod(i, Big36)
		chars = append(chars, rune(alphabet[x.Uint64()]))
		i.Div(i, Big36)
		if i.Cmp(Big0) == 0 {
			break
		}
	}

	// reverse slice
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// Base36Decode
func Base36Decode(s, alphabet string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	for i := len(s) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(alphabet, string(s[i]))
		if tmp == -1 {
			return []byte("")
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)

		answer.Add(answer, tmp1)
		j.Mul(j, Big36)
	}
	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(s); numZeros++ {
		if s[numZeros] != alphabet[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)
	return val
}
