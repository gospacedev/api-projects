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

type ChuckNorrisJokes struct {
	Text string `json:"value"`
}

func getChuckNorrisJokes() (ChuckNorrisJokes, error) {
    var fact ChuckNorrisJokes

	resp, err := client.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return ChuckNorrisJokes{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return ChuckNorrisJokes{}, err
	}

	return fact, nil
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()

	win := a.NewWindow("Chuck Norris Jokes")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Get Chuck Norris Jokes", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Joke", func() {
		fact, err := getChuckNorrisJokes()

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
