package book_test

import (
	"challenge/book"
	"strings"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("Read", func(t *testing.T) {
		data := []struct {
			have string
			want string
		}{
			{
				have: "",
				want: "",
			},
			{
				have: "one line",
				want: "one line",
			}, {
				have: "line one\nline two",
				want: "line one\nline two",
			},
		}
		for _, item := range data {
			b := book.New(strings.NewReader(item.have))
			if content := b.Read(); content != item.want {
				t.Errorf("Wanted: %v, got: %v", item.want, content)
			}
		}
	})
	t.Run("ReadLines_no_content", func(t *testing.T) {
		b := book.New(strings.NewReader(""))
		content, err := b.ReadLines(0, 20)
		if err == nil {
			t.Fatalf("Expected error when reading lines from empty book. Got: %v", content)
		}
	})
	t.Run("ReadLines", func(t *testing.T) {
		data := []struct {
			have string
			want []string
		}{
			{
				have: "one line",
				want: []string{"one line"},
			}, {
				have: "line one\nline two",
				want: []string{"line one", "line two"},
			},
		}
		for _, item := range data {
			b := book.New(strings.NewReader(item.have))
			content, err := b.ReadLines(0, 20)
			if err != nil {
				t.Fatal(err)
			}
			if len(item.want) != len(content) {
				t.Fatalf("Incorrect number of lines returned. Expected: %v, got: %v for input:\n%v", len(item.want), len(content), item.have)
			}
			for i := range content {
				if item.want[i] != content[i] {
					t.Errorf("Failed to get expected line for input:\n%v\nWanted: %v, got: %v\n",
						item.want[i], content[i], item.have)
				}
			}
		}
	})
}
