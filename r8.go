// Package r8 emulates a simple CPU called the R8.
package r8

const (
	OpDEC_A = 64
	OpHALT  = 0
	OpINC_A = 48
	OpLD_A  = 16
	OpADD_A = 80
)

type CPU struct {
	A   byte
	PC  uint16
	Mem [65536]byte
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Fetch() byte {
	value := cpu.Mem[cpu.PC]
	cpu.PC++
	return value
}

func (cpu *CPU) Step() bool {
	opcode := cpu.Fetch()
	switch opcode {
	case OpADD_A:
		operand := cpu.Fetch()
		cpu.A += operand
	case OpDEC_A:
		cpu.A--
	case OpHALT:
		return false
	case OpINC_A:
		cpu.A++
	case OpLD_A:
		operand := cpu.Fetch()
		cpu.A = operand
	}
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}
