package mocks

type MockGenerator struct {
	name string
}

func (g *MockGenerator) Generate() string {
	return g.name
}

func NewMockGenerator(n string) *MockGenerator {
	return &MockGenerator{name: n}
}
