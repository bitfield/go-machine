package main

import (
	"fmt"
	"os"

	r8 "github.com/bitfield/go-machine"
)

var usage = "Usage: r8 FILE"

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
	cpu := r8.NewCPU()
	copy(cpu.Mem[:], program)
	var cmd string
	for {
		code := cpu.Mem[cpu.PC : cpu.PC+3]
		asm, _ := r8.Disassemble(code)
		fmt.Println("  PC  A | NEXT")
		fmt.Printf("%04d %02d | %s\n> ", cpu.PC, cpu.A, asm)
		fmt.Scanln(&cmd)
		if cmd == "g" {
			cpu.Run()
		} else {
			cpu.Step()
		}
	}
}
