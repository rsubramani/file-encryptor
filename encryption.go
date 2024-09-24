package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "golang.org/x/crypto/pbkdf2"
    "hash"
    "io"
    "os"
    "golang.org/x/crypto/sha3"
)

// EncryptFile takes a file and password, then encrypts the file using AES-256.
func EncryptFile(inputFile, outputFile, password string) error {
    key := deriveKey(password)

    inFile, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer inFile.Close()

    outFile, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer outFile.Close()

    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }

    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return err
    }

    outFile.Write(iv)
    stream := cipher.NewCFBEncrypter(block, iv)

    writer := &cipher.StreamWriter{S: stream, W: outFile}
    if _, err := io.Copy(writer, inFile); err != nil {
        return err
    }

    return nil
}

// DecryptFile takes the encrypted file and password to reverse the encryption process.
func DecryptFile(inputFile, outputFile, password string) error {
    key := deriveKey(password)

    inFile, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer inFile.Close()

    outFile, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer outFile.Close()

    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(inFile, iv); err != nil {
        return err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }

    stream := cipher.NewCFBDecrypter(block, iv)
    reader := &cipher.StreamReader{S: stream, R: inFile}

    if _, err := io.Copy(outFile, reader); err != nil {
        return err
    }

    return nil
}

func deriveKey(password string) []byte {
    salt := []byte("random-salt") // In a real app, generate a unique salt for each encryption
    return pbkdf2.Key([]byte(password), salt, 4096, 32, func() hash.Hash {
        return sha3.New256()
    })
}
