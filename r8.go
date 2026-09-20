// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	A   byte
	PC  uint16
	Mem [65536]int
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	opcode := cpu.Mem[cpu.PC]
	cpu.PC++
	switch opcode {
	case 48: // `inc`
		cpu.A++
	case 64: // `dec`
		cpu.A--
	}
}
