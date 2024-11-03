package main

import "testing"

func TestSearchString(t *testing.T) {
	tests := []struct {
		s, text string
		want    int
	}{
		{"cat", "dogmice cat horse goat animal", 8},
		{"anteater", "dogmice horse goat animal", -1},
		{"elephant", "cat", -1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := Search(tt.s, tt.text)
			if got != tt.want {
				t.Errorf("for %s -- got %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
