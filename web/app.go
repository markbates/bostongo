package web

import (
	"bytes"
	_ "embed"
	"html/template"
	"net/http"
	"os"

	"github.com/markbates/bostongo"
)

//go:embed template.html
var htmlTemplate string

type App struct {
	DefaultPath string
}

// snippet: serve

func (a App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http.Error(w, "nil request", http.StatusBadRequest)
		return
	}

	f, err := a.parseForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := a.walk(&f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// snippet: serve

func (a App) walk(f *form) error {
	wk := bostongo.Walker{
		PrintDirs: len(f.PrintDirs) > 0,
		SkipFiles: len(f.SkipFiles) > 0,
	}
	bb := &bytes.Buffer{}

	cab := os.DirFS(f.WalkPath)
	if err := wk.Walk(cab, bb); err != nil {
		return err
	}

	f.Results = bb.String()
	return nil
}

func (a App) parseForm(r *http.Request) (form, error) {
	vals := r.URL.Query()

	f := form{
		WalkPath:  vals.Get("walkPath"),
		PrintDirs: vals.Get("printDirs"),
		SkipFiles: vals.Get("skipFiles"),
	}

	dp := a.DefaultPath
	if len(dp) == 0 {
		dp = "testdata"
	}

	if len(f.WalkPath) == 0 {
		f.WalkPath = dp
	}

	if len(f.PrintDirs) > 0 {
		f.PrintDirs = "checked"
	}

	if len(f.SkipFiles) > 0 {
		f.SkipFiles = "checked"
	}

	return f, nil
}

type form struct {
	WalkPath  string
	PrintDirs string
	SkipFiles string
	Results   string
}
