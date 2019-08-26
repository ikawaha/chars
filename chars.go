package chars

import (
	"fmt"
	"strconv"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/text/unicode/runenames"
)

type Char rune

func New(hex string) Char {
	code, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return '\uFFFD'
	}
	return Char(code)
}

func (c Char) CodePoint() string {
	return fmt.Sprintf("U+%02X",c)
}

func (c Char) UTF8() []string {
	buf :=make([]byte, utf8.UTFMax)
	n:=utf8.EncodeRune(buf, rune(c))
	ret := make([]string, 0, n)
	for i := range  buf[:n] {
		ret = append(ret, fmt.Sprintf("%02X", buf[i]))
	}
	return ret
}

func (c Char) UTF16() []string{
	v := utf16.Encode([]rune{rune(c)})
	return []string{fmt.Sprintf("%04X",v[0])}
}

func (c Char) Name() string {
	return runenames.Name(rune(c))
}