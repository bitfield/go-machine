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
	cpu.Mem[0] = r8.OpINC_A
	cpu.Step()
	if cpu.A != 1 {
		t.Errorf("want A == 1, got %d", cpu.A)
	}
}

func TestIncWrapsAFrom255To0(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = r8.OpINC_A
	cpu.A = 255
	cpu.Step()
	if cpu.A != 0 {
		t.Errorf("want A == 0, got %d", cpu.A)
	}
}

func TestDecWrapsAFrom0To255(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Mem[0] = r8.OpDEC_A
	cpu.A = 0
	cpu.Step()
	if cpu.A != 255 {
		t.Errorf("want A == 255, got %d", cpu.A)
	}
}

func TestStepWrapsPCFrom65535To0(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.PC = 65535
	cpu.Step()
	if cpu.PC != 0 {
		t.Errorf("want PC == 0, got %d", cpu.PC)
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
	cpu.Mem[0] = r8.OpINC_A
	cpu.Mem[1] = r8.OpHALT
	cpu.Run()
	if cpu.A != 1 {
		t.Errorf("want A == 1, got %d", cpu.A)
	}
	if cpu.PC != 2 {
		t.Errorf("want PC == 2, got %d", cpu.PC)
	}
}

func TestLdLoadsAccumulator(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 10
	cpu.Mem[0] = r8.OpLD_A
	cpu.Mem[1] = 5
	cpu.Step()
	if cpu.A != 5 {
		t.Errorf("want A == 5, got %d", cpu.A)
	}
	if cpu.PC != 2 {
		t.Errorf("want PC == 2, got %d", cpu.PC)
	}
}
