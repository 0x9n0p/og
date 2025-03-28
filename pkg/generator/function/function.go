package function

import (
	"context"
	"fmt"
)

type Function struct {
	Name      string
	Arguments string
	Returns   string
	Body      string
}

func (f *Function) Generate(ctx context.Context) (string, error) {
	return fmt.Sprintf(`func %s(%s) %s {
%s
}`, f.Name, f.Arguments, f.Returns, f.Body), nil
}
