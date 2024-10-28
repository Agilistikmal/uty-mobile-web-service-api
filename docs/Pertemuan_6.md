# Pertemuan 6 <!-- omit in toc -->
## Data diri

| Nama                | NPM        |
| ------------------- | ---------- |
| Agil Ghani Istikmal | 5220411040 |

---

## Daftar Isi

- [Data diri](#data-diri)
- [Daftar Isi](#daftar-isi)
- [Konversi Figma ke Flutter](#konversi-figma-ke-flutter)


## Konversi Figma ke Flutter

## API OTP

OTP digunakan untuk mengirim kode verifikasi kepada user. Ini untuk membuktikan bahwa yang mengakses adalah user tersebut untuk mengurangi resiko dihack. <br>
Saya menggunakan [WAHA](https://waha.devlike.pro/) untuk mengirimkan OTP melalui **WhatsApp**. WAHA adalah 3rd party API open source yang dapat mengirimkan pesan melalui whatsapp. Saya menghosting sendiri WAHA di VPS yang dapat diakses melalui https://waha.safatanc.com dengan nomor [+6285888881550](https://wa.me/6285888881550) <br>

### Send OTP

Untuk mengirimkan pesan cukup memanggil API dari WAHA. <br>
Melalui endpoint POST /api/sendText <br>

Contohnya POST https://waha.safatanc.com/api/sendText

```json
{
	"session": "default",
	"chatId":  "6281346173829@c.us",
	"text":    "Hi there!"
}
```

<p align="center">
	<img src="./assets/waha1.png" />
</p>

### Send OTP - Golang

Berikut adalah cara mengirim OTP melalui golang. OTP dikirim saat melakukan register dan login. <br>

1. OTP dibuat dengan 4 digit huruf random
2. OTP disimpan dalam tabel OTP dengan expired time 10 menit
3. OTP dikirim melalui whatsapp ke nomor user

#### OTP Model

```go
package model

import "time"

type OTP struct {
	Username  string    `json:"username,omitempty" gorm:"primaryKey"`
	Code      string    `json:"code,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

```

#### OTP Repository

Repository untuk membuat random code dan menyimpannya di database.

```go
package repository

...

func (r *OTPRepository) Create(username string) (*model.OTP, error) {
	code := pkg.RandomString(4)

	otp := &model.OTP{
		Username:  username,
		Code:      code,
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}

	err := r.db.Save(&otp).Error
	if err != nil {
		return nil, err
	}

	return otp, nil
}

```

#### OTP Service

Service untuk logic saat ingin membuat code dan mengirimkannya ke user. <br>
Template message OTP diambil dari file `config.yml`

```go
package service 

...

func (s *OTPService) Generate(username string) (*model.OTP, error) {
	user, err := s.userRepository.Find(username)
	if err != nil {
		return nil, err
	}

	otp, err := s.otpRepository.Create(user.Username)
	if err != nil {
		return nil, err
	}

	var otpMessageBuffer bytes.Buffer
	otpMessageTemplate := template.Must(template.New("otp_message").Parse(viper.GetString("otp.message")))
	otpMessageTemplate.Execute(&otpMessageBuffer, map[string]string{
		"Username": otp.Username,
		"Code":     otp.Code,
	})

	otpMessage := strings.ReplaceAll(otpMessageBuffer.String(), "\n", `\n`)

	body := []byte(fmt.Sprintf(`{
		"session": "default",
		"chatId":  "%s",
		"text":    "%s"
	}`, user.Phone[1:]+"@c.us", otpMessage))

	endpoint := viper.GetString("waha.base_url") + "/api/sendText"

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to send otp")
	}

	return otp, nil
}
```

## API Payment Gateway

## Integrasi Flutter dan API

