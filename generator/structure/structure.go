package structure

import (
	"fmt"
)

type Structure struct {
	StructName string
	Fields     []Field
	Methods    []Method

	Generator *Generator
}

func (g *Structure) Generate() error {
	if err := g.Generator.StructName(g.StructName); err != nil {
		return fmt.Errorf("set struct name: %w", err)
	}

	if g.Fields != nil && len(g.Fields) > 0 {
		for _, field := range g.Fields {
			if err := g.Generator.Field(field); err != nil {
				return fmt.Errorf("add field: %w", err)
			}
		}
	}

	if g.Methods != nil && len(g.Methods) > 0 {
		for _, method := range g.Methods {
			if err := g.Generator.Method(method); err != nil {
				return fmt.Errorf("add method: %w", err)
			}
		}
	}

	if err := g.Generator.Finalize(); err != nil {
		return fmt.Errorf("finalize: %w", err)
	}

	return nil
}
