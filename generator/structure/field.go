package structure

import "fmt"

type Field struct {
	Name string
	Type string
	Tags string
}

func (f *Field) Code() string {
	if f.Name != "" {
		f.Name = f.Name + " "
	}

	if f.Type != "" {
		f.Type = f.Type + " "
	}

	if f.Tags != "" {
		return fmt.Sprintf("%s%s`%s`", f.Name, f.Type, f.Tags)
	}

	return fmt.Sprintf("%s%s", f.Name, f.Type)
}
