package bostongo

import (
	"bytes"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Walker(t *testing.T) {
	t.Parallel()

	const root = "testdata"
	const filesOnly = "a/a.md\na/b/b.md\na/b/c/c.md\n"
	const dirsOnly = "a\na/b\na/b/c\n"
	const all = "a\na/a.md\na/b\na/b/b.md\na/b/c\na/b/c/c.md\n"

	tcs := []struct {
		name string
		cab  fs.FS
		wk   Walker
		exp  string
		err  bool
	}{
		{name: "files only", cab: os.DirFS(root), exp: filesOnly},
		{name: "dirs only", cab: os.DirFS(root), exp: dirsOnly, wk: Walker{SkipFiles: true, PrintDirs: true}},
		{name: "all", cab: os.DirFS(root), exp: all, wk: Walker{PrintDirs: true}},
		{name: "nil fs", err: true},
		{name: "bad fs", cab: os.DirFS("bad"), err: true},
	}

	for _, tc := range tcs {
		// snippet: test
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			bb := &bytes.Buffer{}

			err := tc.wk.Walk(tc.cab, bb)
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
