package main

import "testing"

func TestFindHouse_SmallTargets(t *testing.T) {
	tests := []struct {
		name       string
		target     int
		multiplier int
		maxVisits  int
		workers    int
		expected   int
	}{
		{
			name:       "Part1_SmallTarget_70",
			target:     70,
			multiplier: 10,
			maxVisits:  0,
			workers:    4,
			expected:   4, // House 4 gets 10+20+40 = 70
		},
		{
			name:       "Part1_SmallTarget_120",
			target:     120,
			multiplier: 10,
			maxVisits:  0,
			workers:    4,
			expected:   6, // House 6 gets 10+20+30+60 = 120
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findHouse(tt.target, tt.multiplier, tt.maxVisits, tt.workers)
			if got != tt.expected {
				t.Errorf("expected house %d, got %d", tt.expected, got)
			}
		})
	}
}
