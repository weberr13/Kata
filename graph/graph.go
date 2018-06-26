package graph

//Graph of strings
type Graph struct {
	theMap map[string][]string
}

//NewGraph single directional
func NewGraph() (g *Graph) {
	return &Graph{theMap: make(map[string][]string)}
}

//AddEdge to the graph
func (g Graph) AddEdge(node, edge string) {
	g.theMap[node] = append(g.theMap[node], edge)
}

//BFS Breadth first Search
func (g Graph) BFS(start string) (nodes []string) {
	visited := make(map[string]struct{})
	visited[start] = struct{}{}
	active := []string{start}
	var last string
	for len(active) > 0 {
		last, active = active[len(active)-1], active[:len(active)-1]
		nodes = append(nodes, last)
		for _, next := range g.theMap[last] {
			_, ok := visited[next]
			if !ok {
				visited[next] = struct{}{}
				active = append(active,next)
			}
		}
	}
	return nodes
}