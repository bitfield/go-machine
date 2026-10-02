// Package r8 emulates a simple CPU called the R8.
package r8

const (
	OpDEC  = 64
	OpHALT = 0
	OpINC  = 48
	OpLD   = 16
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
	case OpHALT:
		return false
	case OpINC:
		cpu.A++
	case OpDEC:
		cpu.A--
	case OpLD:
		operand := cpu.Mem[cpu.PC]
		cpu.PC++
		cpu.A += operand
	}
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}
