package book_test

import (
	"challenge/book"
	"challenge/stats"
	"strings"
	"testing"
)

func TestBook_Read(t *testing.T) {
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
		b := book.New(strings.NewReader(item.have), stats.Stats{})
		if content := b.Read(); content != item.want {
			t.Errorf("Wanted: %v, got: %v", item.want, content)
		}
	}
}
func TestBook_ReadLines(t *testing.T) {
	t.Run("no_content", func(t *testing.T) {
		b := book.New(strings.NewReader(""), stats.Stats{})
		content, err := b.ReadLines(0, 20)
		if err == nil {
			t.Fatalf("Expected error when reading lines from empty book. Got: %v", content)
		}
	})
	t.Run("with content", func(t *testing.T) {
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
			b := book.New(strings.NewReader(item.have), stats.Stats{})
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

func TestBook_WordFrequency(t *testing.T) {
	data := []struct {
		have string
		want map[string]int
	}{
		{
			have: "",
			want: map[string]int{},
		},
		{
			have: "foo foo foo",
			want: map[string]int{"foo": 3},
		}, {
			have: "foo bar baz\nfoo baz",
			want: map[string]int{"foo": 2, "bar": 1, "baz": 2},
		},
	}
	for _, item := range data {
		b := book.New(strings.NewReader(item.have), stats.Stats{})
		got := b.WordFrequency()
		if len(got) != len(item.want) {
			t.Fatalf("Incorrect number of entries returned. Expected: %v, got: %v for input:\n%v", len(item.want), len(got), item.have)
		}
		for k := range item.want {
			if got[k] != item.want[k] {
				t.Errorf("Incorrect frequency. Expected: %v, got: %v\nFor input:\n%v", item.want[k], got[k], item.have)
			}
		}
	}
}
