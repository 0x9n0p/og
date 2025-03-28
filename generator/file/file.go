package file

import (
	"context"
	"fmt"

	"github.com/0x9n0p/og/generator/function"
	"github.com/0x9n0p/og/generator/structure"
)

type File struct {
	PackageName string
	Imports     []Import
	Structures  []structure.Structure
	Functions   []function.Function
}

func (g *File) Generate(ctx context.Context) (string, error) {
	generated := fmt.Sprintf("package %s", g.PackageName)

	if g.Imports != nil && len(g.Imports) > 0 {
		generated += "\n\nimport ("

		for _, p := range g.Imports {
			ps, err := p.Generate(ctx)
			if err != nil {
				return "", fmt.Errorf("generate import: %w", err)
			}

			generated += "\n\t" + ps
		}

		generated += "\n)"
	}

	if g.Structures != nil && len(g.Structures) > 0 {
		for _, s := range g.Structures {
			ss, err := s.Generate(ctx)
			if err != nil {
				return "", fmt.Errorf("generate structure: %w", err)
			}

			generated += "\n\n" + ss
		}
	}

	if g.Functions != nil && len(g.Functions) > 0 {
		for _, s := range g.Functions {
			ss, err := s.Generate(ctx)
			if err != nil {
				return "", fmt.Errorf("generate function: %w", err)
			}

			generated += "\n\n" + ss
		}
	}

	return generated, nil
}
