// Package r8 emulates a simple CPU called the R8.
package r8

import (
	"encoding/binary"
	"fmt"
)

const (
	OpADD_A = 0x50
	OpADD_B = 0x51
	OpBEQ   = 0xF1
	OpBRA   = 0xF0
	OpCMP_A = 0x70
	OpCMP_B = 0x71
	OpDEC_A = 0x40
	OpDEC_B = 0x41
	OpHALT  = 0x00
	OpINC_A = 0x30
	OpINC_B = 0x31
	OpJMP   = 0xF7
	OpLD_A  = 0x10
	OpLD_B  = 0x11
	OpTRAP  = 0xF9

	RegA = 0x00
	RegB = 0x01

	TrapPRINT = 0x20
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
	cpu.PC = 0x0000
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
	case OpCMP_A:
		operand := cpu.Fetch()
		cpu.Zero = cpu.A == operand
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
		if operand == TrapPRINT {
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
	cpu.Zero = value == 0x00
}

type Reg byte

func Disassemble(code []byte) (string, int) {
	switch code[0] {
	case OpBEQ:
		return fmt.Sprintf("beq %s", format_byte(code[1:])), 2
	case OpCMP_A:
		return fmt.Sprintf("cmp a, %s", format_byte(code[1:])), 2
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
	return fmt.Sprintf("0x%02X", code[0])
}

func format_word(code []byte) string {
	if len(code) < 2 {
		return "???"
	}
	addr_lo, addr_hi := code[0], code[1]
	addr := binary.LittleEndian.Uint16([]byte{addr_lo, addr_hi})
	return fmt.Sprintf("0x%04X", addr)
}

func GetTarget(opcode byte) Reg {
	return Reg(opcode & 0x0F)
}
