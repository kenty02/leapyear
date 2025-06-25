package main

import (
	"testing"
)

func Test_leapyear(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		y    int
		want bool
	}{
		{"Year 2000", 2000, true},
		{"Year 1900", 1900, false},
		{"Year 2004", 2004, true},
		{"Year 2021", 2021, false},
		{"Year 1600", 1600, true},
		{"Year 1700", 1700, false},
		{"Year 1800", 1800, false},
		{"Year 2400", 2400, true},
		{"Year 2020", 2020, true},
		{"Year 2022", 2022, false},
		{"Year 2023", 2023, false},
		{"Year 2024", 2024, true},
		{"Year 2025", 2025, false},
		{"Year 2026", 2026, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leapyear(tt.y)
			if tt.want != got {
				t.Errorf("leapyear() = %v, want %v", got, tt.want)
			}
		})
	}
}
