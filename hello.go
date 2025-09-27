package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Height   int
	Width    int
	tileSize int
	State    [][]int
}

func NewGame(width, height, tileSize int) *Game {
	g := &Game{Width: width, Height: height, tileSize: tileSize}
	g.State = CreateGameState(g.Width/g.tileSize, g.Height/g.tileSize)

	return g
}

func (g *Game) Draw() {
	for y := range g.State {
		for x := 0; x < len(g.State[y]); x++ {
			if g.State[y][x] == 1 {

				pixelX := x * g.tileSize
				pixelY := y * g.tileSize

				rl.DrawRectangle(int32(pixelX), int32(pixelY), int32(g.tileSize), int32(g.tileSize), rl.RayWhite)
			}
		}
	}
}

func (g *Game) Update() {
	newState := CreateGameState(len(g.State[0]), len(g.State))

	for indexY, cellY := range g.State {
		for indexX, cellX := range cellY {
			neighbours := CountNeighbours(indexX, indexY, g.State)
			newState[indexY][indexX] = IsCellAlive(cellX, neighbours)
		}
	}

	g.State = newState
}

func CountNeighbours(x, y int, gameState [][]int) int {
	count := 0

	for cellX := x - 1; cellX <= x+1; cellX++ {
		for cellY := y - 1; cellY <= y+1; cellY++ {
			if cellY < 0 || cellX < 0 || cellY >= len(gameState) || cellX >= len(gameState[0]) {
				continue
			}

			if cellY == y && cellX == x {
				continue
			}

			if gameState[cellY][cellX] == 1 {
				count++
			}
		}
	}

	return count
}

func IsCellAlive(current, neighbours int) int {
	switch {
	case neighbours < 2:
		return 0
	case neighbours == 2:
		return current
	case neighbours == 3:
		return 1
	case neighbours > 3:
		return 0
	}

	return 0
}

func CreateGameState(newWidth, newHeight int) [][]int {
	newState := make([][]int, newHeight)

	for i := range newHeight {
		newState[i] = make([]int, newWidth)
	}

	return newState
}

func CreateGliders(x, y int, gameState *[][]int) {
	(*gameState)[y][x+1] = 1
	(*gameState)[y+1][x+2] = 1
	(*gameState)[y+2][x] = 1
	(*gameState)[y+2][x+1] = 1
	(*gameState)[y+2][x+2] = 1
}

func addPattern(x, y int, pattern [][]int, gameState *[][]int) {
	for row := range pattern {
		for col := 0; col < len(pattern[row]); col++ {
			if pattern[row][col] == 1 {
				(*gameState)[y+row][x+col] = 1
			}
		}
	}
}

func CreateGliderGun(x, y int, gameState *[][]int) {
	pattern := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	addPattern(x, y, pattern, gameState)
}

func CreatePulsar(x, y int, gameState *[][]int) {
	pattern := [][]int{
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
	}

	addPattern(x, y, pattern, gameState)
}

func CreatePentadecathlon(x, y int, gameState *[][]int) {
	// Create a slice of the pattern
	pattern := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
	}

	addPattern(x, y, pattern, gameState)
}

func main() {
	var game = NewGame(800, 600, 10)

	CreateGliderGun(0, 0, &game.State)
	CreatePentadecathlon(40, 10, &game.State)
	CreatePulsar(60, 20, &game.State)

	rl.InitWindow(int32(game.Width), int32(game.Height), "Game of life")

	defer rl.CloseWindow()

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		game.Update()
		game.Draw()

		rl.EndDrawing()
	}
}
