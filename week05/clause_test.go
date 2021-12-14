package demo

import (
	"testing"
)

func Test_Clause_String(t *testing.T) {
	t.Parallel()

	q := Clauses{}
	res := q.String()
	exp := ""
	if res != exp {
		t.Fatalf("expected %s but got %s", exp, res)
	}

	q = Clauses{"Brand": "Ford"}
	res = q.String()
	exp = `"Brand" = "Ford"`
	if res != exp {
		t.Fatalf("expected %s but got %s", exp, res)
	}
}

func Test_Clause_Match(t *testing.T) {
	t.Parallel()

	m := Model{"Brand": "Ford"}
	q := Clauses{"Brand": "Ford"}
	res := q.Match(m)
	exp := true
	if res != exp {
		t.Fatalf("expected %t but got %t", exp, res)
	}
}

func Test_Clause_Match_Error(t *testing.T) {
	t.Parallel()

	m := Model{}
	q := Clauses{"Brand": "Ford"}
	res := q.Match(m)
	exp := false
	if res != exp {
		t.Fatalf("expected %t but got %t", exp, res)
	}
}
