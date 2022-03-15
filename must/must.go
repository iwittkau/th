package must

import (
	"context"
	"testing"
)

func Return[v any, t any](tb testing.TB, in v, from func(v) (t, error)) t {
	tb.Helper()
	res, err := from(in)
	if err != nil {
		tb.Fatal(err)
	}

	return res
}

func ReturnCtx[v any, t any](tb testing.TB, ctx context.Context, in v, from func(context.Context, v) (t, error)) t {
	tb.Helper()
	res, err := from(ctx, in)
	if err != nil {
		tb.Fatal(err)
	}

	return res
}
