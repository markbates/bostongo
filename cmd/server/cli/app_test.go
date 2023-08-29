package cli

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_App_Main(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	a := &App{}
	srv := httptest.NewServer(a)
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	go func() {
		err := a.Main(ctx, "", nil)
		r.NoError(err)
	}()

	resp, err := http.Get(srv.URL)
	r.NoError(err)

	r.Equal(200, resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	r.NoError(err)

	act := string(b)

	exp := "Walker - File Ranger"

	r.Contains(act, exp)

	cancel()

	<-ctx.Done()

	err = ctx.Err()
	r.NotEqual(context.DeadlineExceeded, err)

	_, err = http.Get(srv.URL)
	r.Error(err)
}

func Test_App_Getenv(t *testing.T) {
	t.Parallel()

	os.Setenv("EDITOR", "code")

	tcs := []struct {
		name string
		c    *App
		k    string
		exp  string
	}{
		{name: "default", exp: "code", c: &App{}},
		{name: "missing key", k: "unknown", exp: "", c: &App{}},
		{name: "good key", exp: "vim", c: &App{
			Env: &Env{
				data: map[string]string{
					"EDITOR": "vim",
				},
			},
		}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			k := tc.k
			if len(k) == 0 {
				k = "EDITOR"
			}

			act := tc.c.Getenv(k)

			r.Equal(tc.exp, act)
		})
	}

}
