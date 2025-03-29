package generator

import (
	"context"
	"fmt"
	"io"
	"text/template"

	"github.com/0x9n0p/og/payload"
)

type Generator struct {
	Template    *template.Template
	Payload     payload.Payload
	Destination io.Writer
}

func (g *Generator) Generate(ctx context.Context) error {
	err := g.Template.Execute(g.Destination, g.Payload)
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}
