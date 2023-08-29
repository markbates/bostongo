package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/markbates/garlic"
	"github.com/markbates/iox"
	"github.com/stretchr/testify/require"
)

func garlicDir(t testing.TB) (*Garlic, string) {
	t.Helper()

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	fp := filepath.Join(dir, "garlic", "bostongo")
	main := filepath.Join(fp, "main.go")

	r := require.New(t)

	_, err = os.Stat(main)
	r.ErrorIs(err, os.ErrNotExist)

	r.NoError(os.MkdirAll(fp, 0755))
	f, err := os.Create(filepath.Join(fp, "go.mod"))
	r.NoError(err)
	fmt.Fprintln(f, "module demo")
	r.NoError(f.Close())

	g := &Garlic{}

	err = g.Main(context.Background(), dir, []string{})
	r.NoError(err)

	_, err = os.Stat(main)
	r.NoError(err)

	return g, dir
}

func Test_Garlic_Generator(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	g, dir := garlicDir(t)

	// can't run twice
	err := g.Main(context.Background(), dir, []string{})
	r.Error(err)
}

// snippet: garlic-works
func Test_Garlic_Works(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// generate the garlic directory
	// and the main.go file
	_, dir := garlicDir(t)

	// snippet: io
	bb := &strings.Builder{}
	oi := iox.IO{
		Out: bb,         // use the strings.Builder as STDOUT
		Err: io.Discard, // discard STDERR
	}
	// snippet: io

	clove := &garlic.Garlic{
		Cmd:  &App{},        // the App to run, if no local command is found
		FS:   os.DirFS(dir), // the filesystem to use
		IO:   oi,            // IO to be used
		Name: "bostongo",    // the name of the command to run
	}

	// call the walker command through garlic
	err := clove.Main(context.Background(), dir, []string{"walker"})
	r.NoError(err)

	// assert the output
	exp := `Hello from Garlic!
go.mod
go.sum
main.go`

	act := bb.String()
	act = strings.TrimSpace(act)

	r.Equal(exp, act)
}

// snippet: garlic-works
