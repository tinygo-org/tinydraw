package tinydraw

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Set to true to generate missing reference images for new tests
const generateMissingReferenceImages = false

type testDisplay struct {
	*image.RGBA
}

func (d *testDisplay) Size() (x, y int16) {
	return int16(d.Rect.Dx()), int16(d.Rect.Dy())
}

func (d *testDisplay) SetPixel(x, y int16, c color.RGBA) {
	d.Set(int(x), int(y), c)
}

func newTestDisplay() testDisplay {
	i := image.NewRGBA(image.Rect(0, 0, 320, 240))
	return testDisplay{
		RGBA: i,
	}
}

func compareWithReference(name string, d testDisplay) error {

	referenceFilePath := filepath.Join("testdata", name+".png")

	referenceFile, err := os.Open(referenceFilePath)
	switch err.(type) {
	case *os.PathError:
		if generateMissingReferenceImages {
			saveImgFile, err := os.Create(referenceFilePath)
			if err != nil {
				return err
			}

			err = png.Encode(saveImgFile, d.RGBA)
			return err
		} else {
			return fmt.Errorf("Missing reference image %v. Set generateMissingReferenceImages in tinydraw_test.go to true to autogenerate.", referenceFilePath)
		}
	default:
		if err != nil {
			return err
		}
	}

	referenceImage, err := png.Decode(referenceFile)
	if err != nil {
		return err
	}

	sx, sy := d.Size()

	differenceDetected := false

	for y := 0; y < int(sy); y++ {
		for x := 0; x < int(sx); x++ {
			refPixel := referenceImage.At(x, y)
			testPixel := d.RGBA.At(x, y)

			refR, refG, refB, refA := refPixel.RGBA()
			testR, testG, testB, testA := testPixel.RGBA()

			if refR != testR || refG != testG || refB != testB || refA != testA {
				differenceDetected = true
				d.RGBA.Set(x, y, color.RGBA{255, 0, 0, 255}) //Mark the pixel as bad for debugging
			}
		}
	}

	if differenceDetected {
		diffFilePath := filepath.Join("testdata", "DIFF_"+name+".png")
		saveImgFile, err := os.Create(diffFilePath)
		if err != nil {
			return err
		}

		err = png.Encode(saveImgFile, d.RGBA)
		return fmt.Errorf("Image contains differences see %v" + diffFilePath)
	}

	return nil
}

// Display sends the buffer (if any) to the screen.
func (d *testDisplay) Display() error {
	return nil
}

func TestFilledTriangleClockMinuteHand(t *testing.T) {

	black := color.RGBA{0, 0, 0, 255}
	unfilledDisplay := newTestDisplay()
	filledDisplay := newTestDisplay()

	x := int16(160)
	y := int16(120)

	const (
		hourRadius   = 55
		minuteRadius = 95
		hourWidth    = 6
		minuteWidth  = 5
	)

	testTime, _ := time.Parse("15:04:05", "10:02:25")
	for i := 0; i < 60; i++ {
		// Draw the clock hands.
		minuteAngle := float64(testTime.Minute()) / 60 * 2 * math.Pi
		mx, my := math.Sincos(minuteAngle)

		x0 := x - int16(minuteWidth*my)
		y0 := y - int16(minuteWidth*mx)
		x1 := x + int16(minuteWidth*my)
		y1 := y + int16(minuteWidth*mx)
		x2 := x + int16(minuteRadius*mx)
		y2 := y - int16(minuteRadius*my)

		FilledTriangle(

			&filledDisplay,
			x0, y0,
			x1, y1,
			x2, y2,
			black)

		Triangle(
			&unfilledDisplay,
			x0, y0,
			x1, y1,
			x2, y2,
			black)

		testTime = testTime.Add(1 * time.Minute)
	}

	err := compareWithReference("TestFilledTriangleClockMinuteHand", filledDisplay)
	if err != nil {
		t.Errorf("Fail %v", err)
	}

	err = compareWithReference("TestTriangleClockMinuteHand", unfilledDisplay)
	if err != nil {
		t.Errorf("Fail %v", err)
	}
}

type triangleTest struct {
	name string
	x0   int16
	y0   int16
	x1   int16
	y1   int16
	x2   int16
	y2   int16
}

func TestTriangles(t *testing.T) {

	tests := []triangleTest{
		{"TopAndBottom1", 160, 0, 80, 240, 213, 120}, //Basic triangle with top and bottom half
		{"TopAndBottom2", 160, 0, 213, 240, 80, 120}, //Basic triangle with top and bottom half
		{"NoTop", 80, 50, 160, 50, 120, 200},         //No top
		{"NoBottom", 80, 200, 160, 200, 120, 50},     //No Bottom
		{"HorizLine", 80, 200, 160, 200, 213, 200},   //Horizontal line
		{"VertLine", 80, 80, 80, 120, 80, 200},       //Vertical line

		{"ClockHand1", 160, 115, 160, 125, 255, 120}, //Previously rendered too Wide
		{"ClockHand2", 160, 116, 160, 124, 254, 129}, //Previusly fill extendended to the left outside of triangle
		{"ClockHand3", 165, 120, 155, 120, 160, 215}, //Too Wide
	}

	colors := []color.RGBA{
		{255, 0, 0, 255},
		{0, 255, 0, 255},
		{0, 0, 255, 255},
		{255, 255, 0, 255},
		{0, 255, 255, 255},
		{0, 0, 0, 255},
	}

	for _, test := range tests {

		unfilledDisplay := newTestDisplay()

		//Test all permutations of coordinates
		Triangle(&unfilledDisplay, test.x0, test.y0, test.x1, test.y1, test.x2, test.y2, colors[0])
		Triangle(&unfilledDisplay, test.x0, test.y0, test.x2, test.y2, test.x1, test.y1, colors[1])
		Triangle(&unfilledDisplay, test.x1, test.y1, test.x0, test.y0, test.x2, test.y2, colors[2])
		Triangle(&unfilledDisplay, test.x1, test.y1, test.x2, test.y2, test.x0, test.y0, colors[3])
		Triangle(&unfilledDisplay, test.x2, test.y2, test.x0, test.y0, test.x1, test.y1, colors[4])
		Triangle(&unfilledDisplay, test.x2, test.y2, test.x1, test.y1, test.x0, test.y0, colors[5])

		err := compareWithReference(fmt.Sprintf("TestTriangles_%v", test.name), unfilledDisplay)
		if err != nil {
			t.Errorf("Fail %v", err)
		}

		filledDisplay := newTestDisplay()

		//Test all permutations of coordinates
		FilledTriangle(&filledDisplay, test.x0, test.y0, test.x1, test.y1, test.x2, test.y2, colors[0])
		FilledTriangle(&filledDisplay, test.x0, test.y0, test.x2, test.y2, test.x1, test.y1, colors[1])
		FilledTriangle(&filledDisplay, test.x1, test.y1, test.x0, test.y0, test.x2, test.y2, colors[2])
		FilledTriangle(&filledDisplay, test.x1, test.y1, test.x2, test.y2, test.x0, test.y0, colors[3])
		FilledTriangle(&filledDisplay, test.x2, test.y2, test.x0, test.y0, test.x1, test.y1, colors[4])
		FilledTriangle(&filledDisplay, test.x2, test.y2, test.x1, test.y1, test.x0, test.y0, colors[5])

		err = compareWithReference(fmt.Sprintf("TestTriangles_%v_Filled", test.name), filledDisplay)
		if err != nil {
			t.Errorf("Fail %v", err)
		}
	}
}
