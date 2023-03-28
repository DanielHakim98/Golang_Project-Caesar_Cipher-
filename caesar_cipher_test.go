package main

import "testing"

func TestCaesarInputA(t *testing.T) {
	var word = "A"
	var move int64 = 1
	output, err := caesar(word, move)

	expected := "B"
	if output != expected || err != nil {
		t.Fatalf("Expected '%v', got '%v'\n", expected, output)
	}
}

func TestCaesarInputHelloWorld(t *testing.T) {
	var word = "Hello, World!"
	var move int64 = 5
	output, err := caesar(word, move)

	expected := "Mjqqt, Btwqi!"
	if output != expected || err != nil {
		t.Fatalf("Expected '%v', got '%v'\n", expected, output)
	}
}

func TestCaesarZAnd1(t *testing.T) {
	var word = "Z"
	var move int64 = 1
	output, err := caesar(word, move)

	expected := "A"
	if output != expected || err != nil {
		t.Fatalf("Expected '%v', got '%v'\n", expected, output)
	}
}
func TestCaesarEdgeCasePositive75(t *testing.T) {
	var word = "Hello, World!"
	var move int64 = 75

	expected := "Ebiil, Tloia!"
	output, err := caesar(word, move)
	if output != expected || err != nil {
		t.Fatalf("Expected '%v', got '%v'\n", expected, output)
	}
}
func TestCaesarEdgeCaseNegative29(t *testing.T) {
	var word = "Hello, World!"
	var move int64 = -29

	expected := "Ebiil, Tloia!"
	output, err := caesar(word, move)
	if output != expected || err != nil {
		t.Fatalf("Expected '%v', got '%v'\n", expected, output)
	}
}
