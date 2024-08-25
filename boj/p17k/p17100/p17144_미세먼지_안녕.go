package p17100

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Solve17144(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	r, _ := strconv.Atoi(params[0])
	c, _ := strconv.Atoi(params[1])
	t, _ := strconv.Atoi(params[2])

	currentRoom := make([][]int, r)
	for i := 0; i < r; i++ {
		scanner.Scan()
		currentRoom[i] = parseLine(scanner.Text(), c)
	}

	var purifierTop, purifierBottom int
	for y := 0; y < r; y++ {
		if currentRoom[y][0] == -1 {
			purifierTop = y
			purifierBottom = y + 1
			break
		}
	}

	for i := 0; i < t; i++ {
		currentRoom = spread(currentRoom)
		purify(currentRoom, purifierTop, purifierBottom)
	}

	ans := 0
	for _, row := range currentRoom {
		for _, cell := range row {
			if cell > 0 {
				ans += cell
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func parseLine(line string, c int) []int {
	fields := strings.Fields(line)
	result := make([]int, c)
	for i := 0; i < c; i++ {
		result[i], _ = strconv.Atoi(fields[i])
	}
	return result
}

type MoveParams struct {
	y, x           int
	dy, dx         int
	limitY, limitX int
}

func movePurifier(room [][]int, params MoveParams, prev int) int {
	y, x := params.y, params.x

	for (y != params.limitY || params.dy == 0) && (x != params.limitX || params.dx == 0) {
		room[y][x], prev = prev, room[y][x]

		if ny := y + params.dy; ny >= 0 && ny < len(room) {
			y = ny
		} else {
			break
		}

		if nx := x + params.dx; nx >= 0 && nx < len(room[0]) {
			x = nx
		} else {
			break
		}
	}
	return prev
}

func executePurifierMovements(room [][]int, movements []MoveParams) {
	prev := 0
	for _, params := range movements {
		prev = movePurifier(room, params, prev)
	}
}

func purify(room [][]int, purifierTop int, purifierBottom int) {
	c := len(room[0])
	r := len(room)

	topMovements := []MoveParams{
		{purifierTop, 1, 0, 1, purifierTop, c - 1},
		{purifierTop, c - 1, -1, 0, 0, c - 1},
		{0, c - 1, 0, -1, 0, 0},
		{0, 0, 1, 0, purifierTop, 0},
	}

	bottomMovements := []MoveParams{
		{purifierBottom, 1, 0, 1, purifierBottom, c - 1},
		{purifierBottom, c - 1, 1, 0, r - 1, c - 1},
		{r - 1, c - 1, 0, -1, r - 1, 0},
		{r - 1, 0, -1, 0, purifierBottom, 0},
	}

	executePurifierMovements(room, topMovements)
	executePurifierMovements(room, bottomMovements)
}

func spreadDust(y int, x int, room [][]int, newRoom [][]int) {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	amount := room[y][x] / 5
	spreadCount := 0
	r := len(room)
	c := len(room[0])

	for _, dir := range directions {
		ny, nx := y+dir[0], x+dir[1]
		if ny >= 0 && ny < r && nx >= 0 && nx < c && room[ny][nx] != -1 {
			newRoom[ny][nx] += amount
			spreadCount++
		}
	}

	newRoom[y][x] -= amount * spreadCount
}

func spread(room [][]int) [][]int {
	newRoom := make([][]int, len(room))
	for i := range room {
		newRoom[i] = make([]int, len(room[i]))
		copy(newRoom[i], room[i])
	}

	for y := range room {
		for x := range room[y] {
			if room[y][x] > 0 {
				spreadDust(y, x, room, newRoom)
			}
		}
	}
	return newRoom
}
