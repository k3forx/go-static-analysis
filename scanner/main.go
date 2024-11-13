package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

// https://go.dev/play/p/t3XAEtYYP7e も参考になる
// src := []byte(`sum := n + 1`)

// fset := token.NewFileSet()
// file := fset.AddFile("tmp.go", fset.Base(), len(src))

// var s scanner.Scanner
// s.Init(file, src, nil, scanner.Mode(0))
// for {
// 	pos, tok, lit := s.Scan()
// 	if tok == token.EOF {
// 		break
// 	}
// 	fmt.Printf("%+v\t%s\t%q\n", fset.Position(pos), tok, lit)
// }

func main() {
	f, err := os.Open("a.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	b, _ := io.ReadAll(f)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile(f.Name(), fset.Base(), len(b))
	s.Init(file, b, nil, scanner.ScanComments)

	var (
		oldLine = 1
		tokens  = []string{}
	)
	for {
		pos, tok, _ := s.Scan()
		if tok == token.EOF {
			fmt.Printf("Line: %d: %s\n", oldLine, strings.Join(tokens, " "))
			break
		}
		currentLine := fset.Position(pos).Line
		if currentLine != oldLine {
			fmt.Printf("Line: %d: %s\n", oldLine, strings.Join(tokens, " "))
			oldLine = currentLine
			tokens = []string{}
		}
		tokens = append(tokens, tok.String())
	}
}
