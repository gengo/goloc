package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gengo/goloc"
)

var usageTemplate = `
Goloc is a tool for counting the number of statements in go files.

Usage:

	cat path/to/file | goloc

`

func main() {
	flag.Usage = usageExit
	flag.Parse()

	bio := bufio.NewReader(os.Stdin)
	lines := []string{}
	line, err := bio.ReadString('\n')
	for ; err == nil; line, err = bio.ReadString('\n') {
		lines = append(lines, line)
	}
	src := strings.Join(lines, "")

	stmts, err := goloc.CountStatements(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(2)
	}
	fmt.Println(stmts)
}

func usageExit() {
	fmt.Fprint(os.Stderr, usageTemplate)
	os.Exit(2)
}
