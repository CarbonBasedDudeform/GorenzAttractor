package main

import (
	"fmt"
	"github.com/fogleman/gg"
)

type point struct {
	x float64
	y float64
	z float64
}

func step(old_p point, dt float64) point {
	sigma := 10.0
	beta := 8.0/3.0
	rho := 28.0

	dx := sigma * (old_p.y - old_p.x)
	dy := old_p.x * (rho - old_p.z) - old_p.y
	dz := old_p.x * old_p.y - beta * old_p.z
	
	new_p := point{old_p.x + dx * dt,old_p.y + dy * dt, old_p.z + dz * dt}
	
	return new_p;
}

func main() {
	fmt.Println("Printing an image.")
	width := 800
	height := 600

	dc := gg.NewContext(width, height)
	dc.SetRGBA(1, 1, 1, 1.0)

	p := point{0,0.1,0};
	t := 0.01

	dc.Translate(800.0/2.0, 50.0)
	scale := 0.1;

	for x := 0; x < 10000; x++ {
		old_p := p;
		p = step(p, t);
		dc.Push()
		dc.DrawLine(old_p.x/scale, old_p.z/scale, p.x/scale, p.z/scale);
		dc.Stroke()
		dc.Pop()
	}

	dc.SavePNG("Lorenz2022.png")
}