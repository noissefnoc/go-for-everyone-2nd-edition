package main

import (
	"flag"
	"fmt"
	"strings"
)

type strSliceValue []string

// Value interface implementation
func (v *strSliceValue) Set(s string) error {
	strs := strings.Split(s, ",")
	*v = append(*v, strs...)
	return nil
}

func (v *strSliceValue) String() string {
	return strings.Join(([]string)(*v), ",")
}

func main() {
	// use original flag
	var species []string
	flag.Var((*strSliceValue)(&species), "species", "")

	// print
	fmt.Printf("%s\n", species)
}
