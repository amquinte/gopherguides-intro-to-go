package assignment04

import (
	"bytes"
	"testing"
)

func Test_Venue_Entertain_Crew(t *testing.T) {
	t.Parallel()

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	c := Crew{}

	err := v.Entertain(0, &c)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	act := bb.String()
	exp := "Crew has performed for 0 people.\nCrew has completed teardown.\n"

	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func Test_Venue_Entertain_Paul(t *testing.T) {
	t.Parallel()

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	p := Beatles{
		Member:      "Paul",
		MinAudience: 100,
	}

	err := v.Entertain(1000, &p)

	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	act := bb.String()
	exp := "Paul has completed setup.\nPaul has performed for 1000 people.\n"

	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func Test_Venue_Entertain_Paul_Error(t *testing.T) {
	t.Parallel()

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	p := Beatles{
		Member:      "Paul",
		MinAudience: 1000,
	}

	err := v.Entertain(100, &p)

	if err == nil {
		t.Fatalf("expected error %s, got nil", err)
	}

}

func Test_Venue_Entertain_John(t *testing.T) {
	t.Parallel()

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	j := Beatles{
		Member:      "John",
		MinAudience: 100,
	}

	err := v.Entertain(1000, &j)

	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	act := bb.String()
	exp := "John has performed for 1000 people.\n"

	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func Test_Venue_Entertain_Crew_Error(t *testing.T) {
	t.Parallel()
	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	c := Crew{}

	err := v.Entertain(100, &c)

	if err == nil {
		t.Fatalf("expected error %s, got nil", err)
	}
}

func Test_Venue_Entertain_Crew_Teardown_Error(t *testing.T) {
	t.Parallel()
	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	c := Crew{Completed: true}

	err := v.Entertain(0, &c)

	if err == nil {
		t.Fatalf("expected error %s, got nil", err)
	}
}
func Test_Venue_Entertain_Helped_Error(t *testing.T) {
	t.Parallel()
	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}
	b := Beatles{Member: "Ringo", Helped: true}

	err := v.Entertain(100, &b)
	if err == nil {
		t.Fatalf("expected error %s, got nil", err)
	}
}

func Test_Venue_No_Acts_Error(t *testing.T) {
	t.Parallel()
	v := Venue{}
	err := v.Entertain(100)

	if err == nil {
		t.Fatalf("expected error %s, got nil", err)
	}
}
