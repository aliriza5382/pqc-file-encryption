// Paket ve kütüphanelerin dahil edilmesi
package main

import (
	"fmt" // Konsola mesaj yazdırma
	"log" // Hata ayıklama mesajları yönetme

	"pqc-file-encryption/pqc" // Post-kuantum şifreleme işlemlerini içeren modül

	"github.com/sqweek/dialog" // Kullanıcıdan dosya seçmesini sağlamak için kütüphane
)

func main() {
	for {
		fmt.Println("\n===================================")
		fmt.Println(" Post-Kuantum Şifreleme Aracı")
		fmt.Println(" 1️-) TXT Dosyası Seç ve Şifrele")
		fmt.Println(" 2️-) TXT Dosyası Seç ve Şifreyi Çöz")
		fmt.Println(" 3️-) Çıkış")
		fmt.Println("===================================")
		fmt.Print(" Seçiminizi girin: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			// **Şifrelenecek metin dosyasını seç**
			inputPath, err := dialog.File().Title("Şifrelenecek Dosyayı Seç").Load()
			if err != nil {
				log.Println("Dosya seçme iptal edildi.")
				continue
			}

			// **Anahtarları yükle veya üret**
			publicKey, privateKey, err := pqc.LoadKeys()
			if err != nil {
				fmt.Println("Anahtar bulunamadı, yeni anahtar oluşturuluyor...")
				publicKey, privateKey, err = pqc.GenerateKyberKeyPair()
				if err != nil {
					log.Fatalf("Anahtar üretme hatası: %v", err)
				}
				pqc.SaveKeys(publicKey, privateKey)
			}

			// **Dosyayı şifrele**
			// Kyber512 ile üretilen açık anahtar (public key) kullanılarak AES-GCM şifreleme yapılır ve dosya şifrelenir.
			err = pqc.EncryptFile(inputPath, inputPath, publicKey)
			if err != nil {
				log.Fatalf("Şifreleme hatası: %v", err)
			}
			fmt.Println("Dosya başarıyla şifrelendi:", inputPath)

		case "2":
			// **Şifrelenmiş dosyayı seç**
			inputPath, err := dialog.File().Title("Çözülecek Dosyayı Seç").Load()
			if err != nil {
				log.Println("Dosya seçme iptal edildi.")
				continue
			}

			// **Anahtarları yükle**
			_, privateKey, err := pqc.LoadKeys()
			if err != nil {
				log.Fatalf("Anahtar bulunamadı, önce bir dosya şifreleyin!")
			}

			// **Dosyanın şifresini çöz**
			// AES-GCM ve Kyber512 ile şifre çözme işlemi yapılır ve dosya orijinal haline geri döndürülür.
			err = pqc.DecryptFile(inputPath, inputPath, privateKey)
			if err != nil {
				log.Fatalf("Şifre çözme hatası: %v", err)
			}
			fmt.Println("Dosya başarıyla çözüldü:", inputPath)

		case "3":
			fmt.Println("Çıkış yapılıyor...")
			return

		default:
			fmt.Println("Geçersiz seçim! Lütfen 1, 2 veya 3 girin.")
		}
	}
}
