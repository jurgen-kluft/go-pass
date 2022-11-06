package clipboard_test

import (
	"testing"

	. "github.com/jurgen-kluft/go-pass/clipboard"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "Dutch (Nederlands)"

	err := Write(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := Read()
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	expected1 := "French: éèêëàùœç"
	expected2 := "Specific UTF-8 characters: 𐍈𐌹𐌻𐌰𐌿𐌰"

	err := Write(expected1)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	if actual1 != expected1 {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = Write(expected2)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	if actual2 != expected2 {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	text := "Chinese (中文): 你好世界"
	for i := 0; i < b.N; i++ {
		Write(text)
	}
}
