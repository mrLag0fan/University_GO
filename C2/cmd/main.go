package main

import (
	"C2/internal/data"
	"C2/internal/graph_generator"
	"C2/internal/scunnner"
	"fmt"
)

func main() {
	V := scunnner.ReadV()
	M := scunnner.ReadM()

	fmt.Println("V:", V)
	fmt.Println("M:", M)

	graph_generator.GenerateGraph()

	fmt.Println(data.R)
}
