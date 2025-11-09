// Package bf if brain fuckery
package bf

import (
	"fmt"
	"io"
)

const (
	tapeSize = 30_000
)

type Interpreter struct {
	tape         []byte
	pos          int
	instructions string
	w            io.Writer
}

func New(w io.Writer) *Interpreter {
	client := &Interpreter{
		tape: make([]byte, tapeSize),
		w:    w,
	}
	return client
}

func (i *Interpreter) LoadInstructions(instructions string) {
	i.instructions = instructions
}

func (i *Interpreter) ExecutionLoop() error {
	if len(i.instructions) == 0 {
		return fmt.Errorf("instructions are empty")
	}
	for _, instruction := range i.instructions {
		switch instruction {
		case '>':
			i.shiftRight()
		case '<':
			i.shiftLeft()
		case '+':
			i.increment()
		case '-':
			i.decrement()
		case '.':
			i.output()
		default:
			continue
		}
	}

	return nil
}

func (i *Interpreter) shiftRight() {
	if i.pos == tapeSize-1 {
		i.pos = 0
		return
	}
	i.pos++
}

func (i *Interpreter) shiftLeft() {
	if i.pos == 0 {
		i.pos = tapeSize - 1
		return
	}
	i.pos--
}

func (i *Interpreter) increment() {
	i.tape[i.pos]++
}

func (i *Interpreter) decrement() {
	i.tape[i.pos]--
}

func (i *Interpreter) output() {
	_, _ = fmt.Fprintf(i.w, "%c", i.tape[i.pos])
}
