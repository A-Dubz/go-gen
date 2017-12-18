package main

import (
	"bytes"
	"testing"
)

func TestMain(m *testing.M) {

}

func TestGenString(t *testing.T) {
	tests := []struct {
		buf    bytes.Buffer
		set    string
		length int
		want   int
	}{
		{bytes.Buffer{}, stringSet, 5, 5},
		{bytes.Buffer{}, stringSet, 10, 10},
		{bytes.Buffer{}, stringSet, 25, 25},
	}

	for _, tc := range tests {
		genString(&tc.buf, tc.set, tc.length)

		if got := len(tc.buf.String()); got != tc.want {
			t.Errorf("Expected %d, Got: %d", tc.length, got)
		}
	}
}

func BenchmarkGenerateStr(b *testing.B) {
	benchs := []struct {
		name   string
		buf    bytes.Buffer
		set    string
		length int
	}{
		{"Run at length = 5", bytes.Buffer{}, stringSet, 5},
		{"Run at length = 10", bytes.Buffer{}, stringSet, 10},
		{"Run at length = 25", bytes.Buffer{}, stringSet, 25},
		{"Run at length = 50", bytes.Buffer{}, stringSet, 50},
		{"Run at length = 100", bytes.Buffer{}, stringSet, 100},
	}

	for _, bm := range benchs {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				genString(&bm.buf, bm.set, bm.length)
				bm.buf.Reset()
			}
		})
	}
}
