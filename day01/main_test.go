package main

import "testing"

func TestPart1(t *testing.T) {
	want := 69836
	value := Part1()

	if value != want {
		t.Errorf("TestPart1() = %v, want %v", value, want)
	}
}

func TestPart2(t *testing.T) {
	want := 207968
	value := Part2()

	if value != want {
		t.Errorf("TestPart2() = %v, want %v", value, want)
	}
}
