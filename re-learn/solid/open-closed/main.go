package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type ProductSpecification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type AndSpecification struct {
	first, second ProductSpecification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type Filter struct{}

func (f *Filter) Filter(products []Product, spec ProductSpecification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

type Product struct {
	name  string
	color Color
	size  Size
}

func main() {
	apple := Product{"apple", green, small}
	tree := Product{"tree", green, large}
	house := Product{"house", blue, large}

	products := []Product{apple, tree, house}
	fmt.Println("Green products:")
	greenSpec := ColorSpecification{green}
	f := Filter{}

	for _, v := range f.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	largGreenSpec := AndSpecification{greenSpec, largeSpec}

	fmt.Println("Large Green products:")
	for _, v := range f.Filter(products, largGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
