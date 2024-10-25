# Pertemuan 5 

## Data diri

| Nama                | NPM        |
| ------------------- | ---------- |
| Agil Ghani Istikmal | 5220411040 |

---

## Daftar Isi

- [Pertemuan 5](#pertemuan-5)
  - [Data diri](#data-diri)
  - [Daftar Isi](#daftar-isi)
  - [Design Figma](#design-figma)
  - [REST API](#rest-api)
    - [HTTP Methods](#http-methods)
      - [GET](#get)
        - [GET /users](#get-users)
  - [API User Login](#api-user-login)


## Design Figma

## REST API

REST API atau RESTful API yang memiliki kepanjangan **RE**presentational **S**tate **T**ransfer adalah salah satu cara berkomunikasi antar client-server menggunakan protokol HTTP. <br>

Gambaran sederhananya, client mengirim request ke endpoint tertentu, dan dikembalikan response nya saat request telah selesai diproses. <br>

REST menggunakan JSON untuk request body dan response body nya.

<p align="center">
  <img src="./assets/rest.png" />
<p>

### HTTP Methods

Karena REST API menggunakan protokol HTTP, maka ada beberapa protokol HTTP yang sering digunakan untuk membuat RESTful API.

#### GET

Method GET biasanya digunakan untuk menampilkan daftar data atau detail tentang data tersebut.

##### GET /user

Contohnya dengan endpoint GET `http://localhost:3000/user`
Ini digunakan untuk mengambil daftar data user yang ada di database. <br>

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": [
    {
      "username": "agilistikmal",
      "full_name": "Agil Ghani Istikmal",
      "phone": "+6281346173829",
      "created_at": "...",
      "updated_at": "..."
    },
    {
      "username": "ghani",
      "full_name": "Ghani Istikmal Agil",
      "phone": "++628123456789",
      "created_at": "...",
      "updated_at": "..."
    },
  ]
}
```

##### GET /user/:username

Ini untuk menampilkan 1 data detail tentang user dengan username tersebut. Contohnya `GET /user/agilistikmal` <br>

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": {
    "username": "agilistikmal",
    "full_name": "Agil Ghani Istikmal",
    "phone": "+6281346173829",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

#### POST

POST digunakan untuk membuat atau input data baru. POST menyertakan body berupa JSON. <br>

##### POST /user

```json
{
  "username": "agilistikmal",
  "full_name": "Agil Ghani Istikmal",
  "phone": "+6281346173829"
}
```

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": {
    "username": "agilistikmal",
    "full_name": "Agil Ghani Istikmal",
    "phone": "+6281346173829",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

#### PUT

PUT digunakan untuk mengubah dengan menyertakan keseluruhan data baru. <br>

##### PUT /user/:username

Contohnya ingin mengubah data user agilistikmal `PUT /user/agilistikmal`

```json
{
  "username": "agil_baru",
  "full_name": "Agil Punya Nama Baru",
  "phone": "+62812345789"
}
```

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": {
    "username": "agil_baru",
    "full_name": "Agil Punya Nama Baru",
    "phone": "+62812345789",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

#### PATCH

PATCH digunakan untuk mengubah sebagian data saja. <br>

##### PATCH /user/:username

Contohnya ingin mengubah data full_name pada user agilistikmal `PATCH /user/agil_baru`

```json
{
  "full_name": "Agil Saja",
}
```

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": {
    "username": "agil_baru",
    "full_name": "Agil saja",
    "phone": "+62812345789",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

#### DELETE

Method DELETE digunakan untuk menghapus data

##### DELETE /user/:username

`DELETE /user/agil_baru` maka akan menghapus data user dengan username agil_baru

Contoh Response:

```json
{
  "status": 200,
  "message": "ok",
  "data": {
    "username": "agil_baru",
    "full_name": "Agil saja",
    "phone": "+62812345789",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

## API User Login