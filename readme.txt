# PQC File Encryption – Post-Kuantum Dosya Şifreleme Aracı

Bu proje, **Kyber512 (Post-Kuantum Kriptografi)** ve **AES-GCM** kullanarak güvenli bir dosya şifreleme ve çözme sistemi sunar. Geleneksel RSA/ECC algoritmalarına kıyasla **kuantum bilgisayarlara karşı dayanıklıdır**.

---

## Özellikler

- **Kyber512 ile Anahtar Değişimi** (Post-Kuantum Güvenli)
- **AES-256-GCM ile Dosya Şifreleme**
- Şifrelenmiş anahtarları dosya içinde güvenli şekilde saklama
- CLI + GUI desteği (Grafiksel kullanım için `gui.go`)
- Anahtar üretimi, yönetimi ve güvenli dosya çözümü

---

## Proje Dosya Yapısı

| Dosya               | Açıklama |
|---------------------|----------|
| `main.go`           | Konsol üzerinden dosya şifreleme/çözme işlemlerini başlatır. |
| `encrypt.go`        | Dosyayı AES-GCM ile şifreler, AES anahtarını Kyber512 ile şifreler. |
| `decrypt.go`        | AES anahtarını Kyber512 ile çözer, dosya şifresini açar. |
| `kyber.go`          | Kyber512 Public/Private key üretimi ve yönetimi |
| `gui.go`            | Dosya seçme, şifreleme ve çözme işlemleri için basit GUI (tkinter) arayüzü |

---

## Geri Bildirim / İletişim

Her türlü geri bildirim veya öneri için:

E-posta: [sahinaliriza888@gmail.com](mailto:sahinaliriza888@gmail.com)  
GitHub: [github.com/aliriza5382](https://github.com/aliriza5382)
