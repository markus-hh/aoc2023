package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/aoc/util"
)

type Node struct {
	label string
	leftNode *Node
	rightNode *Node
}

// L -> false
// R -> true
func parseDirection(directionRune rune) bool {
	if directionRune == 'L' {
		return false
	} else if(directionRune == 'R') {
		return true
	} else {
		fmt.Println("error during directions parsing")
		os.Exit(1)
		return false
	}
}

func parseDirections(line string) (directions []bool) {
	runes := []rune(line)
	return util.MapFunc(runes, parseDirection)
}

func parseLabel(line string) string {
	if !strings.Contains(line, " = ") {
		fmt.Println("error during label parsing")
		os.Exit(1)
		return ""
	}
	return strings.Split(line, " = ")[0]
}

func parseLine(line string) (label string, left string, right string) {
	label = parseLabel(line)
	tupleRaw := strings.Split(line, " = ")[1]
	tupleRaw = strings.Trim(tupleRaw, "()")
	tupleParts := strings.Split(tupleRaw, ", ")
	return label, tupleParts[0], tupleParts[1]
}

func parseInput(lines []string) map[string]*Node {
	labels := util.MapFunc(lines, parseLabel)
	nodeMap := map[string]*Node {}

	for _, label := range labels {
		nodeMap[label] = &Node { label: label }
	}

	for _, line := range lines {
		label, leftNodeLabel, rightNodeLabel := parseLine(line)
		currentNode := nodeMap[label]
		leftNode := nodeMap[leftNodeLabel]
		currentNode.leftNode = leftNode

		rightNode := nodeMap[rightNodeLabel]
		currentNode.rightNode = rightNode

		nodeMap[label] = currentNode
	}

	return nodeMap
}

func findSmallestPathToTarget(nodes []*Node, directions []bool, startingNode *Node, targetLabel string) int {
	nodeAmount := len(nodes)
	pathCache := map[string][]bool {}

	currentlyVisitedNodes := []*Node { startingNode }

	for stepCount := 0; true; stepCount++ {
		currentlyVisitedNodesNextStep := []*Node {}

		directionIndex := stepCount % len(directions)
		goRight := directions[directionIndex]

		for _, currentlyVisitedNode := range currentlyVisitedNodes {
			var nextNode Node
			if goRight {
				nextNode = *currentlyVisitedNode.rightNode
			} else {
				nextNode = *currentlyVisitedNode.leftNode
			}

			if nextNode.label == targetLabel {
				return stepCount + 1
			}

			_, containsLabel := pathCache[nextNode.label]
			if !containsLabel {
				pathCache[nextNode.label] = make([]bool, nodeAmount)
			}

			if pathCache[nextNode.label][directionIndex] {
				continue
			}

			pathCache[nextNode.label][directionIndex] = true
			currentlyVisitedNodesNextStep = append(currentlyVisitedNodesNextStep, &nextNode)
		}

		currentlyVisitedNodes = currentlyVisitedNodesNextStep
	}

	return -1
}

func solvePart1(input string) {
	lines := util.SplitLines(input)

	directions := parseDirections(lines[0])
	nodes := parseInput(lines[2:])
	startingNode := nodes["AAA"]

	nodeSlice := []*Node {}
	for _, node := range nodes {
		nodeSlice = append(nodeSlice, node)
	}
	
	solution := findSmallestPathToTarget(nodeSlice, directions, startingNode, "ZZZ")
	fmt.Println(solution)
}

func main() {
	input := util.FetchInput(8)
	solvePart1(input)
}