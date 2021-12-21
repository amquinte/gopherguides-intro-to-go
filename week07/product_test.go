package week07

import (
	"fmt"
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	e := Employee(1)
	p := Product{Quantity: 1, builtBy: e}
	act := p.BuiltBy()
	exp := e
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}

func Test_Product_Build(t *testing.T) {
	t.Parallel()

	//Tests first if statement - p.IsValid()
	e := Employee(1)
	p := Product{Quantity: 1, builtBy: e}
	act := p.Build(e)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Tests second if statement - e.IsValid()
	e = Employee(1)
	p = Product{Quantity: 0, builtBy: e}
	act = p.Build(e)
	exp = ErrInvalidQuantity(p.Quantity)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Tests good path
	e = Employee(-1)
	p = Product{Quantity: 1, builtBy: e}
	act = p.Build(e)
	exp = ErrInvalidEmployee(e)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}

func Test_ProductIsBuilt(t *testing.T) {
	t.Parallel()
	e := Employee(0)
	p := Product{Quantity: 1, builtBy: e}
	act := p.IsBuilt()
	exp := ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", p))
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}
