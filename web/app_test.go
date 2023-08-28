package web

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_App_ServeHTTP(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	a := App{}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	a.ServeHTTP(w, req)

	r.Equal(200, w.Code)

	exp := "Walker - File Ranger"
	act := w.Body.String()

	r.Contains(act, exp)
}
