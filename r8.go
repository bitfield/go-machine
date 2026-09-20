// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	A   byte
	PC  uint16
	Mem [65536]byte
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() bool {
	opcode := cpu.Mem[cpu.PC]
	cpu.PC++
	switch opcode {
	case 0: // `halt`
		return false
	case 48: // `inc`
		cpu.A++
	case 64: // `dec`
		cpu.A--
	}
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}
