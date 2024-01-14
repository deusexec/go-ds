# Go DS

Useful golang data structures for every day.

## Graph

```go
package main

import (
    "fmt"

    "github.com/deusexec/go-ds/graph"
)

func main() {
    nodeA := graph.NewNode("A", "A")
    nodeB := graph.NewNode("B", "B")
    nodeC := graph.NewNode("C", "C")
    nodeD := graph.NewNode("D", "D")
    nodeE := graph.NewNode("E", "E")
    nodeF := graph.NewNode("F", "F")

    g := graph.NewGraph(nodeA, nodeB, nodeC, nodeD)

    g.AddNode(nodeE)
    g.AddNode(nodeF)

    g.AddEdge(nodeA, nodeB)
    g.AddEdge(nodeA, nodeC)
    g.AddEdge(nodeA, nodeD)
    g.AddEdge(nodeB, nodeA)

    g.AddEdge(nodeC, nodeF)
    g.AddEdge(nodeF, nodeE)
    g.AddEdge(nodeE, nodeD)
    g.AddEdge(nodeD, nodeA)

    g.Dump()

    traversalCallback := func(node *graph.Node[string, string]) {
        fmt.Printf("[%s] ", node.Key())
    }

    BFS := func(start *graph.Node[string, string]) {
        fmt.Print("BFS: ")
        g.Traversal(graph.BFS, start, traversalCallback)
        fmt.Println()
    }

    DFS := func(start *graph.Node[string, string]) {
        fmt.Print("DFS: ")
        g.Traversal(graph.DFS, start, traversalCallback)
        fmt.Println()
    }

    BFS(nodeA)
    BFS(nodeD)

    DFS(nodeA)
    DFS(nodeD)
}
```

## Output

```text
$ go run .
===============
==   Graph   ==
===============
Node: [A] Edges: [A]->[B] [A]->[C] [A]->[D] 
Node: [B] Edges: [B]->[A] 
Node: [C] Edges: [C]->[F] 
Node: [D] Edges: [D]->[A] 
Node: [E] Edges: [E]->[D] 
Node: [F] Edges: [F]->[E] 
===============
BFS: [A] [B] [C] [D] [F] [E] 
BFS: [D] [A] [B] [C] [F] [E] 
DFS: [A] [D] [C] [F] [E] [B] 
DFS: [D] [A] [C] [F] [E] [B] 
```
