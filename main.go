package main

import (
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a headless test app
	app := test.NewApp()
	defer app.Quit()

	// Create some content (e.g., a label)
	label := widget.NewLabel("Hello, Fyne!")
	label.TextStyle.Bold = true

	// Create a headless window and set its content
	win := test.NewWindow(label)
	defer win.Close()

	win.Resize(fyne.NewSize(200, 100))

	img := win.Canvas().Capture()

	file, err := os.Create("screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
