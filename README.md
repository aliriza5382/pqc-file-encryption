# Post-Kuantum Dosya Şifreleme Aracı

Bu proje, Go dili ile geliştirilmiş bir dosya şifreleme sistemidir. AES-GCM ve Kyber512 algoritmaları ile klasik ve post-kuantum güvenlik birleştirilmiştir.

## Özellikler

- TXT dosyalarını şifreleme ve çözme
- AES-GCM ile içerik şifreleme
- Kyber512 (Post-Quantum) ile anahtar şifreleme
- Anahtar üretimi ve kaydetme/yükleme
- GUI üzerinden dosya seçimi (`dialog` kütüphanesi)

## Kullanılan Teknolojiler

- Go (Golang)
- `github.com/cloudflare/circl` – Kyber512
- `crypto/aes`, `crypto/cipher` – AES-GCM
- `github.com/sqweek/dialog` – Dosya seçimi arayüzü

## Dosya Yapısı

| Dosya            | Açıklama                                 |
|------------------|------------------------------------------|
| `main.go`        | Kullanıcı arayüzü ve işlem menüsü        |
| `encrypt.go`     | Dosya şifreleme fonksiyonları            |
| `decrypt.go`     | Dosya şifre çözme fonksiyonları          |
| `kyber.go`       | Anahtar üretme, kaydetme, yükleme işleri |
