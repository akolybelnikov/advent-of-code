package advent_test

import (
	a "scripts/advent"
	"testing"
)

func TestTraverseDirs(t *testing.T) {
	t.Run("traverse default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/system_update/input.txt")
		root, err := a.TraverseDirs(&data)
		if err != nil {
			t.Fatalf("returned an error: %v", err)
		}

		t.Logf("root dir size: %d\n", root.GetSize())

		free := 70000000 - root.GetSize()
		t.Logf("available free %d", free)

		need := 30000000 - free
		t.Logf("need to free %d", need)

		dirs := a.FindSystemDirs(root)
		t.Logf("total dirs in system: %d\n", len(*dirs))

		filterDirs := a.FilterDirsBySize(dirs, need)
		t.Logf("top candidate dir has %d size bytes\n", (*filterDirs)[0].GetSize())

		total := 0
		for _, d := range *dirs {
			if d.GetSize() <= 100000 {
				total += d.GetSize()
			}
		}
		t.Logf("total under 100000 bytes: %d\n", total)
	})
}
