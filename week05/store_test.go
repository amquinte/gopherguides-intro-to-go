package demo

import (
	"errors"
	"testing"
)

func Test_Store_Db(t *testing.T) {
	s := &Store{}
	act := s.db()
	if act == nil {
		t.Fatalf("Was expecting non-nil value")
	}
}

func Test_Store_All(t *testing.T) {
	s := &Store{}
	tn := "cars"
	_, err := s.All(tn)
	exp := ErrTableNotFound{table: tn}
	//expecting ErrTableNotFound
	if err == exp {
		t.Fatalf("expected:%s got:%s", exp, err)
	}
}

func Test_Store_Insert(t *testing.T) {
	s := &Store{}
	tn := "cars"
	act, _ := s.Len(tn)
	exp := 0
	if act != exp {
		t.Fatalf("expected %d got %d", exp, act)
	}

	s.Insert(tn, Model{"Brand": "Ford", "Model": "Mustang", "Year": 2020})
	act, _ = s.Len(tn)
	exp = 1
	if act != exp {
		t.Fatalf("expected %d got %d", exp, act)
	}
}

func Test_Store_Select(t *testing.T) {
	s := &Store{}
	q := Clauses{}
	_, err := s.Select("", q)

	//tests if an error is thrown
	if err == nil {
		t.Fatalf("expected non nil value")
	}

	tn := "cars"
	s.Insert(tn, Model{"Brand": "Ford", "Model": "Mustang", "Year": 2020})
	_, err = s.Select(tn, q)

	//tests len of query
	if err != nil {
		t.Fatalf("length of query is greater than 0")
	}

	//tests when no rows are found
	x := Clauses{"vin": 143566735}
	_, err = s.Select(tn, x)
	exp := &errNoRows{clauses: x, table: tn}
	res := errors.Is(err, exp)
	if !res {
		t.Fatalf("was expecting, %v but got %v instead", exp, err)
	}

	//tests when there is a query
	q = Clauses{"Brand": "Ford"}
	_, err = s.Select(tn, q)
	if err != nil {
		t.Fatalf("expected no error but got %s", err)
	}
}
