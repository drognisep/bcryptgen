package data

// StringObserver enables a dependent object to be updated in response to a state change.
type StringObserver func(appliedValue string)

// StringSubject encapsulates an independent value to be observed.
type StringSubject struct {
	value     string
	observers []StringObserver
}

// Attach notifies the subject of the observer.
func (s *StringSubject) Attach(in StringObserver) {
	s.observers = append(s.observers, in)
}

// SetState sets the subject state to a new value.
func (s *StringSubject) SetState(newVal string) {
	s.value = newVal
	for _, o := range s.observers {
		o(newVal)
	}
}

// GetState gets the current state of the subject.
func (s *StringSubject) GetState() string {
	return s.value
}

// NewStringSubject creates an initialized subject.
func NewStringSubject(initialValue string) *StringSubject {
	return &StringSubject{value: initialValue}
}
