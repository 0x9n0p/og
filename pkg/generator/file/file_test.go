package file_test

import (
	"testing"

	file2 "github.com/0x9n0p/og/pkg/generator/file"
	"github.com/0x9n0p/og/pkg/generator/function"
	structure2 "github.com/0x9n0p/og/pkg/generator/structure"
)

func TestGenerator(t *testing.T) {
	f := file2.File{
		PackageName: "user",
		Imports: []file2.Import{
			{
				Alias: "_",
				Path:  "github.com/example/example",
			},
		},
		Structures: []structure2.Structure{
			{
				StructName: "User",
				Fields: []structure2.Field{
					{
						Name: "Username",
						Type: "string",
						Tags: `json:"username"`,
					},
				},
			},
			{
				StructName: "Wallet",
				Fields: []structure2.Field{
					{
						Name: "Balance",
						Type: "int",
						Tags: `json:"balance"`,
					},
				},
				Methods: []structure2.Method{
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
