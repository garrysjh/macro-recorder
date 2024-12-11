package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Session struct{}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Update() error {
	keyLog(s)
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
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		fmt.Println("W")
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		fmt.Println("A")
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		fmt.Println("S")
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		fmt.Println("D")
	}
}
