package structure_test

import (
	"testing"

	"github.com/0x9n0p/og/generator/structure"
)

func TestGenerator_AddField(t *testing.T) {
	g := structure.Structure{
		StructName: "User",
		Fields: []structure.Field{
			{
				Type: "gorm.Model",
				Tags: `json:"-"`,
			},
			{
				Name: "ID",
				Type: "uuid.UUID",
				Tags: `json:"id" gorm:"primary_key"`,
			},
			{
				Name: "Username",
				Type: "string",
				Tags: `json:"username" gorm:"not null"`,
			},
		},
		Generator: structure.NewGenerator(),
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("generate: %v", err)
	}

	t.Log(g.Generator.Content)
}
