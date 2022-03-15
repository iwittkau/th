package should_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/iwittkau/th/should"
)

func TestReturn_Examples(t *testing.T) {
	data := should.Return[any](t, []string{"test"}, json.Marshal)
	t.Log(string(data))

	f := should.Return(t, "should.go", os.Open)
	defer f.Close()

	res := should.ReturnCtx(t, context.Background(), "test", example)
	t.Log(res)
}

func example(_ context.Context, in string) (string, error) {
	return in, nil
}
