package main

import (
	"machine"

	"image/color"

	"github.com/conejoninja/primitive"
	"github.com/tinygo-org/drivers/waveshare-epd/epd2in13x"
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

	primitive.DrawLine(&display, 10, 10, 94, 10, black)
	primitive.DrawLine(&display, 94, 16, 10, 16, yellow)
	primitive.DrawLine(&display, 10, 20, 10, 202, yellow)
	primitive.DrawLine(&display, 16, 202, 16, 20, black)

	primitive.DrawLine(&display, 40, 40, 80, 80, black)
	primitive.DrawLine(&display, 40, 40, 80, 70, black)
	primitive.DrawLine(&display, 40, 40, 80, 60, black)
	primitive.DrawLine(&display, 40, 40, 80, 50, black)
	primitive.DrawLine(&display, 40, 40, 80, 40, black)

	primitive.DrawLine(&display, 100, 100, 40, 100, yellow)
	primitive.DrawLine(&display, 100, 100, 40, 90, yellow)
	primitive.DrawLine(&display, 100, 100, 40, 80, yellow)
	primitive.DrawLine(&display, 100, 100, 40, 70, yellow)
	primitive.DrawLine(&display, 100, 100, 40, 60, yellow)
	primitive.DrawLine(&display, 100, 100, 40, 50, yellow)

	primitive.DrawRectangle(&display, 30, 120, 20, 20, black)
	primitive.DrawFilledRectangle(&display, 34, 124, 12, 12, yellow)

	primitive.DrawCircle(&display, 52, 180, 20, black)
	primitive.DrawFilledCircle(&display, 52, 180, 16, yellow)

	primitive.DrawTriangle(&display, 60, 110, 100, 130, 84, 150, black)
	primitive.DrawFilledTriangle(&display, 65, 114, 96, 130, 84, 145, yellow)

	display.Display()
	display.WaitUntilIdle()
	println("You could remove power now")
}
