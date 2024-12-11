package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Session struct{}

func NewSession() *Session {
	return &Session{}
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
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Println("W")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		fmt.Println("A")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Println("S")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Println("D")
	}
}
