package must_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/iwittkau/th/must"
)

func TestReturn_Examples(t *testing.T) {
	data := must.Return[any](t, []string{"test"}, json.Marshal)
	t.Log(string(data))

	f := must.Return(t, "must.go", os.Open)
	defer f.Close()

	res := must.ReturnCtx(t, context.Background(), "test", example)
	t.Log(res)
}

func example(_ context.Context, in string) (string, error) {
	return in, nil
}
