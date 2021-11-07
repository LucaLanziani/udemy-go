package main

import "fmt"

// OCP
// open for extention, closed for modification
// Specification pattern

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

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	isSatisfied(p Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) isSatisfied(p Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p Product) bool {
	return p.size == s.size
}

type AndSecification struct {
	first, second Specification
}

func (a AndSecification) isSatisfied(p Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.isSatisfied(v) {
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
	fmt.Printf("Green Products (old):\n")

	g := ColorSpecification{green}
	s := SizeSpecification{large}
	a := AndSecification{g, s}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, a) {
		fmt.Printf(" - %v", v)
	}
}
