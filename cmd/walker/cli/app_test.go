package cli

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_App_Main(t *testing.T) {
	t.Parallel()

	const root = "testdata"
	const filesOnly = "a/a.md\na/b/b.md\na/b/c/c.md\n"
	const dirsOnly = "a\na/b\na/b/c\n"
	const all = "a\na/a.md\na/b\na/b/b.md\na/b/c\na/b/c/c.md\n"

	cab := os.DirFS(root)

	tcs := []struct {
		name string
		app  *App
		args []string
		exp  string
		err  bool
	}{
		{name: "files only/with cab", exp: filesOnly, args: []string{root}, app: &App{Cab: cab}},
		{name: "files only/without cab", exp: filesOnly, args: []string{root}, app: &App{}},
		{name: "dirs only", exp: dirsOnly, args: []string{"-dirs", "-skip-files", root}, app: &App{Cab: cab}},
		{name: "all", exp: all, args: []string{"-dirs", root}, app: &App{Cab: cab}},
		{name: "nil app", err: true},
	}

	for _, tc := range tcs {
		// snippet: test
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			bb := &bytes.Buffer{}

			app := tc.app

			if app != nil {
				app.IO.Out = bb
			}

			ctx := context.Background()
			ctx, cause := context.WithTimeout(ctx, time.Second)
			defer cause()

			err := app.Main(ctx, root, tc.args)

			if tc.err {
				r.Error(err)
				return
			}

			r.NoError(err)

			r.Equal(tc.exp, bb.String())
		})
		// snippet: test
	}

}
