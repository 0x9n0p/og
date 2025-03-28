package file

import (
	"context"
	"fmt"
)

type Import struct {
	Alias string
	Path  string
}

func (f *Import) Generate(ctx context.Context) (string, error) {
	if f.Alias != "" {
		f.Alias = f.Alias + " "
	}

	return fmt.Sprintf("%s%s", f.Alias, f.Path), nil
}
