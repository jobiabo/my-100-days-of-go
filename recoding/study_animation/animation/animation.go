package animation

import (
	"fmt"
	"strings"
)

type Animation struct {
	t string
	n int
	f []string
}

func NewAnimation(t string, n int) *Animation { return &Animation{t: t, n: n} }

func (a *Animation) GenerateWaveFrames() { a.GenerateSpinFrames() }
func (a *Animation) GenerateZoomFrames() { a.GenerateSpinFrames() }

func (a *Animation) GenerateSpinFrames() {
	for i := 0; i < a.n; i++ {
		a.f = append(a.f, strings.TrimRight(strings.Repeat(fmt.Sprintf("%s%d\n", a.t, i), 10), "\n"))
	}
}



func (a *Animation) GetFrame(i int) string {
	if len(a.f) == 0 {
		return ""
	}
	return a.f[i%len(a.f)]
}

func (a *Animation) Play() string { return strings.Join(a.f, "\n\n") }