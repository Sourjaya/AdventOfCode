package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type xCoordinate int
type yCoordinate int
type direction []int

type coordinates struct {
	x xCoordinate
	y yCoordinate
}

type dimension struct {
	xLength xCoordinate
	yLength yCoordinate
}

type obstacle struct {
	pos coordinates
}

type guard struct {
	pos coordinates
	dir direction
}

func main() {
	if len(os.Args) < 2 {
		panic("No input file given")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	obstacles := make(map[string]struct{})
	var g guard
	var dimension dimension
	positions := make(map[string]bool)
	obstacles, g, dimension, positions = prepareInput(file, obstacles, g, positions)

	count, path := countGuardPatrolPositions(obstacles, g, dimension, positions)
	fmt.Println(count)
	//fmt.Println(path)

	fmt.Println(countObstructions(obstacles, g, dimension, positions, path))
}

func createKey(x, y int) string {
	return fmt.Sprint(strconv.Itoa(x) + "," + strconv.Itoa(y))
}

func createPositionKey(x, y int, dir direction) string {
	return fmt.Sprint(createKey(x, y) + "," + strconv.Itoa(dir[0]) + "," + strconv.Itoa(dir[1]))
}

func countObstructions(obstacles map[string]struct{}, g guard, dimension dimension, positions map[string]bool, path [][]int) int {
	count := 0
	for k := 0; k < len(path); k++ {
		obstacleKey := createKey(path[k][0], path[k][1])
		if obstacleKey == createKey(int(g.pos.x), int(g.pos.y)) {
			continue // Skip guard's starting position
		}
		obstacles[obstacleKey] = struct{}{}
		if isTrue, _ := checkForLoop(obstacles, g, dimension, positions); isTrue {
			count++
		}
		delete(obstacles, obstacleKey)
		//fmt.Println("COUNT : ", count)
		//fmt.Println("\n\n--------------------------------------\n\n")
	}
	return count
}

func checkForLoop(obstacles map[string]struct{}, g guard, dimension dimension, pos map[string]bool) (bool, [][]int) {

	i := g.pos.y
	j := g.pos.x
	path := make([][]int, 0)
	positions := make(map[string]bool)

	for j < dimension.xLength && i < dimension.yLength && j >= 0 && i >= 0 && i < dimension.yLength {
		i = i + yCoordinate(g.dir[1])
		j = j + xCoordinate(g.dir[0])
		if i >= dimension.yLength || j >= dimension.xLength || i < 0 || j < 0 {
			break
		}
		if checkObstacle(obstacles, coordinates{x: xCoordinate(j), y: yCoordinate(i)}) {
			i -= yCoordinate(g.dir[1])
			j -= xCoordinate(g.dir[0])
			g.dir = changeDirection(g.dir)
			continue
		}
		g.pos.y = i
		g.pos.x = j
		key := createPositionKey(int(j), int(i), g.dir)
		if !positions[key] {
			positions[key] = true
			path = append(path, []int{int(j), int(i)})
		} else {
			return true, path // Loop detected
		}
	}
	return false, path
	//return count, path
}

func countGuardPatrolPositions(obstacles map[string]struct{}, g guard, dimension dimension, positions map[string]bool) (int, [][]int) {
	i := g.pos.y
	j := g.pos.x
	path := make([][]int, 0)
	count := 1
	for j < dimension.xLength && i < dimension.yLength && j >= 0 && i >= 0 && i < dimension.yLength {
		i = i + yCoordinate(g.dir[1])
		j = j + xCoordinate(g.dir[0])
		if i >= dimension.yLength || j >= dimension.xLength || i < 0 || j < 0 {
			break
		}
		if checkObstacle(obstacles, coordinates{x: xCoordinate(j), y: yCoordinate(i)}) {
			i -= yCoordinate(g.dir[1])
			j -= xCoordinate(g.dir[0])
			g.dir = changeDirection(g.dir)
			continue
		}
		g.pos.y = i
		g.pos.x = j
		key := createKey(int(j), int(i))
		if !positions[key] {
			count++
			positions[key] = true
			path = append(path, []int{int(j), int(i)})
		}
	}
	return count, path
}

func changeDirection(dir direction) direction {
	if dir[0] == 0 && dir[1] == -1 {
		return direction{1, 0}
	}
	if dir[0] == 0 && dir[1] == 1 {
		return direction{-1, 0}
	}
	if dir[0] == 1 && dir[1] == 0 {
		return direction{0, 1}
	}
	return direction{0, -1}
}

func checkObstacle(obstacles map[string]struct{}, pos coordinates) bool {
	key := createKey(int(pos.x), int(pos.y))
	_, ok := obstacles[key]
	return ok
}

func newObstacle(pos coordinates) obstacle {
	return obstacle{pos: pos}
}

func newGuard(pos coordinates, dir direction) guard {
	return guard{pos: pos, dir: dir}
}

func prepareInput(file *os.File, obstacles map[string]struct{}, g guard, positions map[string]bool) (map[string]struct{}, guard, dimension, map[string]bool) {
	scanner := bufio.NewScanner(file)
	i := 0
	inputString := ""

	directions := map[byte][]int{'^': {0, -1}, 'v': {0, 1}, '>': {1, 0}, '<': {-1, 0}}

	for scanner.Scan() {
		inputString = scanner.Text()
		for j := 0; j < len(inputString); j++ {
			if inputString[j] == '#' {
				obstacles[createKey(j, i)] = struct{}{}
			}
			positions[createKey(j, i)] = false
			if inputString[j] == '^' || inputString[j] == 'v' || inputString[j] == '>' || inputString[j] == '<' {
				g = newGuard(coordinates{x: xCoordinate(j), y: yCoordinate(i)}, directions[inputString[j]])
				positions[createKey(j, i)] = true
			}
		}
		i++
	}
	return obstacles, g, dimension{xLength: xCoordinate(len(inputString)), yLength: yCoordinate(i)}, positions
}
