package main

import "testing"

func TestPart1(t *testing.T) {
	want := 13565
	value := Part1()

	if value != want {
		t.Errorf("TestPart1() = %v, want %v", value, want)
	}
}

func TestPart2(t *testing.T) {
	want := 12424
	value := Part2()

	if value != want {
		t.Errorf("TestPart2() = %v, want %v", value, want)
	}
}
