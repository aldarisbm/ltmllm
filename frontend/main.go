package frontend

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/aldarisbm/ltmllm/backend/chatbot"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/aldarisbm/ltmllm/shared"
	"github.com/google/uuid"
	"io"
	"time"
)

func NewWindow(cb *chatbot.ChatBot) fyne.Window {
	mainWindow := setUpApp()
	convo := shared.NewConversation()

	nl := widget.NewLabel("")
	nl.Wrapping = fyne.TextWrapWord

	entry := widget.NewEntry()

	submitButton := widget.NewButtonWithIcon("submit", theme.ColorChromaticIcon(), func() {
		userPrompt := getUserMessage(entry.Text)
		response := &shared.Message{
			Id:         uuid.New(),
			User:       "chatbot",
			CreateTime: time.Now(),
		}
		nl.SetText(fmt.Sprintf("%s: %s\n", userPrompt.User, entry.Text))
		nl.SetText(nl.Text + fmt.Sprintf("%s: ", response.User))

		var responseText string

		msg := cb.ChatStream(entry.Text)
		for {
			m, err := msg.Recv()
			if err == io.EOF {
				break
			}
			nl.SetText(nl.Text + m.Choices[0].Delta.Content)
			responseText += m.Choices[0].Delta.Content
		}
		response.Content = nl.Text
		defer saveMessages(convo, userPrompt, response)
	})

	formBox := widget.NewForm(
		&widget.FormItem{Widget: entry},
		&widget.FormItem{Widget: submitButton},
	)

	mainWindow.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			nl,
			layout.NewSpacer(),
			formBox,
			layout.NewSpacer(),
		),
	)

	return mainWindow
}

func saveMessages(convo shared.Conversation, prompt *shared.Message, response *shared.Message) {
	convo.AddMessage(prompt)
	convo.AddMessage(response)
}

func getUserMessage(text string) *shared.Message {
	userPrompt := &shared.Message{
		Content:    text,
		Id:         uuid.New(),
		User:       "aldarisbm",
		CreateTime: time.Now(),
	}
	return userPrompt
}

func setUpApp() fyne.Window {
	frontend := app.New()

	mainWindow := frontend.NewWindow(config.WindowName)
	mainWindow.Resize(fyne.NewSize(400, 600))
	return mainWindow
}
