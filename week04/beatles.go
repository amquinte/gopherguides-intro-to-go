package assignment04

import (
	"fmt"
)

type Beatles struct {
	Member      string
	IsJohn      bool
	Performed   bool
	MinAudience int
}

func (b Beatles) Name() string {
	if b.Member == "" {
		return "Member forgot their name"
	}
	return ("Hi my name is " + b.Member)
}

func (b *Beatles) Perform(v Venue) error {
	if b.MinAudience > v.Audience {
		b.Performed = false
		return fmt.Errorf("John demands a larger audience")
	}
	b.Performed = true
	return nil
}

func (b *Beatles) Teardown(v Venue) error {
	if b.IsJohn {
		return fmt.Errorf("John does not teardown")
	}
	return nil
}
