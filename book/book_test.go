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

type mockWordFreqCalc struct {
	returnValue map[string]int
}

func (wfc mockWordFreqCalc) WordFrequency(lines []string) map[string]int {
	return wfc.returnValue
}

func TestBook_WordFrequency(t *testing.T) {
	expected := map[string]int{
		"foo": 42,
		"bar": 2,
		"baz": 999,
	}
	b := book.New(strings.NewReader(""), mockWordFreqCalc{expected})
	got := b.WordFrequency()
	if len(got) != len(expected) {
		t.Fatal("Word frequency map doesn't have expected number of entries")
	}
	for k := range expected {
		if expected[k] != got[k] {
			t.Errorf("Unexpected frequency count for key: %v.\nExpected: %v\nGot: %v", k, expected[k], got[k])
		}
	}

}
