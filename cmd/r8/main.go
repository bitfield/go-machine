package main

import (
	"fmt"
	"os"
	"strconv"

	r8 "github.com/bitfield/go-machine"
)

func main() {
	var program []byte
	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseUint(arg, 10, 8)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		program = append(program, byte(value))
	}
	cpu := r8.NewCPU()
	copy(cpu.Mem[:], program)
	for {
		code := cpu.Mem[cpu.PC : cpu.PC+3]
		asm := r8.Disassemble(code)
		fmt.Println("  PC  A | NEXT")
		fmt.Printf("%04d %02d | %s\n> ", cpu.PC, cpu.A, asm)
		fmt.Scanln()
		cpu.Step()
	}
}
