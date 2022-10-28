package main

import "fmt"

type Person1 struct {
	name, position string
}

type personMod func(*Person1)
type PersonBuilder1 struct {
	actions []personMod
}

func (b *PersonBuilder1) Called(name string) *PersonBuilder1 {
	b.actions = append(b.actions, func(p *Person1) {
		p.name = name
	})

	return b
}

func (b *PersonBuilder1) WorksAs(position string) *PersonBuilder1 {
	b.actions = append(b.actions, func(p *Person1) {
		p.position = position
	})

	return b
}

func (b *PersonBuilder1) Build() *Person1 {
	p := Person1{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder1{}
	p := b.Called("Nhat").WorksAs("Programmer").Build()
	fmt.Println(p)
}
