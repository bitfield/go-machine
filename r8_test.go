package r8_test

import (
	"testing"

	r8 "github.com/bitfield/go-machine"
)

func TestNewCPU_InitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	if cpu.PC != 0 {
		t.Errorf("after New, want PC == 0, got %d", cpu.PC)
	}
	if cpu.A != 0 {
		t.Errorf("after New, want A == 0, got %d", cpu.A)
	}
	got := cpu.Mem[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

func TestStepIncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Step()
	if cpu.PC != 1 {
		t.Errorf("want PC == 1, got %d", cpu.PC)
	}
	cpu.Step()
	if cpu.PC != 2 {
		t.Errorf("want PC == 2, got %d", cpu.PC)
	}
}

func TestIncIncrementsA(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48 // `inc a`
	cpu.Step()
	if cpu.A != 1 {
		t.Errorf("want A == 1, got %d", cpu.A)
	}
}

func TestIncWrapsAFrom255To0(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48 // `inc a`
	cpu.A = 255
	cpu.Step()
	if cpu.A != 0 {
		t.Errorf("want A == 0, got %d", cpu.A)
	}
}
