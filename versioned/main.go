package main

type Versioned struct{}

func (m *Versioned) Hello() string {
	return "version 1"
}
