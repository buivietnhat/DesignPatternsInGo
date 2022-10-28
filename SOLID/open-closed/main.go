package main

import "fmt"

type Size int

const (
	small Size = iota
	medium
	large
)

type Color int

const (
	red Color = iota
	green
	blue
)

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type Product struct {
	name  string
	color Color
	size  Size
}

type ProductFilter struct{}

func (f *ProductFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Printf("Green products:\n")
	greenSpec := ColorSpecification{green}
	f := ProductFilter{}
	for _, v := range f.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
}
