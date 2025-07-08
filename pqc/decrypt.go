package pqc

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"

	"github.com/cloudflare/circl/pke/kyber/kyber512"
)

// **AES-GCM ile Dosya Şifre Çözme + Kyber512 ile Anahtar Çözme**
func DecryptFile(inputPath, outputPath string, privateKey *kyber512.PrivateKey) error {
	// Şifrelenmiş dosyayı oku
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// **Dosya bilgilerini al**
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya bilgisi alınamadı: %v", err)
	}

	// **Şifreli olup olmadığını kontrol et**
	if fileInfo.Size() < int64(12+kyber512.CiphertextSize) {
		return fmt.Errorf("Bu dosya şifreli değil! Önce dosyayı şifreleyin.")
	}

	// Nonce, şifrelenmiş AES anahtarı ve şifreli içeriği oku
	nonce := make([]byte, 12)
	if _, err := file.Read(nonce); err != nil {
		return err
	}

	ciphertextKey := make([]byte, kyber512.CiphertextSize)
	if _, err := file.Read(ciphertextKey); err != nil {
		return err
	}

	ciphertext, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Kyber512 ile AES anahtarını çöz
	aesKey := make([]byte, kyber512.PlaintextSize)
	privateKey.DecryptTo(aesKey, ciphertextKey)

	// AES-GCM ile şifreyi çöz
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	stream, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := stream.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Çözülen veriyi dosyaya yaz
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.Write(plaintext)
	return err
}
