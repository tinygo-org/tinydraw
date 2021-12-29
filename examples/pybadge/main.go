package main

import (
	"image/color"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinydraw/examples/initdisplay"
)

func main() {
	display := initdisplay.InitDisplay()

	tinydraw.Line(display, 10, 10, 94, 10, getRainbowRGB(0))
	tinydraw.Line(display, 94, 16, 10, 16, getRainbowRGB(15))
	tinydraw.Line(display, 10, 20, 10, 118, getRainbowRGB(30))
	tinydraw.Line(display, 16, 118, 16, 20, getRainbowRGB(45))

	tinydraw.Line(display, 40, 40, 80, 80, getRainbowRGB(60))
	tinydraw.Line(display, 40, 40, 80, 70, getRainbowRGB(75))
	tinydraw.Line(display, 40, 40, 80, 60, getRainbowRGB(90))
	tinydraw.Line(display, 40, 40, 80, 50, getRainbowRGB(105))
	tinydraw.Line(display, 40, 40, 80, 40, getRainbowRGB(120))

	tinydraw.Line(display, 100, 100, 40, 100, getRainbowRGB(135))
	tinydraw.Line(display, 100, 100, 40, 90, getRainbowRGB(150))
	tinydraw.Line(display, 100, 100, 40, 80, getRainbowRGB(165))
	tinydraw.Line(display, 100, 100, 40, 70, getRainbowRGB(180))
	tinydraw.Line(display, 100, 100, 40, 60, getRainbowRGB(195))
	tinydraw.Line(display, 100, 100, 40, 50, getRainbowRGB(210))

	tinydraw.Rectangle(display, 30, 106, 120, 20, getRainbowRGB(225))
	tinydraw.FilledRectangle(display, 34, 110, 112, 12, getRainbowRGB(130))

	tinydraw.Circle(display, 120, 30, 20, getRainbowRGB(240))
	tinydraw.FilledCircle(display, 120, 30, 16, getRainbowRGB(145))

	tinydraw.Triangle(display, 120, 102, 100, 80, 152, 46, getRainbowRGB(155))
	tinydraw.FilledTriangle(display, 120, 98, 104, 80, 144, 54, getRainbowRGB(90))

	display.Display()

}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
