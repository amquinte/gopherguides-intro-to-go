package assignment04

import (
	//"fmt"
	"testing"
)

func Test_Crew_Name(t *testing.T) {
	t.Parallel()

	c := Crew{}
	act := c.Name()

	exp := "Crew"
	if act != exp {
		t.Fatalf("expected: %s but received %s", exp, act)
	}
}

func Test_Crew_Perform(t *testing.T) {
	t.Parallel()

	c := Crew{}
	v := Venue{Audience: 0}

	err := c.Perform(v)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
}

func Test_Crew_Teardown(t *testing.T) {
	t.Parallel()
	c := Crew{}
	v := Venue{Audience: 100}

	err := c.Teardown(v)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	exp := true
	act := c.Completed
	if exp != act {
		t.Fatalf("expected: true but received false")
	}
}
