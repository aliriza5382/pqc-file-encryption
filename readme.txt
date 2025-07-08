*****
Post-Kuantum Şifreleme (PQC) kullanan bir dosya şifreleme ve şifre çözme aracıdır. Ana amacı, Kyber512 algoritması ile güvenli bir anahtar değişimi yaparak AES-GCM ile dosyaları şifrelemek ve çözmektir.

Kodun Genel Mantığı
Anahtar üretme ve yönetme: Kyber512 algoritması ile bir açık anahtar (public key) ve özel anahtar (private key) oluşturulur. Anahtarlar dosyaya kaydedilir ve gerektiğinde yüklenir.
Dosya şifreleme:
AES-256 için rastgele bir şifreleme anahtarı oluşturulur.
Bu anahtar Kyber512 ile şifrelenir ve dosyanın başına eklenir.
AES-GCM şifreleme algoritması kullanılarak dosya içeriği şifrelenir.
Şifreli dosya, şifrelenmiş AES anahtarı ve bir nonce (başlatma vektörü) ile birlikte kaydedilir.
Dosya şifre çözme:
Şifreli dosya açılır ve Kyber512 ile şifrelenmiş AES anahtarı çözülür.
AES-GCM kullanılarak dosyanın şifresi açılır ve orijinal hali geri yüklenir.
Ana Dosya (main.go)
Bu dosya, bir konsol arayüzü sağlayarak kullanıcının:

TXT dosyası şifrelemesine,
Şifrelenmiş bir TXT dosyasını çözmesine,
Çıkış yapmasına olanak tanır.
İşleyiş
Kullanıcıdan bir seçim yapması istenir.
Seçilen dosyanın yolu alınır.
Eğer şifreleme işlemi yapılıyorsa:
Anahtarlar yüklenir veya oluşturulur.
Dosya AES-GCM ile şifrelenir ve Kyber512 ile güvenli hale getirilir.
Eğer şifre çözme işlemi yapılıyorsa:
Özel anahtar yüklenir.
AES anahtarı çözülür.
AES-GCM ile dosyanın şifresi açılır.
Şifre Çözme Dosyası (decrypt.go)
Bu dosya, AES-GCM ile şifrelenmiş bir dosyanın şifresini çözmekten sorumludur.

Şifrelenmiş dosya açılır ve içindeki nonce ve Kyber512 ile şifrelenmiş AES anahtarı okunur.
Özel anahtar (private key) kullanılarak Kyber512 ile AES anahtarı çözülür.
AES-GCM algoritmasıyla şifre çözme işlemi gerçekleştirilir.
Şifre çözülmüş içerik yeni bir dosyaya yazılır.
Önemli Noktalar
AES anahtarı Kyber512 ile şifrelenmiş olarak dosyanın içinde saklanır.
AES anahtarını çözüp geri almak için özel anahtar gereklidir.
AES-GCM nonce’ı kullanarak şifrelenmiş dosyayı açar.
Şifreleme Dosyası (encrypt.go)
Bu dosya, AES-GCM ve Kyber512 kullanarak bir dosyanın şifrelenmesini sağlar.

Dosya açılır ve içeriği okunur.
Rastgele 256-bit (32 byte) bir AES anahtarı oluşturulur.
Bu anahtar, Kyber512 algoritması ile şifrelenir.
AES-GCM kullanarak dosya içeriği şifrelenir.
Şifrelenmiş içerik, nonce ve Kyber512 ile şifrelenmiş anahtarla birlikte dosyaya kaydedilir.
Önemli Noktalar
AES anahtarı doğrudan saklanmaz, Kyber512 ile şifrelenerek güvence altına alınır.
AES-GCM, şifreleme işlemi için nonce (rastgele başlatma vektörü) kullanır.
AES-256, güçlü ve güvenilir bir şifreleme algoritmasıdır.
Kyber512 Anahtar Yönetimi (kyber.go)
Bu dosya, Kyber512 anahtar çifti üretme, kaydetme ve yükleme işlemlerini yönetir.

GenerateKyberKeyPair(): Yeni bir Kyber512 açık ve özel anahtar çifti üretir.
SaveKeys(): Açık ve özel anahtarları dosyaya kaydeder.
LoadKeys(): Daha önce kaydedilmiş anahtarları dosyadan yükler.
Önemli Noktalar
Anahtarlar kaydedilmezse, her açılışta yeni anahtar üretilir.
Özel anahtar gizli tutulmalıdır. Aksi takdirde şifreleme güvenliği tehlikeye girer.
Kyber512, kuantum bilgisayarlara dayanıklı bir algoritmadır, yani geleneksel RSA ve ECC’den daha güvenlidir.
Genel İşleyiş
Kullanıcı bir dosya seçer.
Program anahtarları oluşturur veya yükler.
Dosya şifrelenir ve şifreli versiyonu kaydedilir.
Dosyanın şifresini çözmek için özel anahtar kullanılır ve içeriği eski haline




*****
AES-GCM (Advanced Encryption Standard - Galois/Counter Mode)
AES-GCM, gizlilik ve bütünlük sağlayan modern bir blok şifreleme modudur. AES (Advanced Encryption Standard), simetrik anahtarlı bir şifreleme algoritmasıdır ve GCM (Galois/Counter Mode), AES için bir çalışma modudur.

