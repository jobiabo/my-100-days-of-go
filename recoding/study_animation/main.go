package main

import (
	"fmt"

	"animation/animation"
	"time"
)

func main() {
	anim := animation.NewAnimation("TEST", 4)

	anim.GenerateSpinFrames()

	fmt.Println(anim.GetFrame(0))

	fmt.Println("------------")

	fmt.Println(anim.Play())

	time.Sleep(500 * time.Millisecond)
}