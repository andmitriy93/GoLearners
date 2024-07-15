package main

import "fmt"

var (
	a, b, c *string
	d  *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("a value %v\t and type %T\n", a, a)
	fmt.Printf("b value %v\t and type %T\n", b, b)
	fmt.Printf("c value %v\t and type %T\n", c, c)
	fmt.Printf("d value %v\t and type %T\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}

