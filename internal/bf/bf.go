// Package bf if brain fuckery
package bf

import (
	"fmt"
	"io"

	"brainfuck/internal/stack"
)

const (
	tapeSize = 30_000
)

type Interpreter struct {
	tape         []byte
	pos          int
	instructions string
	w            io.Writer
	loops        map[int]int
	stack        *stack.Stack
}

func New(w io.Writer) *Interpreter {
	client := &Interpreter{
		tape:  make([]byte, tapeSize),
		w:     w,
		loops: make(map[int]int),
		stack: stack.New(),
	}
	return client
}

func (i *Interpreter) LoadInstructions(instructions string) {
	i.instructions = instructions
}

func (i *Interpreter) ParserLoop() error {
	if len(i.instructions) == 0 {
		return fmt.Errorf("instructions are empty")
	}

	for num, instruction := range i.instructions {
		switch instruction {
		case '[':
			i.stack.Push(num)
		case ']':
			start, exists := i.stack.Pop()
			if !exists {
				continue
			}
			i.loops[start] = num
			i.loops[num] = start
		default:
			continue
		}
	}
	return nil
}

func (i *Interpreter) ExecutionLoop() error {
	if len(i.instructions) == 0 {
		return fmt.Errorf("instructions are empty")
	}
	for inst := 0; inst < len(i.instructions); inst++ {
		switch i.instructions[inst] {
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
		case '[':
			inst = i.startLoop(inst)
		case ']':
			inst = i.endLoop(inst)
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

func (i *Interpreter) startLoop(inst int) int {
	if i.tape[i.pos] == 0 {
		return i.loops[inst]
	}

	return inst
}

func (i *Interpreter) endLoop(inst int) int {
	if i.tape[i.pos] != 0 {
		return i.loops[inst] - 1
	}

	return inst
}
