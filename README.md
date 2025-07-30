# Psikolog Online Backend

Aplikasi ini merupakan REST API sederhana untuk chatbot psikologi. Backend ditulis dengan Go dan framework Gin, serta menggunakan JWT untuk autentikasi.

## Persyaratan
- Go versi 1.20 atau lebih baru

## Menjalankan Aplikasi
1. Clone repositori ini.
2. Masuk ke folder project.
3. Jalankan perintah berikut:

```bash
go run ./cmd/server
```

Server akan berjalan di `http://localhost:8080`. Endpoint- endpoint dapat diakses melalui prefix `/api`.

Contoh untuk registrasi pengguna:

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"email":"test@example.com","password":"secret"}'
```

Setelah login, sertakan token JWT pada header `Authorization: Bearer <token>` untuk mengakses endpoint lain.

