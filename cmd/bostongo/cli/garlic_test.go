package cli

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Garlic(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	dir, err := os.MkdirTemp("", "")
	r.NoError(err)

	defer os.RemoveAll(dir)

	fp := filepath.Join(dir, "garlic", "bostongo", "main.go")

	_, err = os.Stat(fp)
	r.ErrorIs(err, os.ErrNotExist)

	g := &Garlic{}

	err = g.Main(context.Background(), dir, []string{"-skip-tidy"})
	r.NoError(err)

	_, err = os.Stat(fp)
	r.NoError(err)

	err = g.Main(context.Background(), dir, []string{})
	r.Error(err)
}
