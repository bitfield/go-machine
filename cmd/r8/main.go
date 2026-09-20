package main

import (
	"fmt"

	r8 "github.com/bitfield/go-machine"
)

func main() {
	cpu := r8.NewCPU()
	fmt.Println("  PC  A")
	for {
		fmt.Printf("%04d %02d >", cpu.PC, cpu.A)
		fmt.Scanln()
		cpu.Step()
	}
}
