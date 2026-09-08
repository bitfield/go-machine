package r8_test

import (
	"testing"

	r8 "github.com/bitfield/go-machine"
)

func TestNewCPU_InitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	if cpu.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", cpu.PC)
	}
	if cpu.A != 0 {
		t.Errorf("after New, want A == 0, got %d", cpu.A)
	}
	got := cpu.Mem[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

func TestStepTwiceIncrementsPCTwice(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = r8.OpNOP
	cpu.Mem[1] = r8.OpNOP
	cpu.Step()
	if cpu.PC != 1 {
		t.Errorf("want pc == 1, got %d", cpu.PC)
	}
	cpu.Step()
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}

func TestRunRunsUntilHalted(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = r8.OpNOP
	cpu.Mem[1] = r8.OpHALT
	cpu.Run()
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}

func TestIncIncrementsA(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = r8.OpINC
	cpu.Mem[1] = r8.OpHALT
	cpu.Run()
	if cpu.A != 1 {
		t.Errorf("want a == 1, got %d", cpu.A)
	}
}
