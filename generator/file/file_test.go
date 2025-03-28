package file_test

import (
	"testing"

	"github.com/0x9n0p/og/generator/file"
	"github.com/0x9n0p/og/generator/function"
	"github.com/0x9n0p/og/generator/structure"
)

func TestGenerator(t *testing.T) {
	f := file.File{
		PackageName: "user",
		Imports: []file.Import{
			{
				Alias: "_",
				Path:  "github.com/example/example",
			},
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
				Methods: []structure.Method{
					{
						Name:         "Save",
						Arguments:    "",
						Returns:      "error",
						Body:         "\treturn nil",
						StructName:   "*Wallet",
						ReceiverName: "w",
					},
				},
			},
		},
		Functions: []function.Function{
			{
				Name:      "UpdateBalance",
				Arguments: "ctx context.Context, value int",
				Returns:   "error",
				Body:      "\treturn nil",
			},
		},
	}

	generated, err := f.Generate(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(generated)
}
