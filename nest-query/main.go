package main

import "fmt"

type Query struct {
	And   []Query
	Or    []Query
	Child *Query
	Value string
}

func main() {
	q := Query{
		And: []Query{
			{
				Child: &Query{Value: "foo"},
			},
		},
	}

	fmt.Println(q)
}