Özellikleri ve Avantajları:
Gizlilik ve bütünlük sağlar → Hem şifreleme hem de doğrulama işlemi yapar.
Hızlıdır → Paralel işlenebilir ve donanım optimizasyonlarına uygundur.
Yetkilendirilmiş Şifreleme (Authenticated Encryption - AEAD) sunar.
Kimlik doğrulamalı ek veri (AAD - Additional Authenticated Data) desteği vardır.
Düşük gecikme süresi ve yüksek performans sunar.
Çalışma Mantığı:
AES-CTR (Counter Mode) ile veriyi şifreler.
Galois Field (GF(2^128)) tabanlı doğrulama ile veri bütünlüğünü korur.
Nonce (IV) kullanımı: Nonce, tekrar saldırılarını önlemek için her işlemde farklı olmalıdır.
Tag (Etiket) üretimi: Veri değiştirilirse doğrulama başarısız olur.
Uygulama Alanları:
TLS (Transport Layer Security) → HTTPS bağlantılarında kullanılır.
IPsec → VPN ve ağ güvenliğinde.
Donanım hızlandırmalı şifreleme çözümlerinde (Intel AES-NI gibi).
Kyber512 (Post-Kuantum Kriptografi - PQC)
Kyber512, post-kuantum kriptografi (PQC) alanında önerilen lattice (örgü) temelli bir anahtar değişim algoritmasıdır. Geleneksel RSA ve Eliptik Eğri Kriptografisi (ECC) gibi algoritmalar, kuantum bilgisayarlarla kırılabilirken, Kyber bu saldırılara dayanıklıdır.

Özellikleri ve Avantajları:
Post-Kuantum Güvenliği → Shor algoritmasına karşı dayanıklıdır.
Hızlı ve düşük hesaplama maliyetli → Kısa anahtarlarla yüksek güvenlik sağlar.
Anahtar Değişimi için tasarlanmıştır → Kuantum saldırılarına karşı güvenli bir şifreleme yöntemi sunar.
NIST Standartlaşma Süreci → Kyber, NIST’in post-kuantum şifreleme standardizasyon yarışmasını kazanan algoritmalardan biridir.
Çalışma Mantığı:
Lattice (örgü) tabanlı şifreleme → Öğütülmüş Örgü (Learning With Errors - LWE) problemine dayanır.
Kamuya açık anahtar ile şifreleme ve gizli anahtar ile çözme işlemi yapar.
Gürültü ekleme (Noise Injection) → Kuantum bilgisayarların çözmesini zorlaştırır.
Kyber512, Kyber768 ve Kyber1024 gibi farklı güvenlik seviyelerine sahiptir.
Uygulama Alanları:
Post-kuantum güvenli TLS bağlantıları.
Güvenli anahtar değişimi (VPN, mesajlaşma protokolleri).
Kuantum saldırılarına dayanıklı şifreleme çözümleri.
AES-GCM vs Kyber512 Karşılaştırması
Özellik	AES-GCM	Kyber512
Şifreleme Türü	Simetrik	Asimetrik
Kullanım Amacı	Veri şifreleme & doğrulama	Anahtar değişimi
Post-Kuantum Güvenliği	Kuantum saldırılarına karşı savunmasız	Kuantum saldırılarına dayanıklı
Hız & Performans	Çok hızlı, donanım hızlandırmalı	Hızlı ancak asimetrik olduğu için AES’den biraz daha yavaş
Anahtar Boyutu	128-bit, 192-bit, 256-bit	512-bit (Kyber512), 768-bit, 1024-bit seçenekleri
Uygulama Alanları	TLS, IPsec, VPN, IoT	Post-kuantum güvenli TLS, VPN, mesajlaşma
Özetle:

AES-GCM günümüzde en çok kullanılan simetrik şifreleme modlarından biri olup, veri şifreleme ve doğrulama için uygundur.
Kyber512, post-kuantum güvenli bir anahtar değişim protokolüdür ve kuantum bilgisayarlara dayanıklıdır.

Projenin Genel Yapısı
Ana Uygulama (main.go)

Kullanıcıdan bir seçim yapmasını isteyen bir konsol uygulaması içerir.
Kullanıcı 1 tuşuna basarsa bir TXT dosyası seçip şifreleyebilir.
Kullanıcı 2 tuşuna basarsa bir şifreli TXT dosyasını çözebilir.
Kullanıcı 3 tuşuna basarsa programdan çıkış yapar.
Şifreleme Modülü (encrypt.go)

AES-GCM algoritmasını kullanarak dosya içeriğini şifreler.
Kyber512 algoritması ile AES anahtarını güvenli hale getirir.
Kyber512'nin Public Key’i ile AES anahtarını şifreler ve bunu şifrelenmiş dosyaya ekler.
Şifre Çözme Modülü (decrypt.go)

AES anahtarını Kyber512 Private Key ile çözer.
AES-GCM algoritmasını kullanarak şifrelenmiş dosya içeriğini çözer ve orijinal hale getirir.
Kyber512 Anahtar Yönetimi (kyber.go)

Kyber512 Public ve Private Key üretir.
Anahtarları bir dosyaya kaydeder ve yükler (public.key, private.key).
Grafik Arayüz (gui.go)

Kullanıcının dosya seçmesine, şifrelemesine ve şifre çözmesine olanak tanıyan bir arayüz (GUI) sağlar.
Projenin Çalışma Mantığı
Kullanıcı bir dosya şifrelemek istediğinde:

Bir TXT dosyası seçilir.
Eğer daha önce bir anahtar üretilmemişse, Kyber512 Public ve Private Key üretilir.
AES-GCM ile dosya şifrelenir.
AES anahtarı Kyber512 Public Key ile şifrelenerek dosyaya eklenir.
Şifreli dosya oluşturulur.
Kullanıcı şifre çözmek istediğinde:

Şifreli dosya seçilir.
AES anahtarı Kyber512 Private Key ile çözülür.
AES-GCM algoritması kullanılarak dosyanın içeriği eski haline getirilir.
Çözülmüş dosya oluşturulur.