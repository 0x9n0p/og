package structure

import (
	"fmt"
	"strings"
)

const Template = `type {{struct_name}} struct {
    {{where_fields_must_place}}
}`

const (
	StructNameLocation = "{{struct_name}}"
	FieldsLocation     = "{{where_fields_must_place}}"
)

type Generator struct {
	Content string
}

func NewGenerator() *Generator {
	return &Generator{
		Content: Template,
	}
}

func (g *Generator) Finalize() error {
	g.Content = strings.Replace(
		g.Content,
		fmt.Sprintf("\t%s\n", FieldsLocation), "", -1,
	)

	return nil
}

func (g *Generator) StructName(name string) error {
	if name == "" {
		return fmt.Errorf("invalid struct name (%s)", name)
	}

	g.Content = strings.Replace(g.Content, StructNameLocation, name, 1)
	return nil
}

func (g *Generator) Field(f Field) error {
	g.Content = strings.Replace(
		g.Content, FieldsLocation,
		fmt.Sprintf("%s\n\t%s", f.Code(), FieldsLocation), 1,
	)

	return nil
}
