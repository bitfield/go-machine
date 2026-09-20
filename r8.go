// Package r8 emulates a simple CPU called the R8.
package r8

const (
	OpDEC_A = 64
	OpHALT  = 0
	OpINC_A = 48
)

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
	case OpDEC_A:
		cpu.A--
	case OpHALT:
		return false
	case OpINC_A:
		cpu.A++
	}
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}
