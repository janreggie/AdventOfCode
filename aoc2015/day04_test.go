package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2015D04 my input",
			Input:   day04myInput,
			Result1: "346386",
			Result2: "9958218"},
	}
	for _, tt := range testCases {
		tt.Test(Day04, assert)
	}
}

func BenchmarkDay04(b *testing.B) {
	internal.Benchmark(Day04, b, day04myInput)
}

func BenchmarkDay04ST(b *testing.B) {
	internal.Benchmark(Day04ST, b, day04myInput)
}
