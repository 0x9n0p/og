package og

import (
	"context"
	"fmt"
	"io"
	"text/template"
)

type Generator struct {
	Template    *template.Template
	Payload     Payload
	Destination io.Writer
}

func (g *Generator) Generate(ctx context.Context) error {
	err := g.Template.Execute(g.Destination, g.Payload)
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}
