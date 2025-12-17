package day8

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	X int
	Y int
	Z int
}

type Edge struct {
	From Node
	To   Node
}

func (e Edge) GetLength() float64 {
	return math.Sqrt(math.Pow(float64(e.From.X-e.To.X), 2) + math.Pow(float64(e.From.Y-e.To.Y), 2) + math.Pow(float64(e.From.Z-e.To.Z), 2))
}

type Circuit struct {
	nodes map[Node]bool
}

func (c Circuit) Contains(node Node) bool {
	_, exists := c.nodes[node]
	return exists
}

func (c *Circuit) Add(node Node) {
	c.nodes[node] = true
}

func (c *Circuit) Merge(other Circuit) {
	for node := range other.nodes {
		c.Add(node)
	}
}

func (c Circuit) Size() int {
	return len(c.nodes)
}

func createCircuitWithNodes(node1, node2 Node) Circuit {
	circuit := Circuit{
		nodes: make(map[Node]bool),
	}
	circuit.Add(node1)
	circuit.Add(node2)
	return circuit
}

func AssembleCircuitNetwork(input []string, edgeLoopFunc func(edges []Edge, circuits *[]Circuit, loopFunc func(Edge))) ([]Circuit, error) {
	numberOfNodes := len(input)
	nodes := make([]Node, 0, numberOfNodes)
	for _, line := range input {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			return nil, err
		}
		node := Node{
			X: x,
			Y: y,
			Z: z,
		}
		nodes = append(nodes, node)
	}
	edges := make([]Edge, 0, numberOfNodes*(numberOfNodes-1)/2)
	for i := 0; i < numberOfNodes-1; i++ {
		for j := i + 1; j < numberOfNodes; j++ {
			edge := Edge{
				From: nodes[i],
				To:   nodes[j],
			}
			edges = append(edges, edge)
		}
	}
	slices.SortFunc(edges, func(e1 Edge, e2 Edge) int {
		return int(e1.GetLength() - e2.GetLength())
	})
	circuits := make([]Circuit, 0, len(nodes))
	for _, node := range nodes {
		newCircuit := createCircuitWithNodes(node, node)
		circuits = append(circuits, newCircuit)
	}
	edgeLoopFunc(edges, &circuits, func(edge Edge) {
		foundFromCircuitIndex := -1
		foundToCircuitIndex := -1
		for i, circuit := range circuits {
			if circuit.Contains(edge.From) {
				foundFromCircuitIndex = i
			}
			if circuit.Contains(edge.To) {
				foundToCircuitIndex = i
			}
			if foundFromCircuitIndex != -1 && foundToCircuitIndex != -1 {
				break
			}
		}
		if foundFromCircuitIndex == -1 && foundToCircuitIndex == -1 {
			newCircuit := createCircuitWithNodes(edge.From, edge.To)
			circuits = append(circuits, newCircuit)
		} else if foundFromCircuitIndex != -1 && foundToCircuitIndex == -1 {
			circuits[foundFromCircuitIndex].Add(edge.To)
		} else if foundFromCircuitIndex == -1 && foundToCircuitIndex != -1 {
			circuits[foundToCircuitIndex].Add(edge.From)
		} else if foundFromCircuitIndex != foundToCircuitIndex {
			circuits[foundFromCircuitIndex].Merge(circuits[foundToCircuitIndex])
			circuits = append(circuits[:foundToCircuitIndex], circuits[foundToCircuitIndex+1:]...)
		}
	})
	return circuits, nil
}
