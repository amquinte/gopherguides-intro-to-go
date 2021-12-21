package week08

import (
	"context"
	"fmt"
	"testing"
)

func Test_Manager_InvalidEmployeeCount(t *testing.T) {
	t.Parallel()
	m := &Manager{stopped: true}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, act := m.Start(ctx, 0)
	exp := ErrInvalidEmployeeCount(0)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Testing assign for ErrManagerStopped
	act = m.Assign()
	err := ErrManagerStopped{}
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}
	m.Stop()
}

func Test_Manager_ValidEmployeeCount(t *testing.T) {
	t.Parallel()
	m := &Manager{}
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
	p := Product{}
	act = m.Assign(&p)
	err := ErrInvalidMaterials(q)
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}

	//Stops manager and cleans up resources
	m.Stop()
}

func Test_Manager_ValidProduct(t *testing.T) {
	t.Parallel()
	m := &Manager{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	e := 1
	_, act := m.Start(ctx, e)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Testing assign when product is valid
	p := ProductA
	act = m.Assign(p)
	err := error(nil)
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}

	//Stops manager and cleans up resources
	m.Stop()
}

func Test_Manager_Complete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	m := &Manager{}
	ctx, err := m.Start(ctx, 2)
	if err != nil {
		t.Fatal("unexpected error")
	}

	go m.Assign(ProductA)
	go m.Assign(ProductB)

	var completed []CompletedProduct

	go func() {
		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= 2 {
				m.Stop()
			}
		}
	}()

	<-ctx.Done()
}

func Test_Manager_CompleteErrors(t *testing.T) {
	t.Parallel()
	m := Manager{}
	e := Employee(0)
	p := ProductA
	act := m.Complete(e, p)
	exp := ErrInvalidEmployee(0)

	//Tests for when invalid employee
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}

	//Tests when product has not been built
	e = Employee(1)
	act = m.Complete(e, p)
	err := ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", p))
	if act != err {
		t.Fatalf("expected %v, got %v", err, act)
	}
}
