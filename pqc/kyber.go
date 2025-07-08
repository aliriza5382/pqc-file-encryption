package pqc

import (
	cryptoRand "crypto/rand"
	"fmt"
	"os"

	"github.com/cloudflare/circl/pke/kyber/kyber512"
)

// **Kyber512 Anahtar Çifti Üret**
// Bu fonksiyon Kyber512 algoritmasını kullanarak bir açık anahtar (public key)
// ve özel anahtar (private key) çifti üretir.
func GenerateKyberKeyPair() (*kyber512.PublicKey, *kyber512.PrivateKey, error) {
	// Kyber512 anahtar çifti oluşturulur
	publicKey, privateKey, err := kyber512.GenerateKey(cryptoRand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("Anahtar üretme hatası: %v", err)
	}
	return publicKey, privateKey, nil
}

// **Anahtarları Dosyaya Kaydet**
// Bu fonksiyon üretilen açık ve özel anahtarları dosyaya kaydeder.
func SaveKeys(publicKey *kyber512.PublicKey, privateKey *kyber512.PrivateKey) error {
	// Açık anahtar için byte dizisi oluşturulur
	pkData := make([]byte, kyber512.PublicKeySize)
	publicKey.Pack(pkData)

	// Özel anahtar için byte dizisi oluşturulur
	skData := make([]byte, kyber512.PrivateKeySize)
	privateKey.Pack(skData)

	// Açık anahtar "public.key" dosyasına yazılır
	err := os.WriteFile("public.key", pkData, 0644)
	if err != nil {
		return err
	}

	// Özel anahtar "private.key" dosyasına yazılır
	err = os.WriteFile("private.key", skData, 0644)
	return err
}

// **Anahtarları Dosyadan Yükle**
// Bu fonksiyon daha önce kaydedilmiş açık ve özel anahtarları dosyadan yükler.
func LoadKeys() (*kyber512.PublicKey, *kyber512.PrivateKey, error) {
	// "public.key" dosyası okunur
	pkData, err := os.ReadFile("public.key")
	if err != nil {
		return nil, nil, err
	}

	// "private.key" dosyası okunur
	skData, err := os.ReadFile("private.key")
	if err != nil {
		return nil, nil, err
	}

	// Anahtar nesneleri oluşturulur
	publicKey := new(kyber512.PublicKey)
	privateKey := new(kyber512.PrivateKey)

	// Okunan verilerden anahtarlar çıkarılır
	publicKey.Unpack(pkData)
	privateKey.Unpack(skData)

	return publicKey, privateKey, nil
}
