package advent_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	"testing"
)

func TestMonkeyBusiness(t *testing.T) {
	t.Run("process short test input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/monkey/short_input.txt")
		monkeys, f, err := a.MonkeyBusiness(&data)
		if err != nil {
			t.Errorf("encountered an error while creating monkeys: %v", err)
		}
		for i := 0; i < 10000; i++ {
			err = monkeys.Round(f)
			if err != nil {
				t.Errorf("encountered an error while running round %d: %v", err, i)
			}
		}
		for id, mk := range *monkeys {
			t.Logf("monkey %d has total count of %d\n", id, mk.Count)
		}
		t.Log("=======================")
		t.Logf("the level of busy is %d", monkeys.Level())
	})

	t.Run("process default test input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/monkey/input.txt")
		monkeys, factor, err := a.MonkeyBusiness(&data)
		t.Logf("factor is %d\n", factor)
		if err != nil {
			t.Errorf("encountered an error while creating monkeys: %v", err)
		}
		for i := 0; i < 10000; i++ {
			err = monkeys.Round(factor)
			if err != nil {
				t.Errorf("encountered an error while running round %d: %v", err, i)
			}
		}
		for id, mk := range *monkeys {
			t.Logf("monkey %d has total count of %d\n", id, mk.Count)
		}
		t.Log("=======================")
		t.Logf("the level of busy in long input is %d", monkeys.Level())
	})
}
