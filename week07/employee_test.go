package week07

import (
	"context"
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
	defer m.Stop()
	go e.work(context.Background(), m)
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
	defer m.Stop()
	go e.work(context.Background(), m)
	act := m.Assign(&Product{Quantity: 1})
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
	cp := <-m.Completed()
	err := cp.IsValid()
	if err != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
}
