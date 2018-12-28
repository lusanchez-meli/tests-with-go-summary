package main

import (
	"testing"
	"testing/quick"
)

func TestID(t *testing.T) {
	err := quick.Check(func(n ID) bool {
		return validate(n) == nil
	}, nil)
	if err != nil {
		t.Errorf("validate() failed: %v", err)
	}
}
