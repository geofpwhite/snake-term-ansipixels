package main

import (
	"flag"

	"fortio.org/terminal/ansipixels"
	"fortio.org/terminal/ansipixels/tcolor"
	"github.com/geofpwhite/snake-term-ansipixels/snake"
)

func main() {
	fps := flag.Float64("fps", 60, "set fps")
	flag.Parse()
	ap := ansipixels.NewAnsiPixels(*fps)
	ap.Open()
	defer func() {
		ap.Restore()
		ap.ShowCursor()
		ap.MoveCursor(0, 0)
	}()
	ap.HideCursor()
	ap.ClearScreen()
	s := snake.NewSnake(ap.W, ap.H)
	ap.FPSTicks(func() bool {
		ap.ClearScreen()
		ap.WriteAt(0, 0, "%v", ap.Data)
		if len(ap.Data) >= 3 {
			switch ap.Data[2] {
			case 65:
				if s.Dir == snake.R || s.Dir == snake.L {
					s.Dir = snake.U
				}
			case 66:
				if s.Dir == snake.R || s.Dir == snake.L {
					s.Dir = snake.D
				}
			case 67:
				if s.Dir == snake.U || s.Dir == snake.D {
					s.Dir = snake.R
				}
			case 68:
				if s.Dir == snake.U || s.Dir == snake.D {
					s.Dir = snake.L
				}
			}
		}
		if len(ap.Data) > 0 && ap.Data[0] == 'q' {
			return false
		}
		if !s.Next() {
			return false
		}
		mouthCoords := s.Snake[len(s.Snake)-1]
		ap.WriteAt(mouthCoords.X, mouthCoords.Y, "%s ", ap.ColorOutput.Background(tcolor.Red.Color()))
		foodCoords := s.Food
		ap.WriteAt(foodCoords.X, foodCoords.Y, "%s ", ap.ColorOutput.Background(tcolor.Green.Color()))
		for _, coords := range s.Snake[:len(s.Snake)-1] {
			ap.WriteAt(coords.X, coords.Y, "%s ", ap.ColorOutput.Background(tcolor.White.Color()))
		}
		ap.WriteString(tcolor.Reset)

		return true
	})
}
