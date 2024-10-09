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

			// Add left edge
			leftID := fmt.Sprintf("%d-%d", i, j-1)
			leftNode, ok := nodes[leftID]
			if ok {
				node.AddEdgeTo(leftNode, matrix[i][j-1])
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
	startID := fmt.Sprintf("%d-%d", 0, 0)
	toNode := nodes[startID]
	startNode.AddEdgeTo(toNode, matrix[0][0])

	endNode := NewNode("endNode")
	endID := fmt.Sprintf("%d-%d", len(matrix)-1, len(matrix[0])-1)
	fromNode := nodes[endID]
	fromNode.AddEdgeTo(endNode, 0)

	pathMat := Dijkstra(startNode)
	targetCost := pathMat["endNode"].dist
	assert.Equal(t, int64(2297), targetCost)
}
