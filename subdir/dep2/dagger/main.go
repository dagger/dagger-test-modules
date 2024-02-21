package main

type Dep2 struct{}

func (m *Dep2) Fn() string {
	return "hi from dep2"
}
