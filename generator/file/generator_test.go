package file_test

import (
	"testing"

	"github.com/0x9n0p/og/generator/file"
	"github.com/0x9n0p/og/generator/structure"
)

func TestGenerator(t *testing.T) {
	f := file.File{
		PackageName: "user",
		Imports: []file.Import{
			{Path: "github.com/example/example"},
		},
		Structures: []structure.Structure{
			{
				StructName: "User",
				Fields: []structure.Field{
					{
						Name: "Username",
						Type: "string",
						Tags: `json:"username"`,
					},
				},
				Generator: structure.NewGenerator(),
			},
			{
				StructName: "Wallet",
				Fields: []structure.Field{
					{
						Name: "Balance",
						Type: "int",
						Tags: `json:"balance"`,
					},
				},
				Generator: structure.NewGenerator(),
			},
		},
		Generator: file.NewGenerator(),
	}

	if err := f.Generate(); err != nil {
		t.Fatalf("generate: %v", err)
	}

	t.Log(f.Generator.Content)
}
