package main

import (
	"fmt"

	r8 "github.com/bitfield/go-machine"
)

func main() {
	cpu := r8.NewCPU()
	fmt.Println("  PC")
	for {
		fmt.Printf("%04d >", cpu.PC)
		fmt.Scanln()
		cpu.Step()
	}
}
