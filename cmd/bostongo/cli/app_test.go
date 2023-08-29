package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/markbates/iox"
	"github.com/stretchr/testify/require"
)

// snippet: test
func Test_App_Main(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	bb := &bytes.Buffer{}
	oi := iox.IO{
		Out: bb,
	}

	app := &App{
		IO: oi,
	}

	ctx := context.Background()

	err := app.Main(ctx, "testdata", []string{"walker"})
	r.NoError(err)

	// assert the output
	exp := `a/a.md
a/b/b.md
a/b/c/c.md`

	act := bb.String()
	act = strings.TrimSpace(act)

	r.Equal(exp, act)
}

// snippet: test
