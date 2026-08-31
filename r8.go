// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	PC  int
	Mem [65536]int
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	cpu.PC++
}
