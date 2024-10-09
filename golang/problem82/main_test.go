package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	matrix := [][]int64{
		{131, 673, 234, 103, 18},
		{201, 96, 342, 965, 150},
		{630, 803, 746, 422, 111},
		{537, 699, 497, 121, 956},
		{805, 732, 524, 37, 331},
	}

	// Create nodes:
	nodes := map[string]*Node{}
	for i, row := range matrix {
		for j := range row {
			id := fmt.Sprintf("%d-%d", i, j)
			node := NewNode(id)
			nodes[id] = node
		}
	}

	// Create edges:
	for i, row := range matrix {
		for j := range row {
			id := fmt.Sprintf("%d-%d", i, j)
			node := nodes[id]
			// Add right edge
			rightID := fmt.Sprintf("%d-%d", i, j+1)
			rightNode, ok := nodes[rightID]
			if ok {
				node.AddEdgeTo(rightNode, matrix[i][j+1])
			}

			// Add down edge
			downID := fmt.Sprintf("%d-%d", i+1, j)
			downNode, ok := nodes[downID]
			if ok {
				node.AddEdgeTo(downNode, matrix[i+1][j])
			}

			// Add up edge
			upID := fmt.Sprintf("%d-%d", i-1, j)
			upNode, ok := nodes[upID]
			if ok {
				node.AddEdgeTo(upNode, matrix[i-1][j])
			}
		}
	}

	startNode := NewNode("startNode")
	for i := range matrix {
		id := fmt.Sprintf("%d-%d", i, 0)
		toNode := nodes[id]
		startNode.AddEdgeTo(toNode, matrix[i][0])
	}

	endNode := NewNode("endNode")
	for i := range matrix {
		id := fmt.Sprintf("%d-%d", i, len(matrix[0])-1)
		fromNode := nodes[id]
		fromNode.AddEdgeTo(endNode, 0)
	}

	pathMat := Dijkstra(startNode)
	targetCost := pathMat["endNode"].dist
	assert.Equal(t, int64(994), targetCost)
}
