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

func TestStepIncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Step()
	if cpu.PC != 1 {
		t.Errorf("want pc == 1, got %d", cpu.PC)
	}
	cpu.Step()
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}

func TestIncIncrementsA(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48 // `inc`
	cpu.Step()
	if cpu.A != 1 {
		t.Errorf("want a == 1, got %d", cpu.A)
	}
}

func TestIncWrapsAFrom255To0(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48 // `inc`
	cpu.A = 255
	cpu.Step()
	if cpu.A != 0 {
		t.Errorf("want a == 0, got %d", cpu.A)
	}
}

func TestDecWrapsAFrom0To255(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 64 // `dec`
	cpu.A = 0
	cpu.Step()
	if cpu.A != 255 {
		t.Errorf("want a == 255, got %d", cpu.A)
	}
}

func TestStepWrapsPCFrom65535To0(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.PC = 65535
	cpu.Step()
	if cpu.PC != 0 {
		t.Errorf("want pc == 0, got %d", cpu.PC)
	}
}

func TestMemoryIsBytes(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	var value byte = 0
	cpu.Mem[0] = value
}

func TestRunRunsUntilHalted(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48 // `inc`
	cpu.Mem[1] = 0  // `halt`
	cpu.Run()
	if cpu.A != 1 {
		t.Errorf("want a == 1, got %d", cpu.A)
	}
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}
