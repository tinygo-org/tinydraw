package main

import (
	"machine"

	"image/color"

	"github.com/conejoninja/tinydraw"
	"tinygo.org/x/drivers/waveshare-epd/epd2in13x"
)

var display epd2in13x.Device

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display = epd2in13x.New(machine.SPI0, machine.P6, machine.P7, machine.P8, machine.P9)
	display.Configure(epd2in13x.Config{})

	//white := color.RGBA{0, 0, 0, 255}
	yellow := color.RGBA{255, 0, 0, 255}
	black := color.RGBA{1, 1, 1, 255}

	display.ClearBuffer()
	display.ClearDisplay()

	tinydraw.DrawLine(&display, 10, 10, 94, 10, black)
	tinydraw.DrawLine(&display, 94, 16, 10, 16, yellow)
	tinydraw.DrawLine(&display, 10, 20, 10, 202, yellow)
	tinydraw.DrawLine(&display, 16, 202, 16, 20, black)

	tinydraw.DrawLine(&display, 40, 40, 80, 80, black)
	tinydraw.DrawLine(&display, 40, 40, 80, 70, black)
	tinydraw.DrawLine(&display, 40, 40, 80, 60, black)
	tinydraw.DrawLine(&display, 40, 40, 80, 50, black)
	tinydraw.DrawLine(&display, 40, 40, 80, 40, black)

	tinydraw.DrawLine(&display, 100, 100, 40, 100, yellow)
	tinydraw.DrawLine(&display, 100, 100, 40, 90, yellow)
	tinydraw.DrawLine(&display, 100, 100, 40, 80, yellow)
	tinydraw.DrawLine(&display, 100, 100, 40, 70, yellow)
	tinydraw.DrawLine(&display, 100, 100, 40, 60, yellow)
	tinydraw.DrawLine(&display, 100, 100, 40, 50, yellow)

	tinydraw.DrawRectangle(&display, 30, 120, 20, 20, black)
	tinydraw.DrawFilledRectangle(&display, 34, 124, 12, 12, yellow)

	tinydraw.DrawCircle(&display, 52, 180, 20, black)
	tinydraw.DrawFilledCircle(&display, 52, 180, 16, yellow)

	tinydraw.DrawTriangle(&display, 60, 110, 100, 130, 84, 150, black)
	tinydraw.DrawFilledTriangle(&display, 65, 114, 96, 130, 84, 145, yellow)

	display.Display()
	display.WaitUntilIdle()
	println("You could remove power now")
}
