package structure

import (
	"context"
	"fmt"
)

type Structure struct {
	StructName string
	Fields     []Field
	Methods    []Method
}

func (g *Structure) Generate(ctx context.Context) (string, error) {
	generated := fmt.Sprintf("type %s struct {", g.StructName)

	for _, f := range g.Fields {
		s, err := f.Generate(ctx)
		if err != nil {
			return "", fmt.Errorf("generate field %s: %w", f.Name, err)
		}

		generated = generated + fmt.Sprintf("\n\t%s", s)
	}

	generated += "\n}"

	if g.Methods != nil && len(g.Methods) > 0 {
		for _, m := range g.Methods {
			s, err := m.Generate(ctx)
			if err != nil {
				return "", fmt.Errorf("generate method %s: %w", m.Name, err)
			}

			generated = generated + fmt.Sprintf("\n\n%s", s)
		}
	}

	return generated, nil
}
