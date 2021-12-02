package assignment04

import (
	"testing"
)

func Test_Beatles_Name(t *testing.T) {
	t.Parallel()

	a := Beatles{}
	expected := "Member forgot their name"
	act := a.Name()
	if act != expected {
		t.Fatalf("expected: %s but received %s", expected, act)
	}

	a.Member = "Paul"
	expected = a.Member
	act = a.Name()
	if act != a.Member {
		t.Fatalf("expected: %s but received %s", expected, act)
	}
}

func Test_Beatles_Perform(t *testing.T) {
	t.Parallel()

	v := Venue{Audience: 15000}
	b := Beatles{
		Member:      "John",
		IsJohn:      true,
		MinAudience: 10000,
	}

	err := b.Perform(v)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if b.Performed != true {
		t.Fatalf("expected true but got false")
	}
}

func Test_Beatles_Setup(t *testing.T) {
	t.Parallel()
	v := Venue{Audience: 15000}
	b := Beatles{
		Member:      "Paul",
		IsJohn:      false,
		MinAudience: 10000,
	}

	err := b.Setup(v)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if b.Helped == false {
		t.Fatalf("expected true but got false")
	}
}
