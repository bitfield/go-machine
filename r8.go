// Package r8 emulates a simple CPU called the R8.
package r8

import "fmt"

type CPU struct {
	PC     int
	Mem [256]int
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	opcode := cpu.Mem[cpu.PC]
	cpu.PC++
	switch opcode {
	case 1:
		// nothing to do
	default:
		panic(fmt.Sprintf("unimplemented opcode %d", opcode))
	}
}
