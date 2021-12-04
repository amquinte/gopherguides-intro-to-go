package ns

import (
	"context"
	"fmt"
	"testing"
)

func Test_Ns(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	n := &Ns{}

	ctx, err := n.Start(ctx)
	if err != nil {
		t.Fatal("unexpected error")
	}

	fmt.Println("test passed")
}
