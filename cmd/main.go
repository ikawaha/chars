package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ikawaha/chars"
)

func main() {
	var r io.Reader = os.Stdin
	if len(os.Args) >1 {
		r = strings.NewReader(os.Args[1])
	}
	if err :=Run(r); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Run(r io.Reader) error {
	b :=bufio.NewReader(r)
	for {
		c, _, err := b.ReadRune()
		if err != nil{
			if err == io.EOF {
				return nil
			}
			return err
		}
		char := chars.Char(c)
		fmt.Printf("%c\t%v\t%v\t%v\t%v\n", char, char.CodePoint(), char.UTF8(), char.UTF16(), char.Name())
	}
	return nil
}
