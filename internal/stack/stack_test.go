package stack

import "testing"

func TestNew(t *testing.T) {
	s := New()

	if s == nil {
		t.Fatal("stack was nil and should not be")
	}
}

func TestPush(t *testing.T) {
	s := New()
	want := 1
	s.Push(want)
	got, _ := s.Pop()

	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPopEmpty(t *testing.T) {
	s := New()
	_, exists := s.Pop()
	if exists {
		t.Error("Exists is True and should not be")
	}
}

func TestPop(t *testing.T) {
	s := New()
	want := 1
	s.Push(5)
	s.Push(want)
	num, exists := s.Pop()
	if !exists {
		t.Errorf("doesn't exist and should")
	}

	if num != want {
		t.Errorf("got %d want %d", num, want)
	}

	if len(s.items) != 1 {
		t.Errorf("pop failed to remove item")
	}

	if s.items[0] != 5 {
		t.Errorf("pop removed the wrong item")
	}
}
