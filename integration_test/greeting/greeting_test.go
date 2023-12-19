package greeting_test

import (
	"context"
	"io"
	"net/http"
	"sample_app/config"
	"sample_app/integration_test/greeting"
	"testing"

	"github.com/k1LoW/httpstub"
)

func TestGreeting(t *testing.T) {
	ctx := context.Background()
	ts := httpstub.NewServer(t)
	t.Cleanup(ts.Close)
	cfg := config.Greeting{
		URL: ts.URL,
	}
	u, err := greeting.New(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}

	// arrange
	ts.Path("/hello").Method(http.MethodPost).Response(http.StatusOK, nil)

	// act
	ctx = context.Background()
	err = u.Do(ctx)

	// assert
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	{
		reqs := ts.Requests()
		if len(reqs) != 1 {
			t.Errorf("got %v, want 1", len(reqs))
		}
		req := reqs[0]
		t.Cleanup(func() {
			req.Body.Close()
		})
		b, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := "Hello, World!"
		if got != want {
			t.Errorf("got %v, want %s", got, want)
		}
	}
}
