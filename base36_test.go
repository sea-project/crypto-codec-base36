package base36

import (
	"testing"
)

func TestBase36Encode(t *testing.T) {
	str := Base36Encode([]byte("000"), UpperAlphabet)
	t.Log(str)

	res := Base36Decode(str, UpperAlphabet)
	t.Log(string(res))
}