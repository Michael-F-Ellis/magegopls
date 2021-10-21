//go:build mage

package main

import (
	"fmt"
	"os"
	"time"
)

// A trivial template for a .go file with a method that prints the date when it
// was generated.
var templ string = ` 
package mypkg

import (
	"fmt"
)
func DoIt() {
	fmt.Println("Generated at %v")
}
`

// Generate creates "pkg.go" using templ.
func Generate() (err error) {
	f, err := os.OpenFile("pkg.go", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Fprintf(f, templ, time.Now())
	return
}
