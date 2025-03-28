package function

import "fmt"

type Function struct {
	Name      string
	Arguments string
	Returns   string
	Body      string
}

func (f *Function) Code() string {
	return fmt.Sprintf(`func %s(%s) %s {
%s
}`, f.Name, f.Arguments, f.Returns, f.Body)
}
