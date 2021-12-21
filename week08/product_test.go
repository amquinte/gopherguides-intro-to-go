package week08

import (
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	e := Employee(1)
	p := Product{Materials: Materials{Metal: 1}, builtBy: e}
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
	w := &Warehouse{}
	p := Product{Materials: Materials{Metal: 1}, builtBy: e}
	act := p.Build(e, w)
	exp := error(nil)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Tests second if statement - e.IsValid()
	e = Employee(1)
	p = Product{builtBy: e}
	act = p.Build(e, w)
	exp = ErrInvalidMaterials(len(p.Materials))
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Testing for invalid employee
	e = Employee(-1)
	p = Product{Materials: Materials{Metal: 1}, builtBy: e}
	act = p.Build(e, w)
	exp = ErrInvalidEmployee(e)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}

func Test_ProductIsBuilt(t *testing.T) {
	t.Parallel()
	//Testing for when product has not been built by an employee
	p := Product{Materials: Materials{Metal: 1}}
	act := p.IsBuilt()
	exp := error(ErrProductNotBuilt("product is not built: [{metal:1x}]"))
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Testing for when product is invalid
	e := Employee(1)
	p = Product{builtBy: e}
	act = p.IsBuilt()
	exp = ErrInvalidMaterials(len(p.Materials))
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}

	//Testing good path
	p = Product{Materials: Materials{
		Wood: 2,
		Oil:  3,
	}, builtBy: e}
	act = p.IsBuilt()
	exp = error(nil)
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}
