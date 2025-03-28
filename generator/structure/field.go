package structure

import (
	"context"
	"fmt"
)

type Field struct {
	Name string
	Type string
	Tags string
}

func (f *Field) Generate(ctx context.Context) (string, error) {
	if f.Name != "" {
		f.Name = f.Name + " "
	}

	if f.Type != "" {
		f.Type = f.Type + " "
	}

	if f.Tags != "" {
		return fmt.Sprintf("%s%s`%s`", f.Name, f.Type, f.Tags), nil
	}

	return fmt.Sprintf("%s%s", f.Name, f.Type), nil
}
