package main

import (
	"github.com/banthar/Go-SDL/sdl"
)

const (
	SIZE_CELL = 20
	W         = NB_COL * SIZE_CELL
	H         = NB_ROW * SIZE_CELL
)

func main() {

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}
	var screen = sdl.SetVideoMode(W, H, 32, sdl.RESIZABLE)

	grid := genGrid()

	sdl.EnableUNICODE(1)

	sdl.WM_SetCaption("Labyrinthe", "")
	running := true
	if sdl.GetKeyName(270) != "[+]" {
		panic("GetKeyName broken")
	}

	player_x := 1
	player_y := 1
	player := &sdl.Rect{int16(player_x*SIZE_CELL + 2), int16(player_y*SIZE_CELL + 2), SIZE_CELL - 2, SIZE_CELL - 2}

	// Bool pour savoir si on genere au fur et a mesure ou pas
	instant_draw := true

	screen.FillRect(&sdl.Rect{0, 0, W, H}, sdl.MapRGB(screen.Format, 255, 255, 255))
	screen.Flip()

	for running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch e := ev.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			case *sdl.KeyboardEvent:
				//println(e.Keysym.Sym)
				if e.Keysym.Sym == 27 {
					running = false
				}
				if e.Type == sdl.KEYDOWN {
					if e.Keysym.Sym == 32 {
						grid = genGrid()
						player.X = SIZE_CELL + 2
						player.Y = SIZE_CELL + 2
					} else if e.Keysym.Sym == 274 {
						pos_y := player.Y + 1*SIZE_CELL
						pos_row := pos_y / SIZE_CELL
						pos_col := player.X / SIZE_CELL
						if !grid[pos_row][pos_col].nWall && pos_row < NB_ROW-1 {
							player.Y += 1 * SIZE_CELL
						}
					} else if e.Keysym.Sym == 273 {
						pos_y := player.Y - 1*SIZE_CELL
						pos_row := pos_y / SIZE_CELL
						pos_col := player.X / SIZE_CELL
						if !grid[pos_row][pos_col].sWall && pos_row > 0 {
							player.Y -= 1 * SIZE_CELL
						}
					} else if e.Keysym.Sym == 275 {
						pos_x := player.X + 1*SIZE_CELL
						pos_row := player.Y / SIZE_CELL
						pos_col := pos_x / SIZE_CELL
						if !grid[pos_row][pos_col].oWall && pos_col < NB_COL-1 {
							player.X += 1 * SIZE_CELL
						}
					} else if e.Keysym.Sym == 276 {
						pos_x := player.X - 1*SIZE_CELL
						pos_row := player.Y / SIZE_CELL
						pos_col := pos_x / SIZE_CELL
						if !grid[pos_row][pos_col].eWall && pos_col > 0 {
							player.X -= 1 * SIZE_CELL
						}
					}
					//println(player.Y/SIZE_CELL, player.X/SIZE_CELL)
					screen.FillRect(&sdl.Rect{0, 0, W, H}, sdl.MapRGB(screen.Format, 255, 255, 255))
					if player.Y/SIZE_CELL == NB_ROW-2 && player.X/SIZE_CELL == NB_COL-2 {
						grid = genGrid()
						player.X = SIZE_CELL + 2
						player.Y = SIZE_CELL + 2
					}
				}
			}
		}
		// Draw ici
		if instant_draw {
			if visitedCells < totalCells {
				grid = next(grid)
			}
		} else {
			for visitedCells < totalCells {
				grid = next(grid)
			}
		}
		drawLaby(screen, grid)
		screen.FillRect(player, sdl.MapRGB(screen.Format, 0, 255, 0))
		screen.Flip()
	}
}

func drawLaby(screen *sdl.Surface, grid [NB_ROW][NB_COL]Cell) {
	for r := 1; r < NB_ROW-1; r++ {
		for c := 1; c < NB_COL-1; c++ {
			cell := grid[r][c]
			pos_row := r * SIZE_CELL
			pos_col := c * SIZE_CELL
			if cell.nWall {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 0, 0))
			} else {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 255, 255))
			}
			if cell.sWall {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row + SIZE_CELL), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 0, 0))
			} else {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row + SIZE_CELL), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 255, 255))
			}
			if cell.oWall {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row), 2, SIZE_CELL}, sdl.MapRGB(screen.Format, 255, 0, 0))
			} else {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row), 2, SIZE_CELL}, sdl.MapRGB(screen.Format, 255, 255, 255))
			}
			if cell.eWall {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row + SIZE_CELL), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 0, 0))
			} else {
				screen.FillRect(&sdl.Rect{int16(pos_col), int16(pos_row + SIZE_CELL), SIZE_CELL, 2}, sdl.MapRGB(screen.Format, 255, 255, 255))
			}
		}
	}
	//// Bordure tout a droite
	screen.FillRect(&sdl.Rect{int16(W - SIZE_CELL), SIZE_CELL, 2, H - SIZE_CELL*2}, sdl.MapRGB(screen.Format, 255, 0, 0))
	//// Bourdure tout en haut
	screen.FillRect(&sdl.Rect{SIZE_CELL, SIZE_CELL, W - SIZE_CELL*2, 2}, sdl.MapRGB(screen.Format, 255, 0, 0))
	//// Bordure tout en bas
	screen.FillRect(&sdl.Rect{SIZE_CELL, H - SIZE_CELL, W - SIZE_CELL*2 + 2, 2}, sdl.MapRGB(screen.Format, 255, 0, 0))
}
