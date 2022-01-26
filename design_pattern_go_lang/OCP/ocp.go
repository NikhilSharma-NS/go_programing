package main

import (
	"fmt"
)

// open for extension , closed for modification
// Specification

type Size int
type Color int

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	blue
	green
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(p []Product, color Color) []*Product {

	result := make([]*Product, 0)

	for i, v := range p {
		if v.color == color {
			result = append(result, &p[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(p []Product, size Size) []*Product {

	result := make([]*Product, 0)

	for i, v := range p {
		if v.size == size {
			result = append(result, &p[i])
		}
	}
	return result
}

//

type Specificition interface {
	isSatisfied(p *Product) bool
}

type ColorSpecition struct {
	color Color
}

func (c ColorSpecition) isSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecition struct {
	size Size
}

type AndSpecition struct {
	first, second Specificition
}

func (a AndSpecition) isSatisfied(p *Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}
func (s SizeSpecition) isSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct {
}

func (f *BetterFilter) Filter(p []Product, spec Specificition) []*Product {

	res := make([]*Product, 0)
	for i, v := range p {
		if spec.isSatisfied(&v) {
			res = append(res, &p[i])
		}

	}
	return res
}

func main() {
	apple := Product{"apple", green, small}
	tree := Product{"tree", green, large}
	house := Product{"house", red, large}

	product := []Product{apple, tree, house}
	fmt.Println("all product", product)
	filter := Filter{}
	for _, v := range filter.FilterByColor(product, green) {
		fmt.Printf("Green Product: %s\n", v.name)
	}

	// new
	fmt.Println("---------------------------")
	greenSp := ColorSpecition{green}
	largesizespec := SizeSpecition{large}
	andspec := AndSpecition{greenSp, largesizespec}

	filer_new := BetterFilter{}
	for _, v := range filer_new.Filter(product, greenSp) {
		fmt.Printf("Green Product: %s\n", v.name)

	}

	fmt.Println("---------------------------")

	for _, v := range filer_new.Filter(product, largesizespec) {
		fmt.Printf("large Product: %s\n", v.name)

	}
	fmt.Println("---------------------------")

	for _, v := range filer_new.Filter(product, andspec) {
		fmt.Printf("Green and large Product: %s\n", v.name)

	}
}
