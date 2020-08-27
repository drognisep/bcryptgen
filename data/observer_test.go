package data

import (
	"testing"
)

func TestSubjectStateUpdates(t *testing.T) {
	subject := NewStringSubject("")
	updated := false
	const expected string = "New Value"
	subject.Attach(func(newVal string) {
		updated = true
	})

	subject.SetState(expected)
	if got, want := subject.GetState(), expected; got != want {
		t.Errorf("Expected new subject state to be '%s but got '%s'", want, got)
	}
	if got, want := updated, true; got != want {
		t.Errorf("Updated flag is %v but expected %v", got, want)
	}
}
