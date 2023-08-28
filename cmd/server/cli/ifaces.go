package cli

import "context"

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}
