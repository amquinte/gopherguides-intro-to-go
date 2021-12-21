package week08

import (
	"context"
	"testing"
)

func Test_Warehouse_Start(t *testing.T) {
	t.Parallel()
	w := Warehouse{}
	ctx := context.Background()
	w.Start(ctx)
	w.Stop()
}
