package cli

import (
	"context"
	"io"

	"github.com/markbates/iox"
)

type Commander interface {
	Main(ctx context.Context, pwd string, args []string) error
}

type Printer interface {
	Print(w io.Writer) error
}

type Describer interface {
	Describe() string
}

type SettableIO interface {
	SetIO(io iox.IO)
}
