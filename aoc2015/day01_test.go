package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{")", "-1", "1"},
		{"()())", "-1", "5"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day01, assert)
	}
}
