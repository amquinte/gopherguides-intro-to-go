package week06

import (
	"testing"
)

func Test_Employee_IsValid(t *testing.T) {
	t.Parallel()
	e := Employee(1)
	act := e.IsValid()
	exp := error(nil)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	e = Employee(-1)
	act = e.IsValid()
	exp = ErrInvalidEmployee(-1)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}

func Test_Employee_Work_Err(t *testing.T) {
	t.Parallel()
	e := Employee(-1)
	m := NewManager()
	go e.work(m)
	act := m.Assign(&Product{Quantity: 1})
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
	act = <-m.Errors()
	exp = ErrInvalidEmployee(-1)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
}

func Test_Employee_Work_Ok(t *testing.T) {
	t.Parallel()
	e := Employee(1)
	m := NewManager()
	go e.work(m)
	act := m.Assign(&Product{Quantity: 1})
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
	<-m.Completed()
	// exp = ErrInvalidEmployee(-1)
	// if act != exp {
	// 	t.Fatalf("expected %v, got %v", exp, act)
	// }
	// fmt.Println("Test employee passed")
}
