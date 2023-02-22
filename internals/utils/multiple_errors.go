package utils

import (
	"errors"
	"strings"
)

// MultipleErrs is an error that collects other errors, for when you want to do
// several things and then report all of them.
type MultipleErrs struct {
	errors []error
}

func (e *MultipleErrs) Add(err error) {
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

func (e *MultipleErrs) Ret() error {
	if e == nil || e.IsEmpty() {
		return nil
	}
	return e
}

func (e *MultipleErrs) IsEmpty() bool {
	return e.Len() == 0
}

func (e *MultipleErrs) Len() int {
	return len(e.errors)
}

func (e *MultipleErrs) Error() string {
	asStr := make([]string, len(e.errors))
	for i, x := range e.errors {
		asStr[i] = x.Error()
	}
	return strings.Join(asStr, ". ")
}

func (e *MultipleErrs) Is(target error) bool {
	for _, candidate := range e.errors {
		if errors.Is(candidate, target) {
			return true
		}
	}
	return false
}

func (e *MultipleErrs) As(target interface{}) bool {
	for _, candidate := range e.errors {
		if errors.As(candidate, target) {
			return true
		}
	}
	return false
}
