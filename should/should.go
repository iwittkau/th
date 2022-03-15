package should

import (
	"context"
	"testing"
)

func Return[v any, t any](tb testing.TB, in v, f func(v) (t, error)) t {
	tb.Helper()
	res, err := f(in)
	if err != nil {
		tb.Error(err)
	}

	return res
}

func ReturnCtx[v any, t any](tb testing.TB, ctx context.Context, in v, f func(context.Context, v) (t, error)) t {
	tb.Helper()
	res, err := f(ctx, in)
	if err != nil {
		tb.Error(err)
	}

	return res
}
