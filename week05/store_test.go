package demo

import (
	"fmt"
	"testing"
)

func Test_Store_Db(t *testing.T) {
	s := &Store{}
	act := s.db()
	if act == nil {
		t.Fatalf("Was expecting non-nil value")
	}
	fmt.Println("end of test store db")
}

func Test_Store_All(t *testing.T) {
	s := &Store{}
	tn := "cars"
	_, err := s.All(tn)
	exp := ErrTableNotFound{table: tn}
	if err == exp {
		t.Fatalf("expected:%s got:%s", exp, err)
	}
	fmt.Println("end of test store all")
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
	fmt.Println("end of test store insert")
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
	// s2 := &Store{}
	// _, err = s2.Select(tn, q)
	// exp := errNoRows{clauses: q, table: tn}
	// if err == nil {
	// 	t.Fatalf("was expecting, %v but got nil instead", exp)
	// }
	//tests when there is a query
	q = Clauses{"Brand": "Ford"}
	_, err = s.Select(tn, q)
	if err != nil {
		t.Fatalf("expected no error but got %s", err)
	}

	fmt.Println("end of test store select")
}
