package main

import (
    "errors"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
    "time"
)

// CreatePasswordUI creates the password entry field with a strength label
func CreatePasswordUI() (*widget.Entry, *widget.Label) {
    passwordEntry := widget.NewPasswordEntry()
    strengthLabel := widget.NewLabel("Password Strength: ")

    passwordEntry.OnChanged = func(password string) {
        strength := getPasswordStrength(password)
        strengthLabel.SetText("Password Strength: " + strength)
    }

    return passwordEntry, strengthLabel
}

// CreateFileSelectionArea creates a file selection area using a file dialog
func CreateFileSelectionArea(window fyne.Window, progressBar *widget.ProgressBar, passwordEntry *widget.Entry) (*canvas.Text, fyne.CanvasObject) {
    dragLabel := canvas.NewText("Select a file for encryption", nil)

    // Button to open file dialog
    fileSelectBtn := widget.NewButton("Select File", func() {
        dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
            if err != nil || reader == nil {
                showError("Failed to open file", window)
                return
            }
            defer reader.Close()

            filePath := reader.URI().Path()
            go processFile(filePath, progressBar, passwordEntry.Text, window)
        }, window)
    })

    fileSelectContainer := container.NewVBox(dragLabel, fileSelectBtn)
    return dragLabel, fileSelectContainer
}

// processFile simulates file processing and updates the progress bar
func processFile(file string, progressBar *widget.ProgressBar, password string, window fyne.Window) {
    progressBar.SetValue(0)
    time.Sleep(1 * time.Second)

    // Simulating encryption process
    err := EncryptFile(file, file+".enc", password)
    if err != nil {
        showError("Encryption Failed: "+err.Error(), window)
        return
    }

    for i := 0.0; i <= 1.0; i += 0.1 {
        progressBar.SetValue(i)
        time.Sleep(200 * time.Millisecond)
    }

    showSuccess("Encryption Complete!", window)
}

// showError displays an error dialog
func showError(message string, window fyne.Window) {
    err := errors.New(message)
    dialog.ShowError(err, window)
}

// showSuccess displays a success dialog
func showSuccess(message string, window fyne.Window) {
    dialog.ShowInformation("Success", message, window)
}

