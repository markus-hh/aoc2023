package main

import (
	"fmt"
	"os"

	"example.com/aoc/util"
)

// direction mapping is relevant for flipping
const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

type Tile struct {
	pipeDirections []bool
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
func parseTileRune(tileRune rune) Tile {
	switch tileRune {
	case '|':
		return getTileFromDirections(NORTH, SOUTH)
	case '-':
		return getTileFromDirections(EAST, WEST)
	case 'L':
		return getTileFromDirections(NORTH, EAST)
	case 'J':
		return getTileFromDirections(NORTH, WEST)
	case '7':
		return getTileFromDirections(SOUTH, WEST)
	case 'F':
		return getTileFromDirections(SOUTH, EAST)
	case '.':
		return getTileFromDirections()
	case 'S':
		return getTileFromDirections(NORTH, EAST, SOUTH, WEST)
	default:
		fmt.Println("Error during tile parsing " + string(tileRune))
		os.Exit(1)
		return Tile{}
	}
}

func parseLine(line string) []Tile {
	tileRunes := []rune(line)
	return util.MapFunc(tileRunes, parseTileRune)
}

func parseLines(lines []string) [][]Tile {
	return util.MapFunc(lines, parseLine)
}

func getTileFromDirections(directionIndices ...int) Tile {
	tile := Tile{pipeDirections: make([]bool, 4)}
	for _, directionIndex := range directionIndices {
		tile.pipeDirections[directionIndex] = true
	}
	return tile
}

func findStartingPosition(tiles [][]Tile) (int, int) {
	for y, tileRow := range tiles {
		for x, tile := range tileRow {
			if isStartingPosition(tile) {
				return x, y
			}
		}
	}
	return -1, -1
}

func goInDirection(x int, y int, direction int) (int, int) {
	switch direction {
	case NORTH:
		return x, y - 1
	case EAST:
		return x + 1, y
	case SOUTH:
		return x, y + 1
	case WEST:
		return x - 1, y
	default:
		fmt.Println("Invalid direction")
		os.Exit(1)
		return 0, 0
	}
}

func isInBounds(tiles [][]Tile, x int, y int) bool {
	width := len(tiles[0])
	height := len(tiles)
	return x >= 0 && x < width && y >= 0 && y < height
}

func getTile(tiles [][]Tile, x int, y int) Tile {
	return tiles[y][x]
}

func flipDirection(direction int) int {
	return (direction + 2) % 4
}

func isPipeDirectionValid(tiles [][]Tile, x int, y int, direction int) bool {
	tile := getTile(tiles, x, y)
	if !tile.pipeDirections[direction] {
		return false
	}

	x, y = goInDirection(x, y, direction)
	if !isInBounds(tiles, x, y) {
		return false
	}

	neighborDirection := flipDirection(direction)
	neighborTile := getTile(tiles, x, y)
	return neighborTile.pipeDirections[neighborDirection]
}

func findFirstOtherDirection(tile Tile, currentDirection int) (int, bool) {
	for direction := range tile.pipeDirections {
		if tile.pipeDirections[direction] && direction != currentDirection {
			return direction, true
		}
	}

	return -1, false
}

func isStartingPosition(tile Tile) bool {
	isNotStartingTile := util.Any(tile.pipeDirections, func(value bool) bool { return !value })
	return !isNotStartingTile
}

func findLoopSize(tiles [][]Tile, startingX, startingY, startingDirection int) (int, bool, Tile) {
	x := startingX
	y := startingY
	direction := startingDirection

	for loopSize := 2; ; loopSize++ {
		if !isPipeDirectionValid(tiles, x, y, direction) {
			return -1, false, Tile {}
		}

		x, y = goInDirection(x, y, direction)
		newTile := getTile(tiles, x, y)
		sourceDirection := flipDirection(direction)
		direction, _ = findFirstOtherDirection(newTile, sourceDirection)

		if x == startingX && y == startingY {
			startingTile := getTileFromDirections(startingDirection, sourceDirection)
			return loopSize, true, startingTile
		}
	}
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	tiles := parseLines(lines)
	x, y := findStartingPosition(tiles)

	for direction := 0; direction < 4; direction++ {
		loopSize, loopFound, _ := findLoopSize(tiles, x, y, direction)
		if loopFound {
			maximumStepCount := loopSize / 2
			fmt.Println(maximumStepCount)
			return
		}
	}

	fmt.Println("no loop found")
}

func main() {
	input := util.FetchInput(10)
	solvePart1(input)
}
