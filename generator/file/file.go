package file

import (
	"fmt"

	"github.com/0x9n0p/og/generator/structure"
)

type File struct {
	PackageName string
	Imports     []Import
	Structures  []structure.Structure
	Generator   *Generator
}

func (g *File) Generate() error {
	if err := g.Generator.PackageName(g.PackageName); err != nil {
		return fmt.Errorf("set package name: %w", err)
	}

	if g.Imports != nil && len(g.Imports) > 0 {
		for _, p := range g.Imports {
			if err := g.Generator.Import(p); err != nil {
				return fmt.Errorf("add import: %w", err)
			}
		}
	}

	if g.Structures != nil && len(g.Structures) > 0 {
		for _, s := range g.Structures {
			if err := s.Generate(); err != nil {
				return fmt.Errorf("generate structure: %w", err)
			}

			if err := g.Generator.Struct(s.Generator.Content); err != nil {
				return fmt.Errorf("append structure: %w", err)
			}
		}
	}

	if err := g.Generator.Finalize(); err != nil {
		return fmt.Errorf("finalize: %w", err)
	}

	return nil
}
