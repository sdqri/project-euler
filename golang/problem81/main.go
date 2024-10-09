package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	node     *Node
	prevNode *Node
	dist     int64
}

type PriorityQueue []Item

var _ sort.Interface = PriorityQueue{}

func (pq *PriorityQueue) Push(item Item) {
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() Item {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func NewNode(id string) *Node {
	return &Node{
		ID:       id,
		OutEdges: make([]Edge, 0),
	}
}

type Node struct {
	ID       string
	OutEdges []Edge
}

func (from *Node) AddEdgeTo(to *Node, weight int64) {
	for _, edge := range from.OutEdges {
		if edge.To == to {
			return
		}
	}
	from.OutEdges = append(from.OutEdges, Edge{from, to, weight})
	// to.OutEdges = append(to.OutEdges, Edge{to, from, weight})
}

type Edge struct {
	From   *Node
	To     *Node
	Weight int64
}

func CreatePriorityQueue(pathMap map[string]Item, seenNodes map[string]struct{}) PriorityQueue {
	pq := PriorityQueue{}
	for _, item := range pathMap {
		if _, ok := seenNodes[item.node.ID]; !ok {
			pq.Push(item)
		}
	}

	sort.Sort(pq)
	slices.Reverse(pq)
	return pq
}

func Dijkstra(start *Node) map[string]Item {
	pathMap := map[string]Item{}
	seenNodes := map[string]struct{}{}

	newItem := Item{
		node:     start,
		prevNode: nil,
		dist:     0,
	}
	pathMap[start.ID] = newItem

	frontier := CreatePriorityQueue(pathMap, seenNodes)

	for len(frontier) != 0 {
		currentItem := frontier.Pop()
		for _, edge := range currentItem.node.OutEdges {
			nodeDist, seen := pathMap[edge.To.ID]
			if seen {
				currentDist := currentItem.dist + edge.Weight
				if currentDist < nodeDist.dist {
					newItem := Item{
						node:     edge.To,
						prevNode: currentItem.node,
						dist:     currentDist,
					}
					pathMap[newItem.node.ID] = newItem
				}
			} else {
				newItem := Item{
					node:     edge.To,
					prevNode: currentItem.node,
					dist:     currentItem.dist + int64(edge.Weight),
				}
				pathMap[newItem.node.ID] = newItem
			}
		}

		seenNodes[currentItem.node.ID] = struct{}{}
		frontier = CreatePriorityQueue(pathMap, seenNodes)
	}

	return pathMap
}

func main() {
	file, err := os.Open("./0081_matrix.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	start := time.Now()

	scanner := bufio.NewScanner(file)

	matrix := [][]int64{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		row := []int64{}
		for _, value := range values {
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}

		matrix = append(matrix, row)
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
		}
	}

	// Add nodeZero
	startNode := NewNode("startNode")
	firstNode := nodes["0-0"]
	startNode.AddEdgeTo(firstNode, matrix[0][0])

	target := fmt.Sprintf("%d-%d", len(matrix)-1, len(matrix[0])-1)
	pathMat := Dijkstra(startNode)
	targetCost := pathMat[target].dist
	elapsed := time.Since(start)
	fmt.Printf("minimal path sum from the top left to the bottom right\n")
	fmt.Printf("by only moving right and down = %v (elapsed = %v)\n", targetCost, elapsed)
}

func GetShortestPath(pathMap map[string]Item, target string) []*Node {
	targetItem := pathMap[target]
	path := []*Node{}
	for {
		path = append(path, targetItem.node)
		if targetItem.prevNode == nil {
			break
		}
		targetItem = pathMap[targetItem.prevNode.ID]
	}

	slices.Reverse(path)

	return path
}
