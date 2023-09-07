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
