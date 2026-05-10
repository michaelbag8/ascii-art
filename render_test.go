package main

import (
	"reflect"
	"testing"
)

func TestRenderLines(t *testing.T) {
	blockMaps := map[rune][]string{
		'A': {
			"A1", "A2", "A3", "A4",
			"A5", "A6", "A7", "A8",
		},
		'B': {
			"B1", "B2", "B3", "B4",
			"B5", "B6", "B7", "B8",
		},
	}

	tests := []struct {
		name     string
		segment  string
		expected []string
	}{
		{
			name:    "single character",
			segment: "A",
			expected: []string{
				"A1", "A2", "A3", "A4",
				"A5", "A6", "A7", "A8",
			},
		},
		{
			name:    "multiple characters",
			segment: "AB",
			expected: []string{
				"A1B1",
				"A2B2",
				"A3B3",
				"A4B4",
				"A5B5",
				"A6B6",
				"A7B7",
				"A8B8",
			},
		},
		{
			name:    "empty segment",
			segment: "",
			expected: []string{
				"", "", "", "", "", "", "", "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderLines(tt.segment, blockMaps)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
