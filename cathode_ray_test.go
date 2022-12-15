package advent_test

import (
	a "scripts/advent"
	"testing"
)

func TestSignalStrength(t *testing.T) {
	t.Run("run short input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/cathode_ray/short_test_input.txt")
		cpu, err := a.SignalStrength(&data)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("register value after CPU cycle %d is %d\n", len(cpu.Cycles)-1, cpu.RegisterValue)
		signals, sum := cpu.ReadSignalStrength([]int{1, 2, 3, 4, 5})
		t.Logf("signal strengths are %v with the sum of %d\n", signals, sum)
	})

	t.Run("run opTest input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/cathode_ray/test_input.txt")
		cpu, err := a.SignalStrength(&data)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("register value after CPU cycle %d is %d\n", len(cpu.Cycles)-1, cpu.RegisterValue)
		signals, sum := cpu.ReadSignalStrength([]int{20, 60, 100, 140, 180, 220})
		t.Logf("signal strengths are %v with the sum of %d\n", signals, sum)
	})

	t.Run("run long input input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/cathode_ray/input.txt")
		cpu, err := a.SignalStrength(&data)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("register value after CPU cycle %d is %d\n", len(cpu.Cycles)-1, cpu.RegisterValue)
		signals, sum := cpu.ReadSignalStrength([]int{20, 60, 100, 140, 180, 220})
		t.Logf("signal strengths are %v with the sum of %d\n", signals, sum)
		cpu.RenderCRT()
	})
}
