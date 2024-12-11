package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Session struct {
	start time.Time
}

func NewSession() *Session {
	return &Session{
		start: time.Now(),
	}
}

func init() {
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowTitle("Input Test")
	ebiten.SetWindowSize(NewSession().Layout(1, 1))
}

func (s *Session) Update() error {
	keyLog(s)
	mouseLog(s)
	// fmt.Println(ebiten.CursorPosition())
	return nil
}

func (s *Session) Draw(screen *ebiten.Image) {
}

func (s *Session) Layout(w, h int) (int, int) {
	return 1920, 1280
}

func main() {
	s := NewSession()
	if err := ebiten.RunGameWithOptions(s, &ebiten.RunGameOptions{ScreenTransparent: true}); err != nil {
		fmt.Println("err")
	}
}

func keyLog(s *Session) {
	currTime := time.Since(s.start)
	keys := inpututil.AppendJustPressedKeys(nil)
	for _, key := range keys {
		fmt.Printf("%s pressed at %v\n", key, currTime)
	}
}

func mouseLog(s *Session) {
	currTime := time.Since(s.start)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Left Mouse pressed at %v, %v at %v\n", x, y, currTime)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse3 pressed at %v, %v at %v\n", x, y, currTime)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Right Mouse pressed at %v, %v at %v\n", x, y, currTime)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton3) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse4 pressed at %v, %v at %v\n", x, y, currTime)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton4) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse5 pressed at %v, %v at %v\n", x, y, currTime)
	}
}
