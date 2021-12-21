package week08

import (
	"fmt"
	"testing"
)

func Test_Material_Duration(t *testing.T) {
	t.Parallel()
	m := Materials{Metal: 1}
	act := fmt.Sprint(m.Duration())
	exp := "5ms"
	if act != exp {
		t.Fatalf("expect %v, got %v", exp, act)
	}
}
