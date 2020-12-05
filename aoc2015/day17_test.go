package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2019D17 sample input",
			Input:   day17myInput,
			Result1: "654",
			Result2: "57"},
	}
	for _, tt := range testCases {
		tt.Test(Day17, assert)
	}
}

func BenchmarkDay17(b *testing.B) {
	internal.Benchmark(Day17, b, day17myInput)
}
