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
	got := cpu.Memory[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

func TestStepIncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	cpu.Memory[256] = r8.OpNOP
	cpu.Memory[257] = r8.OpNOP
	cpu.PC = 256
	cpu.Step()
	if cpu.PC != 257 {
		t.Errorf("want pc == 257, got %d", cpu.PC)
	}
	cpu.Step()
	if cpu.PC != 258 {
		t.Errorf("want pc == 258, got %d", cpu.PC)
	}
}
