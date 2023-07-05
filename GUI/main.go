/*
*@Author: Wenqiang
*@Date:   2023/7/5
*@Time:   17:57
 */
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app2.New()
	win := app.NewWindow("Tools")
	win.Resize(fyne.NewSize(900, 600))

	button := widget.NewButton("Logging", func() {

		myWindow := app.NewWindow("Form")

		nameEntry := widget.NewEntry()
		passEntry := widget.NewPasswordEntry()

		form := widget.NewForm(
			&widget.FormItem{Text: "Name", Widget: nameEntry},
			&widget.FormItem{Text: "Pass", Widget: passEntry},
		)
		form.OnSubmit = func() {
			fmt.Println("name:", nameEntry.Text, "pass:", passEntry.Text, "login in")
		}
		form.OnCancel = func() {
			fmt.Println("login canceled")
		}

		myWindow.SetContent(form)
		myWindow.Resize(fyne.NewSize(300, 300))
		myWindow.Show()
	})
	win.SetContent(button)
	win.ShowAndRun()

}
