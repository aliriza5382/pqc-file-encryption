package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pqc-file-encryption/pqc"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Post-Kuantum Şifreleme")

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Şifrelenecek dosyanın yolu")

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("Çıktı dosyasının adı")

	encryptButton := widget.NewButton("Şifrele", func() {
		publicKey, privateKey, err := pqc.LoadKeys()
		if err != nil {
			publicKey, privateKey, err = pqc.GenerateKyberKeyPair()
			if err != nil {
				log.Fatalf("Anahtar üretme hatası: %v", err)
			}
			pqc.SaveKeys(publicKey, privateKey)
		}

		err = pqc.EncryptFile(inputEntry.Text, outputEntry.Text, publicKey)
		if err != nil {
			log.Fatalf("Şifreleme hatası: %v", err)
		}
		fmt.Println("Dosya başarıyla şifrelendi:", outputEntry.Text)
	})

	decryptButton := widget.NewButton("Şifreyi Çöz", func() {
		_, privateKey, err := pqc.LoadKeys()
		if err != nil {
			log.Fatalf("Anahtar bulunamadı, önce bir dosya şifreleyin!")
		}

		err = pqc.DecryptFile(inputEntry.Text, outputEntry.Text, privateKey)
		if err != nil {
			log.Fatalf("Şifre çözme hatası: %v", err)
		}
		fmt.Println("Dosya başarıyla çözüldü:", outputEntry.Text)
	})

	content := container.NewVBox(inputEntry, outputEntry, encryptButton, decryptButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
