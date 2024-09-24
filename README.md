# Go File Encryptor

A simple file encryption tool built in Go with a graphical user interface (GUI) using the Fyne framework. This application allows users to encrypt files using AES-256 encryption. It includes features like password strength checking, a progress bar, error handling, and an option to save the encryption key to a file.

## Features

- **File Encryption**: Encrypt files using AES-256 encryption.
- **Password Strength Checker**: Provides feedback on the strength of the user's password.
- **Progress Bar**: Displays the progress of the encryption process.
- **Error Handling**: Shows error messages for issues like incorrect passwords or missing files.
- **GUI Interface**: User-friendly interface built with the Fyne framework.
- **Save Key Option**: Option to save the encryption key to a file.

## Prerequisites

- **Go Programming Language**: Ensure you have Go installed on your system (version 1.20 or higher).
- **Fyne Framework**: The GUI is built using Fyne, which will be installed as a dependency.

## Installation

1. **Clone the Repository**
```
git clone https://github.com/rsubramani/gofileencryptor.git cd gofileencryptor
```

2. **Initialize Go Modules**
```
go mod init gofileencryptor
```

3. **Install Dependencies**
```
go get fyne.io/fyne/v2 go get golang.org/x/crypto
```

4. **Build the Application**
```
go build
```

## Running the Application

To run the application, use the following command:
```
go run .
```

Alternatively, if you have built the application, you can run the executable:
```
./gofileencryptor
```


## Usage

1. **Launch the Application**

   Run the application using the instructions above. A GUI window will appear.

2. **Password Entry**

   - Enter a password in the password field.
   - The application will display the strength of your password (Weak, Medium, Strong).

3. **Select File for Encryption**

   - Click the **"Select File"** button.
   - A file dialog will open. Navigate to and select the file you wish to encrypt.

4. **Encryption Process**

   - Once a file is selected, the encryption process will begin automatically.
   - The progress bar will update to show the progress.
   - Upon completion, a success message will appear.
   - The encrypted file will be saved in the same directory as the original file, with a `.enc` extension added.

5. **Error Handling**

   - If there is an error during encryption (e.g., file not found), an error message will be displayed.


## Code Overview

### `main.go`

- Initializes the Fyne application.
- Sets up the main window and UI components.
- Runs the main event loop.

### `ui.go`

- Contains functions to create UI components:
  - `CreatePasswordUI()`: Creates the password entry field and strength label.
  - `CreateFileSelectionArea()`: Creates the file selection button and handles the file dialog.
- Handles user interactions and updates the progress bar.

### `encryption.go`

- Contains the `EncryptFile` function that performs the encryption of the file.
- Uses AES-256 encryption with a key derived from the user's password.

### `utils.go`

- Contains utility functions:
  - `getPasswordStrength()`: Evaluates the strength of the user's password.
  - `saveKeyToFile()`: Saves the encryption key to a file (if implemented).

## Dependencies

- **Fyne GUI Framework**: [fyne.io/fyne/v2](https://github.com/fyne-io/fyne)
- **Go Crypto Library**: [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)

## Notes

- The application currently saves encrypted files in the same directory as the original file with an added `.enc` extension.
- Decryption functionality can be added following a similar pattern to encryption.
- Ensure that you keep your encryption keys and passwords secure.

## License

This project is licensed under the MIT License.
