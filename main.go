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

func (s *Session) Update() error {
	keyLog(s)
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
	if err := ebiten.RunGame(s); err != nil {
		fmt.Println("err")
	}
}

func keyLog(s *Session) {
	currTime := time.Since(s.start)
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Printf("W pressed at %v\n", currTime)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		fmt.Printf("A pressed at %v\n", currTime)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Printf("S pressed at %v\n", currTime)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Printf("D pressed at %v\n", currTime)
	}
}
