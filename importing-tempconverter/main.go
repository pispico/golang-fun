package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pispico/tempconverter"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconverter.Fahrenheit(t)
		c := tempconverter.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconverter.FToC(f), c, tempconverter.CToF(c))

	}

}
