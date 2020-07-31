package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

const day17myInput = `50
44
11
49
42
46
18
32
26
40
21
7
18
43
10
47
36
24
22
40
`

func TestDay17(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
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
	for ii := 0; ii < b.N; ii++ {
		Day17(bufio.NewScanner(strings.NewReader(day17myInput)))
	}
}