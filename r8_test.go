package r8_test

import (
	"testing"

	r8 "github.com/bitfield/go-machine"
)

func TestNewInitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	if cpu.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", cpu.PC)
	}
	got := cpu.Mem[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

func TestStepTwiceIncrementsPCTwice(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	cpu.Mem[0] = 1
	cpu.Mem[1] = 1
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
	cpu := r8.New()
	cpu.Mem[0] = 1
	cpu.Mem[1] = 0
	cpu.Run()
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}
