// Package intcode implements programming in "Intcode"
// for Advent of Code 2019.
package intcode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Halt is a module that is built in to the Intcode
var Halt *Module = NewModule(ModuleConfig{
	Opcode:        99,
	Mnemonic:      "HALT",
	Parameterized: false,
	Function: func(ic *Intcode) error {
		return NewHaltError("HALT (99)") // that's its literally function
	},
})

// Intcode implements an "Intcode" computer
// consisting of a program counter and a tape of memory
// as well as a list of "modules"
// which are functions that take in an Intcode
// and may return an error.
type Intcode struct {
	pc           int64     // program counter
	mem          []int64   // memory
	modules      []*Module // all modules it has
	input        []int64   // an input "stack" (FILO)
	output       []int64   // an output "stack"
	relativeBase int64     // the "relative base" of the computer
}

// New generates an Intcode using a memory reel
func New(mem []int64) *Intcode {
	ic := &Intcode{pc: 0, mem: make([]int64, len(mem))}
	copy(ic.mem, mem)
	ic.Install(Halt)
	return ic
}

// NewFromString generates using a memory reel represented by a string
func NewFromString(input string) (*Intcode, error) {

	mem := make([]int64, 0)
	for _, cell := range strings.Split(input, ",") {
		cell = strings.TrimSpace(cell)
		memVal, err := strconv.ParseInt(cell, 10, 64)
		if err != nil {
			return nil, errors.Errorf("could not parse memory cell %v in input", cell)
		}
		mem = append(mem, memVal)
	}
	return New(mem), nil
}

// Details returns intimate details of the Intcode state
func (ic *Intcode) Details() string {

	var sb strings.Builder
	fmt.Fprintf(&sb, "State: PC: %v (%v)\n", ic.PC(), ic.Current())
	fmt.Fprintf(&sb, "Memory: %v\n", ic.Snapshot())
	fmt.Fprintf(&sb, "Input: %v\n", ic.Input())
	fmt.Fprintf(&sb, "Output: %v", ic.Output())
	return sb.String()
}

// Install installs a module
func (ic *Intcode) Install(module *Module) {
	ic.modules = append(ic.modules, module)
}

// UninstallAll removes all modules
func (ic *Intcode) UninstallAll() {
	ic.modules = make([]*Module, 0)
}

// Operate performs instructions on the Intcode computer
// depending on the modules it has
func (ic *Intcode) Operate() error {
	// run through the modules
	for ic.pc < int64(len(ic.mem)) { // while we haven't reached the end yet
		// let's go through all the modules
		isFound := false
		for _, module := range ic.modules {
			opcode := ic.Current()
			if module.parameterized {
				opcode = opcode % 100 // we only care about the last two
			}
			if module.opcode == opcode {
				if err := module.function(ic); err != nil {
					if !IsHalt(err) {
						return NewOperationError(err, ic)
					}
					return err
				}
				isFound = true
				break // out of the loop once we found a module
			}
		}
		if !isFound {
			return NewInvalidOpcodeError(ic.Current(), ic.pc)
		}
	}
	return nil
}

// Snapshot returns a copy of its memory
func (ic *Intcode) Snapshot() (mem []int64) {
	mem = make([]int64, ic.Len())
	for i := range ic.mem {
		mem[i] = ic.mem[i]
	}
	return
}

// allocateMore allocates number amount of memory locations for ic.mem
func (ic *Intcode) allocateMore(amount int64) {
	// make sure amount > 0
	if amount < 0 {
		return
	}
	ic.mem = append(ic.mem, make([]int64, amount)...)
}

// GetLocation returns the value of the memory at a particular location.
// If location is more than the memory length,
// ic.mem is reallocated.
// If location is negative it will simply return an error.
func (ic *Intcode) GetLocation(location int64) (value int64, err error) {
	if location < 0 {
		err = NewOutOfBoundsError(location, ic.Len())
	}
	if location >= ic.Len() {
		ic.allocateMore(location - ic.Len() + 1)
	}
	value = ic.mem[location]
	return
}

// GetNext returns a fragment of memory after Current()
// containing the next count locations
func (ic *Intcode) GetNext(count int64) (mem []int64, err error) {
	mem = make([]int64, count)
	for ii := int64(0); ii < count; ii++ {
		if mem[ii], err = ic.GetLocation(ic.pc + ii + 1); err != nil {
			return
		}
	}
	return
}

