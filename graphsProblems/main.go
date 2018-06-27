package main

import (
	"fmt"

	"github.com/weberr13/Kata/graph"
)
func main() {
	g := graph.NewGraph()

	g.AddEdge("robert", "hannah")
	g.AddEdge("robert", "micah")
	g.AddEdge("robert", "noah")
	g.AddEdge("bonnie", "robert")
	g.AddEdge("micah", "noah")

	fmt.Println("BFS:")
	fmt.Println(g.BFS("robert"))
	fmt.Println(g.BFS("bonnie"))
	fmt.Println("DFS")
	fmt.Println(g.DFS("robert"))
	fmt.Println(g.DFS("bonnie"))
}
