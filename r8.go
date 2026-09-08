// Package r8 emulates a simple CPU called the R8.
package r8

import "fmt"

const (
	OpHALT = 0
	OpNOP  = 1
	OpINC  = 48
)

type CPU struct {
	A   int
	PC  int
	Mem [256]int
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}

func (cpu *CPU) Step() bool {
	opcode := cpu.Mem[cpu.PC]
	cpu.PC++
	switch opcode {
	case OpHALT:
		return false
	case OpNOP:
	// nothing to do
	case OpINC:
		cpu.A++
	default:
		panic(fmt.Sprintf("unimplemented opcode %d", opcode))
	}
	return true
}
