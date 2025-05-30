package main

import "fmt"

func maps() {
	var x map[string]int
	fmt.Println(x)

	m := make(map[string]int)
	fmt.Println(m)

	n := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(n)
	n["d"] = 4
	fmt.Println(n)
	o := n["c"]
	fmt.Println(o)
	p := n["p"]
	fmt.Println(p)

	q, r := n["a"]
	fmt.Println(q, r)
	delete(n, "d")
	fmt.Println(n)

	for k, v := range n {
		fmt.Println(k, v)
	}

	var l map[string]int
	fmt.Println(l)
	fmt.Println(l == nil)
	// l["key"] = 10  // will give error

	a := map[string]int{"x": 1}
	b := a
	b["x"] = 100
	fmt.Println(a["x"]) // Output: 100 (a is affected)

}
