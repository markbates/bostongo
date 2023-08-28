package bostongo

import (
	"fmt"
	"io"
	"io/fs"
)

type Walker struct {
	PrintDirs bool
	SkipFiles bool
}

// Walk walks the file system rooted at cab, writing the path of each file to w.
func (wk Walker) Walk(cab fs.FS, w io.Writer) error {
	if cab == nil {
		return fmt.Errorf("nil fs")
	}

	if w == nil {
		return fmt.Errorf("nil writer")
	}

	err := fs.WalkDir(cab, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		if d.IsDir() {
			if wk.PrintDirs {
				fmt.Fprintln(w, path)
			}

			return nil
		}

		if wk.SkipFiles {
			return nil
		}

		fmt.Fprintln(w, path)

		return nil
	})

	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}

	return err
}
