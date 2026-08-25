// Package r8 emulates a simple CPU called the R8.
package r8

import "fmt"

const OpNOP = 1

type CPU struct {
	PC     uint16
	Memory [65536]byte
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	opcode := cpu.Memory[cpu.PC]
	cpu.PC++
	switch opcode {
	case OpNOP:
		// nothing to do
	default:
		panic(fmt.Sprintf("unimplemented opcode %d", opcode))
	}
}
