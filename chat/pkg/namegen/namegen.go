package namegen

import (
	"fmt"
	"math/rand"
)

type Generator struct{}

func (g *Generator) Generate() string {
	colors := []string{"Красный", "Зелёный", "Жёлтый", "Розовый", "Голубой", "Оранжевый"}

	return fmt.Sprintf("%s чат №%d", colors[rand.Intn(len(colors))], rand.Intn(10000)+1)
}

func NewGenerator() *Generator {
	return &Generator{}
}
