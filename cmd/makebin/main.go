package main

import (
	"fmt"
	"os"
	"strconv"
)

var usage = "Usage: makebin FILE CODE"

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	var program []byte
	for _, arg := range os.Args[2:] {
		value, err := strconv.ParseUint(arg, 10, 8)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		program = append(program, byte(value))
	}
	err := os.WriteFile(os.Args[1], program, 0o600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