// SetLocation sets the value of the memory at some location.
// If location is more than the memory length,
// ic.mem is reallocated.
// If location is negative it will simply return an error.
func (ic *Intcode) SetLocation(location, value int64) (err error) {
	if location < 0 {
		err = NewOutOfBoundsError(location, ic.Len())
		return
	}
	if location >= ic.Len() {
		ic.allocateMore(location - ic.Len() + 1)
	}
	ic.mem[location] = value
	return
}

// PC returns the current value for the program counter
func (ic *Intcode) PC() (pc int64) {
	pc = ic.pc
	return
}

// Current returns the current memory location at the program counter
func (ic *Intcode) Current() (value int64) {
	return ic.mem[ic.pc]
}

// Len returns the length of the memory
func (ic *Intcode) Len() (length int64) {
	return int64(len(ic.mem))
}

// Increment increments the program counter by a set amount
func (ic *Intcode) Increment(value int64) (err error) {
	if (value+ic.pc) > ic.Len() || (value+ic.pc) < 0 {
		// greater than since we can assume pc equals Len()
		return NewOutOfBoundsError(value+ic.pc, ic.Len())
	}
	ic.pc += value
	return
}

// Jump jumps the program counter to some value
func (ic *Intcode) Jump(value int64) (err error) {
	if value > ic.Len() || value < 0 {
		// greater than since we can assume pc equals Len() (where it will just halt)
		return NewOutOfBoundsError(value, ic.Len())
	}
	ic.pc = value
	return
}

// RelativeBase returns the relative base of the ic computer
func (ic *Intcode) RelativeBase() (relativeBase int64) {
	relativeBase = ic.relativeBase
	return
}

// AdjustRelativeBase adjusts the relative base by some amount,
// increasing or decreasing it.
func (ic *Intcode) AdjustRelativeBase(amount int64) {
	ic.relativeBase += amount
}

// SetRelativeBase sets the relative base by some amount.
func (ic *Intcode) SetRelativeBase(amount int64) {
	ic.relativeBase = amount
}

// Rewind jumps PC to zero
func (ic *Intcode) Rewind() {
	ic.Jump(0) // this cannot return any error
}

// Format formats the memory, input, and outputs and sets PC and relative base to zero
// but does not remove installed modules
func (ic *Intcode) Format(mem []int64) {
	ic.mem = make([]int64, len(mem))
	copy(ic.mem, mem)
	ic.SetInput()
	ic.ResetOutput()
	ic.SetRelativeBase(0)
	ic.Rewind()
	return
}

// Input returns a copy of its inputs
func (ic *Intcode) Input() (input []int64) {
	input = make([]int64, len(ic.input))
	for ii := range input {
		input[ii] = ic.input[ii]
	}
	return
}

// PushToInput pushes a value to the input queue
func (ic *Intcode) PushToInput(input int64) {
	ic.input = append(ic.input, input)
	return
}

// SetInput sets the input
func (ic *Intcode) SetInput(inputs ...int64) {
	ic.input = make([]int64, len(inputs))
	for ii := range inputs {
		ic.input[ii] = inputs[ii]
	}
}

// GetInput removes an input from the queue
func (ic *Intcode) GetInput() (input int64, err error) {
	if len(ic.input) == 0 {
		err = fmt.Errorf("input is length zero")
		return
	}
	input = ic.input[len(ic.input)-1]
	ic.input = ic.input[:len(ic.input)-1]
	return
}

// ResetOutput resets the outputs
func (ic *Intcode) ResetOutput() {
	ic.output = make([]int64, 0)
}

// PushToOutput pushes a value to its outputs
func (ic *Intcode) PushToOutput(value int64) {
	ic.output = append(ic.output, value)
}

// Output prints the output
func (ic *Intcode) Output() (output []int64) {
	output = make([]int64, len(ic.output))
	for ii := range output {
		output[ii] = ic.output[ii]
	}
	return
}

// GetOutput removes an output from the stack
func (ic *Intcode) GetOutput() (output int64, err error) {
	if len(ic.output) == 0 {
		err = fmt.Errorf("output is length zero")
		return
	}
	output = ic.output[len(ic.output)-1]
	ic.output = ic.output[:len(ic.output)-1]
	return
}
