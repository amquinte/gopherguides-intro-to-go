package week08

import (
	"testing"
)

func Test_Errors_ErrInvalidMaterials_Error(t *testing.T) {
	t.Parallel()
	err := ErrInvalidMaterials(0)
	exp := "materials must be greater than 0, got 0"
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %v but got %v instead", exp, act)
	}
}

func Test_Errors_ErrProductNotBuilt_Error(t *testing.T) {
	t.Parallel()
	err := ErrProductNotBuilt("")
	exp := ""
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %v but got %v instead", exp, act)
	}
}

func Test_Errors_ErrInvalidEmployee_Error(t *testing.T) {
	t.Parallel()
	err := ErrInvalidEmployee(-1)
	exp := "invalid employee number: -1"
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %v but got %v instead", exp, act)
	}
}

func Test_Errors_ErrInvalidEmployeeCount_Error(t *testing.T) {
	t.Parallel()
	err := ErrInvalidEmployeeCount(-1)
	exp := "invalid employee count: -1"
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %v but got %v instead", exp, act)
	}
}

func Test_Erros_ErrManagerStopped_Error(t *testing.T) {
	t.Parallel()
	err := ErrManagerStopped{}
	exp := "manager is stopped"
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %v but got %v instead", exp, act)
	}
}
