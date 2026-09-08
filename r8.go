// Package r8 emulates a simple CPU called the R8.
package r8

import "fmt"

const (
	OpHALT = 0
	OpNOP  = 1
)

type CPU struct {
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
	default:
		panic(fmt.Sprintf("unimplemented opcode %d", opcode))
	}
	return true
}
