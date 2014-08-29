package goloc

import "testing"

func TestCountStatements(t *testing.T) {
	for _, tt := range testCases {
		got, err := CountStatements(tt.input)
		if err != nil && !tt.wantErr {
			t.Fatalf("%q case got error: %v, want <nil>", tt.name, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("%q case got count = %d, want %d", tt.name, got, tt.want)
		}
	}
}

var testCases = []struct {
	name    string
	input   string
	want    int
	wantErr bool
}{
	// test file with basic functionality
	{
		name: "Basic file",
		input: `// main does some things
package main

import (
	"fmt"
	"strings"
	"nice"
)

// main does this amazing
// thing, really
// wow
func main() {
	fmt.Println("Hello, World!", strings.Split("123", ""))
}`,
		want:    1,
		wantErr: false,
	},

	// test file with no statements
	{
		name: "No statements",
		input: `/*
	Package test is a test for this file
	*/
	package test

	import (
		"fmt"
		"strings"
	)

	// theFunc is just a placeholder
	func theFunc() {
	}`,
		want:    0,
		wantErr: false,
	},

	// test broken file
	{
		name: "No statements",
		input: `for i := 0; i < 10; i++ {
				fmt.Println("test", i)
				theFunc()
	}`,
		want:    0,
		wantErr: true,
	},

	// test file with range statement and more
	{
		name: "More complex file",
		input: `package main

	import (
		"fmt"
		"strings"
	)

	func main() {
		fmt.Println("Starting now")
		a := []int{1,2,3}
		for i := range a {
			fmt.Println(i, a[i])
		}
	}`,
		want:    4,
		wantErr: false,
	},
}
