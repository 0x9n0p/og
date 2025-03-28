package structure

import "fmt"

type Method struct {
	Name         string
	Arguments    string
	Returns      string
	Body         string
	StructName   string
	ReceiverName string
}

func (m *Method) Code() string {
	return fmt.Sprintf(`func (%s %s) %s(%s) %s {
%s
}`, m.ReceiverName, m.StructName, m.Name, m.Arguments, m.Returns, m.Body)
}
