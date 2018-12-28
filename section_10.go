package main

import "fmt"

type ID uint

func validate(id ID) error {
	if id < 0 {
		return fmt.Errorf("invalid")
	}
	return nil
}
