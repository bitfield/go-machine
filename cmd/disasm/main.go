package main

import (
	"fmt"
	"os"

	r8 "github.com/bitfield/go-machine"
)

var usage = "Usage: disasm FILE"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}
	program, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var asm string
	var bytes int
	for i := 0; i < len(program); i += bytes {
		asm, bytes = r8.Disassemble(program[i:])
		fmt.Println(asm)
	}
}
