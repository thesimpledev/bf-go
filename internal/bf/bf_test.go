package bf

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	i := New(&bytes.Buffer{})
	tapeSizeWant := 30_000
	tapePosition := 0
	if i == nil {
		t.Fatal("client should not be null")
	}

	if i.tape == nil {
		t.Fatal("tape should not be null")
	}

	if len(i.tape) != tapeSizeWant {
		t.Errorf("tape is %d and should be %d", len(i.tape), tapeSizeWant)
	}

	if i.pos != tapePosition {
		t.Errorf("tape position got %d want %d", i.pos, tapePosition)
	}
}

func TestLoadInstructions(t *testing.T) {
	i := New(&bytes.Buffer{})
	want := ">>>>>"
	i.LoadInstructions(want)

	if want != i.instructions {
		t.Errorf("want %s, go %s", want, i.instructions)
	}
}

func TestExecuteInstructionsError(t *testing.T) {
	i := New(&bytes.Buffer{})
	i.LoadInstructions("")
	err := i.ExecutionLoop()
	if err == nil {
		t.Error("expected error for empty instructions, got nil")
	}
}

func TestExecuteInstructionsShiftRight(t *testing.T) {
	i := New(&bytes.Buffer{})
	inst := ">x>"
	want := 2
	i.LoadInstructions(inst)
	err := i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftRight(t *testing.T) {
	i := New(&bytes.Buffer{})
	want := 5
	for range want {
		i.shiftRight()
	}

	if i.pos != want {
		t.Errorf("got position %d want position %d", i.pos, want)
	}
}

func TestExecuteInstructionsShiftLeft(t *testing.T) {
	i := New(&bytes.Buffer{})
	inst := ">x><"
	want := 1
	i.LoadInstructions(inst)
	err := i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftLeft(t *testing.T) {
	i := New(&bytes.Buffer{})
	want := 3
	for range 5 {
		i.shiftRight()
	}

	for range 2 {
		i.shiftLeft()
	}

	if i.pos != want {
		t.Errorf("got position %d want position %d", i.pos, want)
	}
}

func TestShiftLeftNegative(t *testing.T) {
	i := New(&bytes.Buffer{})
	i.shiftLeft()
	want := 29_999

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftRightOverflow(t *testing.T) {
	i := New(&bytes.Buffer{})
	want := 0
	i.pos = 29_999
	i.shiftRight()

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestExecuteInstructionsIncrement(t *testing.T) {
	i := New(&bytes.Buffer{})
	inst := ">+"
	want := byte(1)
	i.LoadInstructions(inst)
	err := i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestIncrement(t *testing.T) {
	i := New(&bytes.Buffer{})
	i.increment()
	want := byte(1)
	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestExecuteInstructionsDecrement(t *testing.T) {
	i := New(&bytes.Buffer{})
	inst := ">-"
	want := byte(255)
	i.LoadInstructions(inst)
	err := i.ExecutionLoop()
	if err != nil {
		t.Errorf("expected no error got error %v", err)
	}

	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestDecrement(t *testing.T) {
	i := New(&bytes.Buffer{})
	i.decrement()
	want := byte(255)
	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	i := New(&buffer)
	i.decrement()
	want := string(byte(255))
	i.output()
	got := buffer.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExecuteInstructionOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	i := New(&buffer)
	i.LoadInstructions(">-.")
	want := string(byte(255))
	err := i.ExecutionLoop()
	if err != nil {
		t.Errorf("expected no error got error %v", err)
	}
	got := buffer.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestParserLoop(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		want         string
	}{
		{
			name:         "skip loop output B",
			instructions: "[impossible]+++++++++++[>++++++<-]>.",
			want:         "B",
		},
		{
			name:         "output A",
			instructions: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			want:         "Hello World!\n",
		},
		{
			name:         "output 5",
			instructions: "+++++[>++++++++++<-]>+++.",
			want:         "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			i := New(buffer)
			i.LoadInstructions(tt.instructions)
			err := i.ParserLoop()
			if err != nil {
				t.Error("expected no error, got error")
			}
			err = i.ExecutionLoop()
			if err != nil {
				t.Error("expected no error, got error")
			}
			if buffer.String() != tt.want {
				t.Errorf("got %q want %q", buffer.String(), tt.want)
			}
		})
	}
}
