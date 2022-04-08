package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type AntiBoredom struct {
	Text string `json:"activity"`
}

func getAntiBoredom() (AntiBoredom, error) {
    var fact AntiBoredom

	resp, err := client.Get("https://www.boredapi.com/api/activity/")
	if err != nil {
		return AntiBoredom{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return AntiBoredom{}, err
	}

	return fact, nil
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()

	win := a.NewWindow("Anti-Boredom")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Got Bored?", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Activity", func() {
		fact, err := getAntiBoredom()

		if err != nil {
			dialog.ShowError(err, win)
		} else {
			factText.SetText(fact.Text)
		}
	})
	
	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), factText)

	win.SetContent(vBox)
	win.ShowAndRun()
}
