package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/aldarisbm/ltmllm/backend/llm"
	"github.com/aldarisbm/ltmllm/config"
)

func NewWindow(cb llm.ChatBot) fyne.Window {
	myApp := app.New()
	mainWindow := myApp.NewWindow(config.WindowName)
	mainWindow.Resize(fyne.NewSize(600, 800))
	mainWindow.SetFixedSize(true)

	nl := widget.NewLabel("")
	c := widget.NewCard("Chat", "Chat with the bot", nl)
	mainWindow.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			c,
			layout.NewSpacer(),
			widget.NewButtonWithIcon("send", theme.ConfirmIcon(), func() {
				msg := cb.Chat("tell me about the history of golang")
				c.SetContent(widget.NewLabel(msg))
			}),
			layout.NewSpacer()))

	return mainWindow
}
