package structure_test

import (
	"testing"

	structure2 "github.com/0x9n0p/og/pkg/generator/structure"
)

func TestGenerator_AddField(t *testing.T) {
	g := structure2.Structure{
		StructName: "User",
		Fields: []structure2.Field{
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
		Methods: []structure2.Method{
			{
				Name:         "Create",
				Arguments:    "ctx context.Context",
				Returns:      "error",
				Body:         "\treturn nil",
				StructName:   "*User",
				ReceiverName: "u",
			},
			{
				Name:         "Save",
				Arguments:    "ctx context.Context",
				Returns:      "error",
				Body:         "\treturn nil",
				StructName:   "*User",
				ReceiverName: "u",
			},
		},
	}

	s, err := g.Generate(t.Context())
	if err != nil {
		t.Fatalf("generate: %v", err)
	}

	t.Log(s)
}
