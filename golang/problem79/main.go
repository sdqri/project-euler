package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	ID       string
	OutEdges []Edge
}

func NewNode(id string) *Node {
	return &Node{
		ID:       id,
		OutEdges: make([]Edge, 0),
	}
}

func (from *Node) AddEdgeTo(to *Node, weight int64) {
	for _, edge := range from.OutEdges {
		if edge.To == to {
			return
		}
	}
	from.OutEdges = append(from.OutEdges, Edge{from, to, weight})
}

type Edge struct {
	From   *Node
	To     *Node
	Weight int64
}

func TraverseGraph(start *Node, visited map[string]bool) string {
	if visited[start.ID] {
		return ""
	}
	visited[start.ID] = true

	var result strings.Builder
	for _, edge := range start.OutEdges {
		result.WriteString(fmt.Sprintf("\"%s\" -> \"%s\" [label=\"%d\"];\n", edge.From.ID, edge.To.ID, edge.Weight))
		result.WriteString(TraverseGraph(edge.To, visited))
	}
	return result.String()
}

func GenerateGraphViz(root *Node) string {
	visited := make(map[string]bool)
	var graph strings.Builder
	graph.WriteString("digraph G {\n")
	graph.WriteString(TraverseGraph(root, visited))
	graph.WriteString("}")
	return graph.String()
}

func FilterEdges(edges []Edge, memMap map[string]struct{}) []Edge {
	filteredEdges := []Edge{}
	for _, e := range edges {
		if _, ok := memMap[e.To.ID]; !ok {
			filteredEdges = append(filteredEdges, e)
		}
	}

	return filteredEdges
}

func GetLongestPath(root *Node, target *Node, edgeMap map[string][]Edge) []*Node {
	path := []*Node{}
	memMap := map[string]struct{}{}

	path = append(path, target)

	if root.ID == target.ID {
		return path
	}

outerLoop:
	for {
		if path[len(path)-1] == root {
			break outerLoop
		}

		for _, edges := range edgeMap {
			filteredEdges := FilterEdges(edges, memMap)
			if len(filteredEdges) == 1 {
				if filteredEdges[0].To.ID == path[len(path)-1].ID {
					memMap[path[len(path)-1].ID] = struct{}{}
					path = append(path, filteredEdges[0].From)
				}
			}
		}
	}
	return path
}

func main() {
	start := time.Now()

	file, err := os.Open("./0079_keylog.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var keylog [][]int64

	nodes := map[string]*Node{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")
		row := []int64{}
		for _, value := range values {
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}

			// Add node
			if _, ok := nodes[value]; !ok {
				nodes[value] = NewNode(value)
			}

			row = append(row, num)
		}
		keylog = append(keylog, row)
	}

	for _, row := range keylog {
		node0 := nodes[fmt.Sprintf("%d", row[0])]
		node1 := nodes[fmt.Sprintf("%d", row[1])]
		node2 := nodes[fmt.Sprintf("%d", row[2])]
		node0.AddEdgeTo(node1, 1)
		node1.AddEdgeTo(node2, 1)
	}

	edgeMap := map[string][]Edge{}
	nodesWithInEdge := map[string]struct{}{}
	for _, node := range nodes {
		edgeMap[node.ID] = node.OutEdges
		for _, edge := range node.OutEdges {
			nodesWithInEdge[edge.To.ID] = struct{}{}
		}
	}

	var rootNode *Node
	var targetNode *Node
	for nodeID, node := range nodes {
		if _, ok := nodesWithInEdge[nodeID]; !ok {
			rootNode = node
		}

		if len(node.OutEdges) == 0 {
			targetNode = node
		}
	}

	path := GetLongestPath(rootNode, targetNode, edgeMap)
	slices.Reverse(path)

	var passwordBuilder strings.Builder
	for _, node := range path {
		passwordBuilder.WriteString(node.ID)
	}

	elapsed := time.Since(start)
	fmt.Printf("password = %v (elapsed = %v)\n", passwordBuilder.String(), elapsed)

	// Create graphviz dot file
	dotFlag := flag.Bool("dot", false, "Create graphviz dot file")
	flag.Parse()

	if *dotFlag {
		dot := GenerateGraphViz(rootNode)
		file, err = os.Create("graph.dot")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(dot)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		writer.Flush()

	}
}
