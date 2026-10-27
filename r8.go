// Package r8 emulates a simple CPU called the R8.
package r8

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

const (
	OpADD_A = 80
	OpBEQ   = 241
	OpBRA   = 240
	OpDEC_A = 64
	OpHALT  = 0
	OpINC_A = 48
	OpJMP   = 247
	OpLD_A  = 16
	OpTRAP  = 249
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
	case OpBEQ:
		displacement := cpu.Fetch()
		if cpu.Zero {
			cpu.PC += uint16(displacement)
		}
	case OpBRA:
		displacement := cpu.Fetch()
		cpu.PC += uint16(displacement)
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
	case OpTRAP:
		operand := cpu.Fetch()
		if operand == 32 {
			fmt.Printf("%c", cpu.A)
		}
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

func Disassemble(code []byte) (string, int) {
	switch code[0] {
	case OpBEQ:
		return fmt.Sprintf("beq %s", format_byte(code[1:])), 2
	case OpDEC_A:
		return "dec a", 1
	case OpHALT:
		return "halt", 1
	case OpINC_A:
		return "inc a", 1
	case OpJMP:
		return fmt.Sprintf("jmp %s", format_word(code[1:])), 3
	case OpLD_A:
		return fmt.Sprintf("ld a, %s", format_byte(code[1:])), 2
	case OpTRAP:
		return fmt.Sprintf("trap %s", format_byte(code[1:])), 2
	default:
		return "???", 1
	}
}

func format_byte(code []byte) string {
	if len(code) < 1 {
		return "???"
	}
	return strconv.Itoa(int(code[0]))
}

func format_word(code []byte) string {
	if len(code) < 2 {
		return "???"
	}
	addr_lo, addr_hi := code[0], code[1]
	addr := binary.LittleEndian.Uint16([]byte{addr_lo, addr_hi})
	return strconv.Itoa(int(addr))
}
