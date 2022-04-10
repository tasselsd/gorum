package core

type NavStack struct {
	stack []map[string]string
}

func NewNavStack() *NavStack {
	s := NavStack{}
	return &s
}

func (s *NavStack) Push(name, uri string) *NavStack {
	link := make(map[string]string)
	link["name"] = name
	link["uri"] = uri
	s.stack = append(s.stack, link)
	return s
}

func (s *NavStack) List() []map[string]string {
	return s.stack
}
