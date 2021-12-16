package week06

import (
	"fmt"
	"testing"
)

func Test_Manager_InvalidEmployeeCount(t *testing.T) {
	t.Parallel()
	m := NewManager()
	act := m.Start(0)
	exp := ErrInvalidEmployeeCount(0)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Tests stop function - hits if statement that returns
	//Manager can only be stopped once
	m.Stop()
	m.Stop()

	//Testing assign for ErrManagerStopped
	act = m.Assign()
	err := ErrManagerStopped{}
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}
}

func Test_Manager_ValidEmployeeCount(t *testing.T) {
	t.Parallel()
	m := NewManager()
	e := 1
	act := m.Start(e)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Testing assign when product is invalid
	q := 0
	p := Product{Quantity: q}
	act = m.Assign(&p)
	err := ErrInvalidQuantity(q)
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}

	//Stops manager and cleans up resources
	m.Stop()
}

func Test_Manager_ValidProduct(t *testing.T) {
	t.Parallel()
	m := NewManager()
	e := 1
	act := m.Start(e)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Testing assign when product is invalid
	q := 1
	p := Product{Quantity: q}
	act = m.Assign(&p)
	err := error(nil)
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}

	//Stops manager and cleans up resources
	m.Stop()
}

func Test_Manager_CompleteErrors(t *testing.T) {
	t.Parallel()
	m := NewManager()
	e := Employee(0)
	p := Product{Quantity: 1}
	act := m.Complete(e, &p)
	exp := ErrInvalidEmployee(0)

	//Tests for when invalid employee
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Tests when product has not been built
	e = Employee(1)
	act = m.Complete(e, &p)
	err := ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", p))
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}
}
