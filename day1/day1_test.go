package day1

import (
	"os"
	"testing"
)

func TestSampleSumCalibrationValuesPart1(t *testing.T) {
	input, _ := os.ReadFile("part1_sample_input.txt")

	expected := 142
	got := sumPart1(string(input))

	if got != expected {
		t.Fatalf("Got: %d Expected: %d", got, expected)
	}
}

func TestSampleSumCalibrationValuesPart2(t *testing.T) {
	input, _ := os.ReadFile("part2_sample_input.txt")

	expected := 281
	got := sumPart2(string(input))

	if got != expected {
		t.Fatalf("Got: %d Expected: %d", got, expected)
	}
}

func TestSumCalibrationValuesPart1(t *testing.T) {
	input, _ := os.ReadFile("input.txt")

	expected := 54338
	got := sumPart1(string(input))

	if got != expected {
		t.Fatalf("Got: %d Expected: %d", got, expected)
	}
}

func TestSumCalibrationValuesPart2(t *testing.T) {
	input, _ := os.ReadFile("input.txt")

	expected := 54338
	got := sumPart2(string(input))

	if got != expected {
		t.Fatalf("Got: %d Expected: %d", got, expected)
	}
}
