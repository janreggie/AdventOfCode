package aoc2015

import (
	"strconv"
	"sync"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// eggnog represents an amount of eggnog.
type eggnog uint

// containerCount represents a number of containers
type containerCount uint

// contain tries to contain some amount of eggnog into containers
// that can hold a variable amount of eggnog
// and returns the number of possibilities.
func contain(amount eggnog, containers []eggnog) containerCount {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1 // there is only one arrangement: store nothing into no containers!
	}
	if len(containers) == 0 {
		return 0 // there are no containers to store them to
	}
	return contain(amount, containers[1:]) + contain(amount-containers[0], containers[1:])
}

// containOnlyIn contains an amount of eggnog using only a set limit of containers
// and returns the number of possible ways this can be done.
func containOnlyIn(amount eggnog, limit containerCount, containers []eggnog) containerCount {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1 // there is only one arrangement: store nothing into no containers!
	}
	if len(containers) == 0 || limit == 0 {
		return 0 // there are no containers to store them to
	}
	return containOnlyIn(amount-containers[0], limit-1, containers[1:]) +
		containOnlyIn(amount, limit, containers[1:])
}

// Day17 solves the seventeenth day puzzle "No Such Thing as Too Much".
//
// Input
//
// A file containing 20 lines each of which describes
// a container and how much it contains.
// For example:
//
//	50
//	44
//	11
//	49
//	42
//
// There are no guarantees that the list in the input containers have been
// sorted. All containers should be non-negative integers no more than 150.
func Day17(input string) (answer1, answer2 string, err error) {

	containers := make([]eggnog, 0)
	for _, line := range aoc.SplitLines(input) {
		size, e := strconv.ParseUint(line, 0, 32)
		if e != nil {
			err = errors.Wrapf(e, "could not parse %q", line)
			return
		}
		containers = append(containers, eggnog(size))
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		answer1 = strconv.FormatUint(uint64(contain(150, containers)), 10)
		wg.Done()
	}()

	// now determine the least amount...
	// This is a lazy solution and just iterates ii until we found an answer.
	go func() {
		for ii := containerCount(0); ii < containerCount(len(containers)); ii++ {
			if counted := containOnlyIn(150, ii, containers); counted > 0 {
				answer2 = strconv.FormatUint(uint64(counted), 10)
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return
}
