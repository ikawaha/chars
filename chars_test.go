package chars

import (
	"fmt"
	"testing"
)

func TestRune_Name(t *testing.T) {
	cases := []rune{'😀'}

	for i, c := range cases {
		fmt.Println(i, Char(c).Name())
	}
}
