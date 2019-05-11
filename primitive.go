package primitive

import (
	"image/color"

	"github.com/tinygo-org/drivers"
)

// DrawLine draws a line between two points
func DrawLine(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, color color.RGBA) {
	if x0 == x1 {
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		for ; y0 <= y1; y0++ {
			display.SetPixel(x0, y0, color)
		}
	} else if y0 == y1 {
		if x0 > x1 {
			x0, x1 = x1, x0
		}
		for ; x0 <= x1; x0++ {
			display.SetPixel(x0, y0, color)
		}
	} else { // Bresenham
		dx := x1 - x0
		if dx < 0 {
			dx = -dx
		}
		dy := y1 - y0
		if dy < 0 {
			dy = -dy
		}
		steep := dy > dx
		if steep {
			x0, x1, y0, y1 = y0, y1, x0, x1
		}
		if x0 > x1 {
			x0, x1, y0, y1 = x1, x0, y1, y0
		}
		dx = x1 - x0
		dy = y1 - y0
		ystep := int16(1)
		if dy < 0 {
			dy = -dy
			ystep = -1
		}
		err := dx / 2
		for ; x0 <= x1; x0++ {
			if steep {
				display.SetPixel(y0, x0, color)
			} else {
				display.SetPixel(x0, y0, color)
			}
			err -= dy
			if err < 0 {
				y0 += ystep
				err += dx
			}
		}
	}
}

// DrawRectangle draws a rectangle given a point and size
func DrawRectangle(display drivers.Displayer, x int16, y int16, w int16, h int16, color color.RGBA) {
	DrawLine(display, x, y, x+w, y, color)
	DrawLine(display, x, y, x, y+h, color)
	DrawLine(display, x+w, y, x+w, y+h, color)
	DrawLine(display, x, y+h, x+w, y+h, color)
}

// DrawFilledRectangle draws a filled rectangle given a point and size
func DrawFilledRectangle(display drivers.Displayer, x int16, y int16, w int16, h int16, color color.RGBA) {
	for i := x; i <= x+w; i++ {
		DrawLine(display, i, y, i, y+h, color)
	}
}

// DrawCircle draws a circle given a point and radius
func DrawCircle(display drivers.Displayer, x0 int16, y0 int16, r int16, color color.RGBA) {
	f := 1 - r
	ddfx := int16(1)
	ddfy := -2 * r
	x := int16(0)
	y := r
	display.SetPixel(x0, y0+r, color)
	display.SetPixel(x0, y0-r, color)
	display.SetPixel(x0+r, y0, color)
	display.SetPixel(x0-r, y0, color)
	for x < y {
		if f >= 0 {
			y--
			ddfy += 2
			f += ddfy
		}
		x++
		ddfx += 2
		f += ddfx

		display.SetPixel(x0+x, y0+y, color)
		display.SetPixel(x0-x, y0+y, color)
		display.SetPixel(x0+x, y0-y, color)
		display.SetPixel(x0-x, y0-y, color)
		display.SetPixel(x0+y, y0+x, color)
		display.SetPixel(x0-y, y0+x, color)
		display.SetPixel(x0+y, y0-x, color)
		display.SetPixel(x0-y, y0-x, color)
	}
}

// DrawFilledCircle draws a circle given a point and radius
func DrawFilledCircle(display drivers.Displayer, x0 int16, y0 int16, r int16, color color.RGBA) {
	f := 1 - r
	ddfx := int16(1)
	ddfy := -2 * r
	x := int16(0)
	y := r
	DrawLine(display, x0, y0-r, x0, y0+r, color)
	for x < y {
		if f >= 0 {
			y--
			ddfy += 2
			f += ddfy
		}
		x++
		ddfx += 2
		f += ddfx

		DrawLine(display, x0+x, y0-y, x0+x, y0+y, color)
		DrawLine(display, x0+y, y0-x, x0+y, y0+x, color)
		DrawLine(display, x0-x, y0-y, x0-x, y0+y, color)
		DrawLine(display, x0-y, y0-x, x0-y, y0+x, color)
	}
}

// DrawTriangle draws a triangle given three points
func DrawTriangle(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, x2 int16, y2 int16, color color.RGBA) {
	DrawLine(display, x0, y0, x1, y1, color)
	DrawLine(display, x0, y0, x2, y2, color)
	DrawLine(display, x1, y1, x2, y2, color)
}


// DrawFilledTriangle draws a filled triangle given three points
func DrawFilledTriangle(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, x2 int16, y2 int16, color color.RGBA) {
	if y0 > y1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}
	if y1 > y2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}
	if y0 > y1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}

	if y0 == y2 { // y0 = y1 = y2 : it's a line
		a := x0
		b := x0
		if x1 < a {
			a = x1
		} else if x1 > b {
			b = x1
		}
		if x2 < a {
			a = x2
		} else if x2 > b {
			b = x2
		}
		DrawLine(display, a, y0, b, y0, color)
		return
	}

	dx01 := x1 - x0
	dy01 := y1 - y0
	dx02 := x2 - x0
	dy02 := y2 - y0
	dx12 := x2 - x1
	dy12 := y2 - y1

	sa := int16(0)
	sb := int16(0)
	a := int16(0)
	b := int16(0)

	last := y1 - 1
	if y1 == y2 {
		last = y1
	}

	for y := y0; y <= last; y++ {
		a = x0 + sa/dy01
		b = x0 + sb/dy02
		sa += dx01
		sb += dx02
		DrawLine(display, a, y, b, y, color)
	}

	sa = dx12 * (last - y1)
	sb = dx02 * (last - y0)

	for y := last; y <= y2; y++ {
		a = x1 + sa/dy12
		b = x0 + sb/dy02
		sa += dx12
		sb += dx02
		DrawLine(display, a, y, b, y, color)
	}
}
