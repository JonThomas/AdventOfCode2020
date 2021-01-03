package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
)

// cube[w][z][y][x]
// 1 = active
// 0 = inactive
type cube [][][][]int

var endAtCycle int = 6
var cubeSize int

func main() {

	file, err := files.ReadFile("./Day17Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	theCube, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	printCube(0, theCube)

	theCube = moveToCycle(1, theCube)
	//printCube(1, theCube)
	theCube = moveToCycle(2, theCube)
	theCube = moveToCycle(3, theCube)
	theCube = moveToCycle(4, theCube)
	theCube = moveToCycle(5, theCube)
	theCube = moveToCycle(6, theCube)

	//printCube(6, theCube)

	answer := countActive(theCube)

	fmt.Printf("END. Answer is: %d.\n", answer)
}

func countActive(theCube cube) int {
	answer := 0
	for w := 0; w < cubeSize; w++ {
		for z := 0; z < cubeSize; z++ {
			for y := 0; y < cubeSize; y++ {
				for x := 0; x < cubeSize; x++ {
					answer += theCube[w][z][y][x]
				}
			}
		}
	}
	return answer
}

func moveToCycle(cycle int, theCube cube) cube {

	newCube := initializeCube(cubeSize)

	start := endAtCycle - cycle
	end := cubeSize - endAtCycle + cycle

	fmt.Printf("Calculating cycle %d. Top left = (%d,%d,%d)\n", cycle, start, start, start)

	for w := start; w < end; w++ {
		for z := start; z < end; z++ {
			for y := start; y < end; y++ {
				for x := start; x < end; x++ {
					activeNeighbors := countAciveNeighbors(x, y, z, w, theCube)
					if theCube[w][z][y][x] == 1 { // Active cube
						if activeNeighbors == 2 || activeNeighbors == 3 {
							newCube[w][z][y][x] = 1
						} else {
							newCube[w][z][y][x] = 0
						}
					} else {
						// Inactive cube
						if activeNeighbors == 3 {
							newCube[w][z][y][x] = 1
						} else {
							newCube[w][z][y][x] = 0
						}
					}
				}
			}
		}
	}
	return newCube
}

func countAciveNeighbors(x int, y int, z int, w int, theCube cube) int {
	activeNeighbors := 0
	for newW := w - 1; newW <= w+1; newW++ {
		for newX := x - 1; newX <= x+1; newX++ {
			for newY := y - 1; newY <= y+1; newY++ {
				for newZ := z - 1; newZ <= z+1; newZ++ {
					if newX == x && newY == y && newZ == z && newW == w {
						continue
					}
					activeNeighbors += checkIfActive(newX, newY, newZ, newW, theCube)
				}
			}
		}
	}
	return activeNeighbors
}

func checkIfActive(x int, y int, z int, w int, theCube cube) int {

	if x < 0 || x >= cubeSize ||
		y < 0 || y >= cubeSize ||
		z < 0 || z >= cubeSize ||
		w < 0 || w >= cubeSize {
		return 0
	}
	return theCube[w][z][y][x]
}

func printCube(cycle int, theCube cube) {
	fmt.Printf("\nAfter %d cycles:\n", cycle)

	start := endAtCycle - cycle
	end := len(theCube) - endAtCycle + cycle

	zEnd := 0
	if cycle == 0 {
		zEnd = 1 + start
	} else {
		zEnd = cubeSize - start - 1
	}

	for w := start; w < zEnd; w++ {
		for z := start; z < zEnd; z++ {
			fmt.Printf("\nz=%d, w=%d. Top left = (%d,%d,%d)\n", z-start-cycle, w-start-cycle, start, start, start)

			for y := start + 1; y < end; y++ {
				for x := start; x < end; x++ {
					if theCube[w][z][y][x] == 1 {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}

				}
				fmt.Println()
			}
		}
	}
}

func parseInput(fileLines []string) (cube, error) {

	fileWidth := len(fileLines[0])
	cubeSize = fileWidth + 2*endAtCycle

	thisCube := initializeCube(cubeSize)

	w := endAtCycle
	z := endAtCycle
	y := endAtCycle
	x := endAtCycle

	for _, line := range fileLines {

		for _, char := range line {
			if char == '#' {
				thisCube[w][z][y][x] = 1
			} else {
				thisCube[w][z][y][x] = 0
			}
			x++
		}
		y++
		x = endAtCycle
	}

	return thisCube, nil
}

func initializeCube(cubeSize int) cube {

	var thisCube = make(cube, cubeSize)

	// initialize cube with 4 dimesions of size (cubeSize)
	for i := range thisCube {
		thisCube[i] = make([][][]int, cubeSize)
		for j := range thisCube[i] {
			thisCube[i][j] = make([][]int, cubeSize)
			for k := range thisCube[i][j] {
				thisCube[i][j][k] = make([]int, cubeSize)
			}
		}
	}
	return thisCube
}
