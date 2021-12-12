package tinydraw // import "tinygo.org/x/tinydraw"

import (
	"image/color"

	"errors"

	"tinygo.org/x/drivers"
)

//This implementation of bresenham always steps in increasing y direction. This enables it to be used to draw filled triangles
type bresenham struct {
	cx int16
	cy int16

	x0 int16
	y0 int16

	x1 int16
	y1 int16

	adx   int16
	dx    int16
	dy    int16
	steep bool
	xstep int16
	err   int16
}

func newBresenham(x0, y0, x1, y1 int16) bresenham {

	b := bresenham{
		x0: x0,
		y0: y0,
		x1: x1,
		y1: y1,
	}

	b.dx = x1 - x0
	b.adx = b.dx
	b.xstep = 1
	if b.adx < 0 {
		b.adx = -b.adx
		b.xstep = -1
	}
	b.dy = y1 - y0
	b.steep = b.dy > b.adx
	if b.steep {
		b.err = b.dy / 2
	} else {
		b.err = b.adx / 2
	}

	return b
}

func (b *bresenham) Next() bool {
	if b.cx == b.x1 && b.cy == b.y1 {
		return false
	}

	b.cx = b.x0
	b.cy = b.y0

	if b.steep {
		b.err -= b.adx
		if b.err < 0 {
			if b.x0 != b.x1 {
				b.x0 += b.xstep
			}
			b.err += b.dy
		}

		if b.y0 != b.y1 {
			b.y0++
		}
	} else {
		b.err -= b.dy
		if b.err < 0 {
			if b.y0 != b.y1 {
				b.y0++
			}
			b.err += b.adx
		}

		if b.x0 != b.x1 {
			b.x0 += b.xstep
		}
	}

	return true
}

func (b *bresenham) Cur() (rx int16, ry int16) {
	return b.cx, b.cy
}

func horizLine(display drivers.Displayer, x0 int16, y0 int16, x1 int16, color color.RGBA) {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	for ; x0 <= x1; x0++ {
		display.SetPixel(x0, y0, color)
	}
}

func Line(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, color color.RGBA) {

	if y0 > y1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}

	if x0 == x1 {
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		for ; y0 <= y1; y0++ {
			display.SetPixel(x0, y0, color)
		}
	} else if y0 == y1 {
		horizLine(display, x0, y0, x1, color)
	} else { // Bresenham

		b := newBresenham(x0, y0, x1, y1)

		for b.Next() {
			x, y := b.Cur()
			display.SetPixel(x, y, color)
		}
	}
}

// Rectangle draws a rectangle given a point, width and height
func Rectangle(display drivers.Displayer, x int16, y int16, w int16, h int16, color color.RGBA) error {
	if w <= 0 || h <= 0 {
		return errors.New("empty rectangle")
	}
	Line(display, x, y, x+w-1, y, color)
	Line(display, x, y, x, y+h-1, color)
	Line(display, x+w-1, y, x+w-1, y+h-1, color)
	Line(display, x, y+h-1, x+w-1, y+h-1, color)
	return nil
}

// FilledRectangle draws a filled rectangle given a point, width and height
func FilledRectangle(display drivers.Displayer, x int16, y int16, w int16, h int16, color color.RGBA) error {
	if w <= 0 || h <= 0 {
		return errors.New("empty rectangle")
	}
	for i := x; i < x+w; i++ {
		Line(display, i, y, i, y+h-1, color)
	}
	return nil
}

// Circle draws a circle given a point and radius
func Circle(display drivers.Displayer, x0 int16, y0 int16, r int16, color color.RGBA) {
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

// FilledCircle draws a filled circle given a point and radius
func FilledCircle(display drivers.Displayer, x0 int16, y0 int16, r int16, color color.RGBA) {
	f := 1 - r
	ddfx := int16(1)
	ddfy := -2 * r
	x := int16(0)
	y := r
	Line(display, x0, y0-r, x0, y0+r, color)
	for x < y {
		if f >= 0 {
			y--
			ddfy += 2
			f += ddfy
		}
		x++
		ddfx += 2
		f += ddfx

		Line(display, x0+x, y0-y, x0+x, y0+y, color)
		Line(display, x0+y, y0-x, x0+y, y0+x, color)
		Line(display, x0-x, y0-y, x0-x, y0+y, color)
		Line(display, x0-y, y0-x, x0-y, y0+x, color)
	}
}

// Triangle draws a triangle given three points
func Triangle(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, x2 int16, y2 int16, color color.RGBA) {

	if y0 > y1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}
	if y0 > y2 {
		x0, y0, x2, y2 = x2, y2, x0, y0
	}
	if y1 > y2 {
		x2, y2, x1, y1 = x1, y1, x2, y2
	}

	Line(display, x0, y0, x1, y1, color)
	Line(display, x0, y0, x2, y2, color)
	Line(display, x1, y1, x2, y2, color)
}

type rowExtent struct {
	minX int16
	maxX int16
}

func (r *rowExtent) set(x int16) {
	if x < r.minX {
		r.minX = x
	}

	if x > r.maxX {
		r.maxX = x
	}
}

func minMaxforRow(b *bresenham) (int16, int16) {

	cx, cy := b.Cur()

	minx := cx
	maxx := cx

	for b.Next() {
		nx, ny := b.Cur()
		if ny != cy {
			break //Next line
		}

		if nx < minx {
			minx = nx
		}

		if nx > maxx {
			maxx = nx
		}
	}

	return minx, maxx
}

func combineMinMax(min1, max1, min2, max2 int16) (min int16, max int16) {

	if min2 < min1 {
		min = min2
	} else {
		min = min1
	}

	if max2 > max1 {
		max = max2
	} else {
		max = max1
	}

	return
}

func FilledTriangle(display drivers.Displayer, x0 int16, y0 int16, x1 int16, y1 int16, x2 int16, y2 int16, c color.RGBA) {

	//Sort by Y coordinate
	if y0 > y1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}
	if y0 > y2 {
		x0, y0, x2, y2 = x2, y2, x0, y0
	}
	if y1 > y2 {
		x2, y2, x1, y1 = x1, y1, x2, y2
	}

	//Prepare to draw all 3 lines that form the triangle
	b01 := newBresenham(x0, y0, x1, y1)
	b02 := newBresenham(x0, y0, x2, y2)
	b12 := newBresenham(x1, y1, x2, y2)

	b01.Next()
	b02.Next()
	b12.Next()

	//Top half of triangle
	for y := y0; y < y1; y++ {

		//Identify the min and max X coords for this row
		b01minx, b01maxx := minMaxforRow(&b01)
		b02minx, b02maxx := minMaxforRow(&b02)

		minx, maxx := combineMinMax(b01minx, b01maxx, b02minx, b02maxx)

		//Draw a line from the min to max to fill this row
		Line(display, minx, y, maxx, y, c)
	}

	//Bottom half of triangle (continues drawing of b02)
	for y := y1; y < y2; y++ {

		//Identify the min and max X coords for this row
		b12minx, b12maxx := minMaxforRow(&b12)
		b02minx, b02maxx := minMaxforRow(&b02)

		minx, maxx := combineMinMax(b12minx, b12maxx, b02minx, b02maxx)

		//Draw a line from the min to max to fill this row
		horizLine(display, minx, y, maxx, c)
	}

}
