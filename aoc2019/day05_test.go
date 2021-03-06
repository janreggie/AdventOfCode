package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2019D05 my input",
			Input:   day05myInput,
			Result1: "9938601",
			Result2: "4283952"},
	}
	for _, tt := range testCases {
		tt.Test(Day05, assert)
	}
}

func BenchmarkDay05(b *testing.B) {
	aoc.Benchmark(Day05, b, day05myInput)
}
