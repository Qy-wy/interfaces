package main

import "fmt"

type sport interface {
	play() string
	equipment() []string
}

type football struct {
	name string
}

type valleyball struct {
	name string
}

func (f football) play() string {
	return "You need to kick the ball into the opponent's goal"
}

func (v valleyball) play() string {
	return "You need to make sure that the ball falls into the opponent's territory"
}

func (f football) equipment() []string {
	return []string{"ball", "gates", "boots"}
}

func (v valleyball) equipment() []string {
	return []string{"ball", "valleyball-net"}
}

func input(s sport) {
	fmt.Println(s)
	fmt.Println("The point of the game:", s.play())
	fmt.Println("Equipment to play:", s.equipment())
	fmt.Println()
}

func main() {
	f := football{"football"}
	v := valleyball{"valleyball"}

	input(f)
	input(v)
}
