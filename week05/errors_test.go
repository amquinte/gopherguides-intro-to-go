package demo

import (
	"fmt"
	"testing"
)

func Test_ErrorTableNotFound(t *testing.T) {
	t.Parallel()
	x := ErrTableNotFound{table: "test"}
	c1 := "error()"
	c2 := "tableNotFound()"
	c3 := "is()"
	c4 := "isErrTableNotFound()"

	table := []struct {
		name  string
		table string
		exp   string
		res   bool
		err   error
	}{
		{name: c1, table: "test", exp: "table not found test", err: nil},
		{name: c2, table: "test", exp: "test", err: nil},
		{name: c3, table: "test", res: false, err: nil},
		{name: c4, table: "test", res: false, err: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == c1 {
				act := x.Error()
				if tt.exp != act {
					t.Fatalf("expected %s but got %s instead", tt.exp, act)
				}
			}

			if tt.name == c2 {
				act := x.TableNotFound()
				if tt.exp != act {
					t.Fatalf("expected %s but got %s instead", tt.exp, act)
				}
			}

			if tt.name == c3 {
				act := x.Is(tt.err)
				if tt.res != act {
					t.Fatalf("expected %t but got %t instead", tt.res, act)
				}
			}

			if tt.name == c4 {
				act := IsErrTableNotFound(tt.err)
				if tt.res != act {
					t.Fatalf("expected %t but got %t instead", tt.res, act)
				}
			}

		})
	}
}

func Test_Error(t *testing.T) {
	t.Parallel()
	x := errNoRows{clauses: Clauses{}, table: "test"}
	exp := "[test] no rows found\nquery: "
	act := x.Error()
	if exp != act {
		t.Fatalf("expected %s but got %s instead", exp, act)
	}
}

func Test_Clauses(t *testing.T) {
	t.Parallel()
	x := &errNoRows{}
	act := x.Clauses()
	if act == nil {
		t.Fatalf("was expecting %v but got nil", act)
	}
}

func Test_RowNotFound(t *testing.T) {
	t.Parallel()
	x := &errNoRows{clauses: Clauses{"brand": "ford"}, table: "test"}
	act, c := x.RowNotFound()
	exp := "test"
	if exp != act {
		t.Fatalf("expected %v but got %v", exp, act)
	}

	expC := 1
	if expC != len(c) {
		t.Fatalf("expected %v but got %v", expC, c)
	}

}

func Test_IsErrNoRow(t *testing.T) {
	t.Parallel()
	x := &errNoRows{}
	exp := true
	act := IsErrNoRows(x)
	if exp != act {
		t.Fatalf("expected %v but got %v", exp, act)
	}

	exp = false
	act = IsErrNoRows(nil)
	if exp != act {
		t.Fatalf("expected %v but got %v", exp, act)
	}
}

func Test_AsErrNoRows(t *testing.T) {
	t.Parallel()
	err := fmt.Errorf("some error")
	_, act := AsErrNoRows(err)
	exp := false
	if act != exp {
		t.Fatalf("expected %v but got %v", exp, act)
	}

	x := &errNoRows{}
	_, act = AsErrNoRows(x)
	exp = true
	if act != exp {
		t.Fatalf("expected %v but got %v", exp, act)
	}
}
