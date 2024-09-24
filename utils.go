package main

import "os"

// Check password strength based on length
func getPasswordStrength(password string) string {
    if len(password) < 8 {
        return "Weak"
    } else if len(password) < 12 {
        return "Medium"
    }
    return "Strong"
}

// Save the encryption key to a file
func saveKeyToFile(key []byte, filename string) error {
    return os.WriteFile(filename, key, 0644)
}
