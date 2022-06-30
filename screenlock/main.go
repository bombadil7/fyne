package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const defaultMessageText = "Blueberry is running. Please do not interrupt."
const (
	width  = 1080
	height = 6000
)

func main() {
	// In case X11 forwarding is enabled switch target to local display
	display := os.Getenv("DISPLAY")
	if display != ":1" {
		os.Setenv("DISPLAY", ":1")
		defer os.Setenv("DISPLAY", display)
	}
	a := app.New()
	w := a.NewWindow("Blueberry in Use!")
	w.Resize(fyne.NewSize(width, height))

	warningText := canvas.NewText("Warning!", color.RGBA{R: 0xd4, A: 0xff})
	warningText.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	warningText.Alignment = fyne.TextAlignCenter
	warningText.TextSize = 48
	warning := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(), warningText, layout.NewSpacer())

	text := ""
	if len(os.Args) == 1 {
		text = defaultMessageText
	} else {
		for _, arg := range os.Args[1:] {
			text += arg
			text += " "
			/*
				if len(text)*36 > width {
					fmt.Println("Inserting a new line")
					text = fmt.Sprintf("%s\n", text)
				} else {
					text += " "
				}
			*/
		}
	}
	messageText := canvas.NewText(text,
		color.RGBA{G: 0xa4, R: 0xf4, A: 0xff})
	messageText.Alignment = fyne.TextAlignCenter
	messageText.TextSize = 36
	message := container.New(layout.NewHBoxLayout(), layout.NewSpacer(),
		messageText, layout.NewSpacer())

	sign := makeSign()

	w.SetContent(container.New(layout.NewVBoxLayout(),
		warning, widget.NewSeparator(), message, widget.NewSeparator(), sign))
	w.ShowAndRun()
}

func makeSign() fyne.CanvasObject {
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 12

	bar := canvas.NewRectangle(color.White)

	c := container.NewWithoutLayout(bg, bar)

	bg.Resize(fyne.NewSize(900, 900))
	bg.Move(fyne.NewPos(50, 50))

	bar.Resize(fyne.NewSize(760, 160))
	bar.Move(fyne.NewPos(120, 420))
	return c
}
