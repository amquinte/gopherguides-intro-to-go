package assignment04

import (
	"fmt"
)

type Crew struct {
	Completed bool
}

func (c Crew) Name() string {
	return "Crew"
}

func (c Crew) Perform(v Venue) error {
	if v.Audience > 0 {
		return fmt.Errorf("Crew are not allowed to perform when there is an audience")
	}
	return nil
}

func (c *Crew) Teardown(v Venue) error {
	if c.Completed {
		return fmt.Errorf("The job is already done")
	}

	c.Completed = true
	return nil
}
