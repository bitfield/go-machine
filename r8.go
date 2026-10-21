// Package r8 emulates a simple CPU called the R8.
package r8

import "encoding/binary"

const (
	OpADD_A = 80
	OpDEC_A = 64
	OpHALT  = 0
	OpINC_A = 48
	OpJMP   = 247
	OpLD_A  = 16
)

type CPU struct {
	A    byte
	PC   uint16
	Mem  [65536]byte
	Zero bool
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Fetch() byte {
	value := cpu.Mem[cpu.PC]
	cpu.PC++
	return value
}

func (cpu *CPU) RunProgram(program []byte) {
	copy(cpu.Mem[:], program)
	cpu.PC = 0
	cpu.Run()
}

func (cpu *CPU) Step() bool {
	opcode := cpu.Fetch()
	switch opcode {
	case OpADD_A:
		operand := cpu.Fetch()
		cpu.A += operand
		cpu.UpdateZero(cpu.A)
	case OpDEC_A:
		cpu.A--
		cpu.UpdateZero(cpu.A)
	case OpHALT:
		return false
	case OpINC_A:
		cpu.A++
		cpu.UpdateZero(cpu.A)
	case OpJMP:
		addr_lo := cpu.Fetch()
		addr_hi := cpu.Fetch()
		addr := binary.LittleEndian.Uint16([]byte{addr_lo, addr_hi})
		cpu.PC = addr
	case OpLD_A:
		operand := cpu.Fetch()
		cpu.A = operand
		cpu.UpdateZero(cpu.A)
	}
	return true
}

func (cpu *CPU) Run() {
	for cpu.Step() {
	}
}

func (cpu *CPU) UpdateZero(value byte) {
	cpu.Zero = value == 0
}
