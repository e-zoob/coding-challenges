package lib

import "testing"

func TestCountLines(t *testing.T) {
	exampleString := "Line one\nLine two\nLine three\n"
	result := CountLines(exampleString)
	if result != 3 {
		t.Errorf("CountLines(%q) returned %d, expected 2", exampleString, result)
	}
}
func TestCountWords(t *testing.T) {
	exampleString := "Hello, World!\nThis is a test.\n"
	result := CountWords(exampleString)
	if result != 6 {
		t.Errorf("CountWords(%q) returned %d, expected 6", exampleString, result)
	}
}

func TestCountBytes(t *testing.T) {
	exampleString := "this is an example"
	result := CountBytes(exampleString)
	if result != 18 {
		t.Errorf("CountBytes(%q) returned %d, expected 18", exampleString, result)
	}
}

func TestCountChars(t *testing.T) {
	exampleString := "Hello, World! This is a test.\n"
	result := CountChars(exampleString)
	if result != 30 {
		t.Errorf("CountChars(%q) returned %d, expected 30", exampleString, result)
	}
}
