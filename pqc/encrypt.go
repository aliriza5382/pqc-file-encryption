package pqc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"

	"github.com/cloudflare/circl/pke/kyber/kyber512"
	"github.com/google/uuid"
)

// **UUID tabanlı nonce üretme fonksiyonu**
func GenerateUUIDNonce() ([]byte, error) {
	id := uuid.New() // Yeni UUID oluştur
	nonce := id[:12] // İlk 12 baytı al (AES-GCM için uygun uzunluk)
	return nonce, nil
}

// **AES-GCM ile Dosya Şifreleme + Kyber512 ile Anahtar Şifreleme**
func EncryptFile(inputPath, outputPath string, publicKey *kyber512.PublicKey) error {
	// Dosyayı oku
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	plaintext, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// ***Her şifrelemede rastgele AES anahtarı oluştur (32 byte)
	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return err
	}

	// Kyber512 ile AES anahtarını şifrele
	ciphertextKey := make([]byte, kyber512.CiphertextSize)
	seed := make([]byte, kyber512.EncryptionSeedSize)
	if _, err := rand.Read(seed); err != nil {
		return err
	}
	publicKey.EncryptTo(ciphertextKey, aesKey, seed)

	// AES-GCM ile dosyayı şifrele
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	// ***UUID tabanlı nonce üretimi**  (GCM)
	nonce, err := GenerateUUIDNonce()
	if err != nil {
		return errors.New("Nonce oluşturulamadı")
	}

	// *** GCM Sayıcı
	stream, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// *** Tag (Doğrulama Kodu - Authentication Tag)  (GCM)
	// Bu hem şifrelenmiş veriyi hem de Authentication Tag’i üretir.
	// Bu nedenle şifre çözme sırasında yanlış bir anahtar veya veri değiştirilirse hata alınmasını sağlar.
	ciphertext := stream.Seal(nil, nonce, plaintext, nil)

	// Şifrelenmiş verileri dosyaya yaz
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	outputFile.Write(nonce)               // Nonce yaz
	outputFile.Write(ciphertextKey)       // Şifrelenmiş AES anahtarını yaz
	_, err = outputFile.Write(ciphertext) // Şifrelenmiş içeriği yaz

	return err
}
