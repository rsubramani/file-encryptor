package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    // Use app.NewWithID to resolve the preferences API error
    myApp := app.NewWithID("com.yourapp.fileencryptor")
    myWindow := myApp.NewWindow("File Encryptor")

    // UI elements for the application
    passwordEntry, strengthLabel := CreatePasswordUI()
    progressBar := widget.NewProgressBar()
    dragLabel, dragArea := CreateFileSelectionArea(myWindow, progressBar, passwordEntry)

    content := container.NewVBox(
        dragLabel,
        dragArea,
        passwordEntry,
        strengthLabel,
        progressBar,
    )

    myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(500, 300))
    myWindow.ShowAndRun()
}
