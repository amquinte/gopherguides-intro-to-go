package assignment04

import (
	"fmt"
)

type Beatles struct {
	Member      string
	IsJohn      bool
	Performed   bool
	Helped      bool
	MinAudience int
}

func (b Beatles) Name() string {
	if b.Member == "" {
		return "Member forgot their name"
	}
	return b.Member
}

func (b *Beatles) Perform(v Venue) error {
	if b.MinAudience > v.Audience {
		b.Performed = false
		return fmt.Errorf("John demands a larger audience")
	}
	b.Performed = true
	return nil
}

func (b *Beatles) Setup(v Venue) error {
	if b.IsJohn {
		return fmt.Errorf("John does not help setup")
	}

	if b.Helped == true {
		return fmt.Errorf("Member already helped setup")
	}

	b.Helped = true
	return nil
}
