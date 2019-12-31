package aoc2015

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

// checkFirstBytesOfHash checks the first n nibbles (half-octets) of the MD5 of Augend+string(Addend)
// and inserts the augend to a chan int if the first n nibbles are all 0.
// It is assumed that the input is valid
// and that the channel is not blocked or closed.
func checkFirstBytesOfHash(n int, augend string, addend int, channel chan<- int) {
	builtString := []byte(augend + strconv.Itoa(addend))
	hash := fmt.Sprintf("%x", md5.Sum(builtString))
	for ii := 0; ii < n; ii++ {
		if hash[ii] != '0' {
			return // do nothing
		}
	}
	// once we know first n bits is zero...
	if len(channel) > 0 {
		return
	}
	channel <- addend
}

// Day04 solves the fourth day puzzle
// "The Ideal Stocking Stuffer"
func Day04(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// the entire scanner is read.
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		err = errors.New("first line of file is empty")
		return
	}

	// suppose we have an addends channel
	// which contains all the possible addends for input
	// i.e., input+<-addned
	processes := 16 // how many processes in parallel?
	addends := make(chan int, processes)
	go func() {
		moar := 1
		// add moar to addend channel
		for {
			addends <- moar
			moar++
		}
	}()
	_ = input

	for5Zeroes := make(chan int, 1) // will be blocked until we found answer for 5 zeroes
	for6Zeroes := make(chan int, 1) // will be blocked until we found answer for 6 zeroes
	foundFive, foundSix := false, false

	for ii := 0; ii < processes; ii++ {
		// let's create processes amount of goroutines
		go func() {
			for !foundFive || !foundSix {
				current := <-addends
				if !foundFive {
					checkFirstBytesOfHash(5, input, current, for5Zeroes)
				}
				if !foundSix {
					checkFirstBytesOfHash(6, input, current, for6Zeroes)
				}
			}
		}()
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		answer1 = strconv.Itoa(<-for5Zeroes)
		foundFive = true // since the statement above would empty for5Zeroes
		wg.Done()
	}()
	go func() {
		answer2 = strconv.Itoa(<-for6Zeroes)
		foundSix = true
		wg.Done()
	}()
	wg.Wait()
	return
}

// Day04ST solves the fourth day puzzle "The Ideal Stocking Stuffer"
// but is a single-threaded solution.
func Day04ST(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// the entire scanner is read.
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		err = errors.New("first line of file is empty")
		return
	}
	addend := 1        // to be appended to input
	foundFive := false // have we found five zeroes?
	foundSix := false  // have we found six zeroes?
	for {
		builtString := []byte(fmt.Sprintf("%v%v", input, addend))
		if !foundFive {
			firstFive := fmt.Sprintf("%x", md5.Sum(builtString))[:5]
			if firstFive == "00000" {
				foundFive = true
				answer1 = strconv.Itoa(addend)
			}
		}
		if !foundSix {
			firstSix := fmt.Sprintf("%x", md5.Sum(builtString))[:6]
			if firstSix == "000000" {
				foundSix = true
				answer2 = strconv.Itoa(addend)
			}

		}
		if foundFive && foundSix == true {
			break
		}
		addend++
	}
	return
}
