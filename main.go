package main

import (
	"context"
	"flag"
	"sync/atomic"

	"fortio.org/terminal/ansipixels"
	"fortio.org/terminal/ansipixels/tcolor"
	"github.com/geofpwhite/snaketerm/snake"
)

func main() {
	fps := flag.Float64("fps", 60, "set fps")
	flag.Parse()
	ap := ansipixels.NewAnsiPixels(*fps)
	ap.Open()
	ap.HideCursor()
	ap.ClearScreen()
	s := snake.NewSnake(ap.W, ap.H)
	var quit atomic.Bool
	quit.Store(false)
	go func() {
		for {
			curDir := atomic.LoadInt32(s.Dir)
			if len(ap.Data) >= 3 {
				switch ap.Data[2] {
				case 65:
					if curDir == snake.R || curDir == snake.L && s.CanChangeDir.Load() {
						atomic.StoreInt32(s.Dir, snake.U)
						s.CanChangeDir.Store(false)
					}
				case 66:
					if curDir == snake.R || curDir == snake.L && s.CanChangeDir.Load() {
						atomic.StoreInt32(s.Dir, snake.D)
						s.CanChangeDir.Store(false)
					}
				case 67:
					if curDir == snake.U || curDir == snake.D && s.CanChangeDir.Load() {
						atomic.StoreInt32(s.Dir, snake.R)
						s.CanChangeDir.Store(false)
					}
				case 68:
					if curDir == snake.U || curDir == snake.D && s.CanChangeDir.Load() {
						atomic.StoreInt32(s.Dir, snake.L)
						s.CanChangeDir.Store(false)
					}
				}
			}
			if len(ap.Data) > 0 && ap.Data[0] == 'q' {
				quit.Store(true)
			}
		}
	}()

	ap.FPSTicks(context.Background(), func(ctx context.Context) bool {
		if !s.Next() {
			return false
		}
		ap.ClearScreen()
		ap.StartSyncMode()
		mouthCoords := s.Snake[len(s.Snake)-1]
		ap.WriteAt(mouthCoords.X, mouthCoords.Y, "%s ", ap.ColorOutput.Background(tcolor.Red.Color()))
		foodCoords := s.Food
		ap.WriteAt(foodCoords.X, foodCoords.Y, "%s ", ap.ColorOutput.Background(tcolor.Green.Color()))
		for _, coords := range s.Snake[:len(s.Snake)-1] {
			ap.WriteAt(coords.X, coords.Y, "%s ", ap.ColorOutput.Background(tcolor.White.Color()))
		}
		ap.WriteBg(tcolor.Black.Color())

		ap.EndSyncMode()
		return !quit.Load()
	})
}
