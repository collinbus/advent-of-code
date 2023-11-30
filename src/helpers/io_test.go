package helpers

import "testing"

func TestReadInput(t *testing.T) {
	path := "test.txt"
	expectedLines := 2

	lines := ReadInput(path)

	if len(lines) != expectedLines {
		t.Fatalf("Expected %d lines but got %d\n", expectedLines, len(lines))
	}
	if lines[0] != "Hello world!" {
		t.Fatal("the first line is not correct")
	}
	if lines[1] != "This is a test file :)" {
		t.Fatal("the first line is not correct")
	}
}
