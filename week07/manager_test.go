package week07

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

// snippet: example
// TODO: Implement test cases for the Run function.
func Test_Run(t *testing.T) {
	t.Parallel()
	m := NewManager()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	// defer cancel()

	ctx, act := m.Start(ctx, 1)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	<-ctx.Done()

}

func Test_Manager_Run(t *testing.T) {
	t.Parallel()
	//Context will cancel if user interrupts or after 10 seconds have passed
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	p := Product{Quantity: 1}
	cp, err := Run(ctx, 1, &p)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	exp := p
	act := cp[0].Product
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
	<-ctx.Done()
}

func Test_Manager_InvalidEmployeeCount(t *testing.T) {
	t.Parallel()
	m := NewManager()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, act := m.Start(ctx, 0)
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	e := 1
	_, act := m.Start(ctx, e)
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

// func Test_Manager_ValidProduct(t *testing.T) {
// 	t.Parallel()
// 	m := NewManager()
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	e := 1
// 	_, act := m.Start(ctx, e)
// 	exp := error(nil)
// 	if act != exp {
// 		t.Fatalf("expected %v, got %v", exp, act)
// 	}

// 	//Testing assign when product is invalid
// 	q := 1
// 	p := Product{Quantity: q}
// 	act = m.Assign(&p)
// 	err := error(nil)
// 	if act != err {
// 		t.Fatalf("expected %v, got %v", err, act)
// 	}

// 	//Stops manager and cleans up resources
// 	m.Stop()
// }

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
