package day2

import (
	"os"
	"testing"
)

func TestSampleSumOfPossibleGameIds(t *testing.T) {
	input, _ := os.ReadFile("part1_sample_input.txt")

	expected := 8
	cubes := GameCubes{red: 12, green: 13, blue: 14}
	got := SumOfPossibleGameIds(string(input), cubes)

	if expected != got {
		t.Fatalf("Expected: %d Got: %d", expected, got)
	}
}

func TestSumOfPossibleGameIds(t *testing.T) {
	input, _ := os.ReadFile("input.txt")

	expected := 1853
	cubes := GameCubes{red: 12, green: 13, blue: 14}
	got := SumOfPossibleGameIds(string(input), cubes)

	if expected != got {
		t.Fatalf("Expected: %d Got: %d", expected, got)
	}
}
