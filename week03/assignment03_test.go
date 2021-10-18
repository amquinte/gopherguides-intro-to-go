package main

import (
	"testing"
)

func TestMovieRateError(t *testing.T) {

	m := Movie{Length: 120, Name: "Taken", plays: 0}
	act := m.Rate(75.00)
	exp := "can't review a movie without watching it first"

	//Checks if act is nill after calling Rate() func
	//act should have a value since number of plays is 0
	if act == nil {
		t.Errorf("Was expecting Rate() to return the following statement: %s", exp)
	}
}

func TestMovieRate(t *testing.T) {

	m := Movie{Length: 120, Name: "Taken", plays: 1}
	act := m.Rate(75.00)

	//Checks if the return value of Rate() is nill
	//Should be nill since movie has at least 1 play
	if act != nil {
		t.Errorf("Was expecting error to be nil but instead is %e", act)
	}
}

func TestMoviePlay(t *testing.T) {
	m := Movie{Length: 120, Name: "Taken", plays: 0}
	v := 300
	m.Play(v)
	actV := m.views
	actP := m.plays

	if actV != v {
		t.Errorf("Was expecting %d but received %d", v, actV)
	}

	if actP != 1 {
		t.Errorf("Was expecting 1 but received %d", actP)
	}

	m.Play(v)

	if actV != v {
		t.Errorf("Was expecting %d but received %d", v, actV)
	}

	if actP != 1 {
		t.Errorf("Was expecting 2 but received %d", actP)
	}
}

func TestMovieViewers(t *testing.T) {
	m := Movie{Length: 120, Name: "Taken", plays: 0}
	v := 300
	m.Play(v)
	actV := m.views
	result := m.Viewers()

	if result != actV {
		t.Errorf("Was expecting %d but actually got %d", v, result)
	}

	m.views = 1000
	result = m.Viewers()
	actV = m.views

	if result != actV {
		t.Errorf("Was expecting %d but actually got %d", v, result)
	}
}

func TestMoviePlays(t *testing.T) {
	m := Movie{Length: 120, Name: "Taken", plays: 0}
	v := 300
	m.Play(v)
	actP := m.plays
	result := m.Plays()

	if result != actP {
		t.Errorf("Was expecting %d but actually got %d", 1, result)
	}

	m.plays = 1000
	result = m.Plays()
	actP = m.plays

	if result != actP {
		t.Errorf("Was expecting %d but actually got %d", 1000, result)
	}
}

func TestMovieString(t *testing.T) {
	m := Movie{Length: 102, Name: "Wizard of Oz", rating: 99.0}
	s := m.String()
	exp := "Wizard of Oz (102m) 99.0%"

	if s != exp {
		t.Errorf("Was expecting %s but received %s", exp, s)
	}
}
