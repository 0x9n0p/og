package file

import (
	"fmt"
	"strings"
)

const Template = `package {{package_name}}

import (
	{{where_imports_must_place}}
)
{{where_structures_must_place}}
`

const (
	ImportsLocation     = "{{where_imports_must_place}}"
	PackageNameLocation = "{{package_name}}"
	StructuresLocation  = "{{where_structures_must_place}}"
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
		fmt.Sprintf("\t%s\n", ImportsLocation), "", 1,
	)

	g.Content = strings.Replace(
		g.Content,
		fmt.Sprintf("%s\n", StructuresLocation), "", 1,
	)

	return nil
}

func (g *Generator) PackageName(name string) error {
	if name == "" {
		return fmt.Errorf("invalid package name (%s)", name)
	}

	g.Content = strings.Replace(g.Content, PackageNameLocation, name, 1)
	return nil
}

func (g *Generator) Import(i Import) error {
	g.Content = strings.Replace(
		g.Content, ImportsLocation,
		fmt.Sprintf("%s\n\t%s", i.Code(), ImportsLocation), 1,
	)

	return nil
}

func (g *Generator) Struct(s string) error {
	g.Content = strings.Replace(
		g.Content, StructuresLocation,
		fmt.Sprintf("\n%s\n%s", s, StructuresLocation), 1,
	)

	return nil
}
