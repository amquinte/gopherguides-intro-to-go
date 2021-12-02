package assignment04

import (
	"fmt"
)

type Beatles struct {
	Member      string
	Performed   bool
	Helped      bool
	MinAudience int
}

func (b Beatles) Name() string {
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
	if b.IsJohn() {
		b.Helped = false
		return nil
	}

	if b.Helped == true {
		return fmt.Errorf("Member already helped setup")
	}

	b.Helped = true
	return nil
}

func (b Beatles) IsJohn() bool {
	if b.Name() == "John" {
		return true
	}

	return false
}
