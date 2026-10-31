package r8_test

import (
	"testing"

	r8 "github.com/bitfield/go-machine"
)

func TestNewCPU_InitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	assert_eq_hex16(t, cpu.PC, 0x0000, "wrong PC")
	assert_eq_hex(t, cpu.A, 0x00, "wrong A")
	assert_zero(t, cpu, false)
	assert_eq_hex(t, cpu.Mem[0], 0x00, "wrong memory contents")
}

func TestStepIncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Step()
	assert_eq_hex16(t, cpu.PC, 0x0001, "wrong PC")
	cpu.Step()
	assert_eq_hex16(t, cpu.PC, 0x0002, "wrong PC")
}

func TestIncIncrementsA(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Zero = true
	cpu.RunProgram([]byte{
		r8.OpINC_A,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x01, "wrong A")
	assert_zero(t, cpu, false)
}

func TestIncWrapsAFromFFTo00(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0xFF
	cpu.RunProgram([]byte{
		r8.OpINC_A,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x00, "wrong A")
	assert_zero(t, cpu, true)
}

func TestDecDecrementsA(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0x01
	cpu.RunProgram([]byte{
		r8.OpDEC_A,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x00, "wrong A")
	assert_zero(t, cpu, true)
}

func TestDecWrapsAFrom00ToFF(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0x00
	cpu.Zero = true
	cpu.RunProgram([]byte{
		r8.OpDEC_A,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0xFF, "wrong A")
	assert_zero(t, cpu, false)
}

func TestStepWrapsPCFromFFFFTo0000(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.PC = 0xFFFF
	cpu.Step()
	assert_eq_hex16(t, cpu.PC, 0x0000, "wrong PC")
}

func TestMemoryIsBytes(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	var value byte = 0x00
	cpu.Mem[0x0000] = value
}

func TestRunRunsUntilHalted(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpINC_A,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x01, "wrong A")
	assert_eq_hex16(t, cpu.PC, 0x0002, "wrong PC")
}

func TestLdLoadsAccumulator(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0x0A
	cpu.Zero = true
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x05,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x05, "wrong A")
	assert_eq_hex16(t, cpu.PC, 0x0003, "wrong PC")
	assert_zero(t, cpu, false)
}

func TestLdSetsZeroFlag(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x00,
		r8.OpHALT,
	})
	assert_zero(t, cpu, true)
}

func TestAddAddsToAccumulator(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0x01
	cpu.Zero = true
	cpu.RunProgram([]byte{
		r8.OpADD_A, 0x02,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x03, "wrong A")
	assert_eq_hex16(t, cpu.PC, 0x0003, "wrong PC")
	assert_zero(t, cpu, false)
}

func TestAddSetsZeroFlag(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.A = 0x01
	cpu.RunProgram([]byte{
		r8.OpADD_A, 0xFF,
		r8.OpHALT,
	})
	assert_zero(t, cpu, true)
}

func TestJmpJumpsToSpecifiedAddress(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpJMP, 0x05, 0x00,
		r8.OpINC_A,
		r8.OpHALT,
		r8.OpHALT,
	})
	assert_eq_hex(t, cpu.A, 0x00, "wrong A")
	assert_eq_hex16(t, cpu.PC, 0x0006, "wrong PC")
}

func TestBraAlwaysBranches(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpBRA, 0x01,
		r8.OpHALT,
		r8.OpHALT,
	})
	assert_eq_hex16(t, cpu.PC, 0x0004, "wrong PC")
}

func TestBeqTakesBranchIfZeroFlagSet(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x00,
		r8.OpBEQ, 0x01,
		r8.OpHALT,
		r8.OpHALT,
	})
	assert_eq_hex16(t, cpu.PC, 0x0006, "wrong PC")
}

func TestBeqDoesNotBranchIfZeroFlagClear(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x01,
		r8.OpBEQ, 0x01,
		r8.OpHALT,
		r8.OpHALT,
	})
	assert_eq_hex16(t, cpu.PC, 0x0005, "wrong PC")
}

func TestCmpSetsZeroIfOperandEqual(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x01,
		r8.OpCMP_A, 0x01,
		r8.OpHALT,
	})
	assert_zero(t, cpu, true)
}

func TestCmpClearsZeroIfOperandNotEqual(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.RunProgram([]byte{
		r8.OpLD_A, 0x00,
		r8.OpCMP_A, 0x01,
		r8.OpHALT,
	})
	assert_zero(t, cpu, false)
}

func TestDisassembleCorrectlyDisassemblesCode(t *testing.T) {
	t.Parallel()
	cases := []struct {
		code   []byte
		source string
	}{
		{
			code:   []byte{r8.OpBEQ, 0x03},
			source: "beq 0x03",
		},
		{
			code:   []byte{r8.OpCMP_A, 0x01},
			source: "cmp a, 0x01",
		},
		{
			code:   []byte{r8.OpDEC_A},
			source: "dec a",
		},
		{
			code:   []byte{r8.OpHALT},
			source: "halt",
		},
		{
			code:   []byte{r8.OpINC_A},
			source: "inc a",
		},
		{
			code:   []byte{r8.OpJMP, 0x02, 0x00},
			source: "jmp 0x0002",
		},
		{
			code:   []byte{r8.OpLD_A, 0x0A},
			source: "ld a, 0x0A",
		},
		{
			code:   []byte{r8.OpTRAP, 0x20},
			source: "trap 0x20",
		},
		{
			code:   []byte{0xFF},
			source: "???",
		},
	}
	for _, tc := range cases {
		t.Run(tc.source, func(t *testing.T) {
			t.Parallel()
			source, bytes := r8.Disassemble(tc.code)
			if source != tc.source {
				t.Errorf("%v: want %q, got %q", tc.code, tc.source, source)
			}
			if bytes != len(tc.code) {
				t.Errorf("%s: want %d, got %d", tc.source, len(tc.code), bytes)
			}
		})
	}
}

func assert_eq_hex(t *testing.T, want, got byte, msg string) {
	t.Helper()
	if want != got {
		t.Errorf("%s: want 0x%02X, got 0x%02X", msg, want, got)
	}
}

func assert_eq_hex16(t *testing.T, want, got uint16, msg string) {
	t.Helper()
	if want != got {
		t.Errorf("%s: want 0x%04X, got 0x%04X", msg, want, got)
	}
}

func assert_zero(t *testing.T, cpu *r8.CPU, value bool) {
	if cpu.Zero != value {
		t.Errorf("want zero == %t, got %t", value, cpu.Zero)
	}
}
