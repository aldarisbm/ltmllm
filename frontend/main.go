package main

import (
	"fmt"
	"frontend/config"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	defer tidyUp()
	myApp := app.New()
	mainWindow := myApp.NewWindow(config.WindowName)
	mainWindow.Resize(fyne.NewSize(600, 800))
	mainWindow.SetFixedSize(true)

	//button := canvas.NewText("centered", color.White)
	//centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	mainWindow.SetContent(container.New(layout.NewVBoxLayout(), layout.NewSpacer(), widget.NewButtonWithIcon("send", theme.ConfirmIcon(), func() {
		fmt.Printf("Button tapped\n")
	}), layout.NewSpacer()))

	mainWindow.ShowAndRun()
}

func tidyUp() {
	fmt.Println("exited")
}
