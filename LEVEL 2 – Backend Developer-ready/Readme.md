## Level 2 – Backend Developer-ready
Selamat datang di Level 2! Di sini kita akan fokus pada pengembangan backend dengan Go, mempersiapkan kamu untuk menjadi developer backend yang siap kerja.

## Review Level
Target: Mampu membuat REST API skala kecil-menengah

Materi
Gin/Fiber/Echo	Web framework cepat
GORM / SQLX / raw SQL	ORM atau no ORM
Middleware (auth, logging, CORS)	Kontrol request
Pagination, search, filter API	Pattern industri
JWT Auth, RBAC	Sistem login modern
Swagger docs / OpenAPI	Dokumentasi otomatis

🧪 Latihan:

Buat API sistem buku tamu atau inventaris sekolah

Tambahkan JWT login + role admin/user


### Tujuan Level 2
Kamu akan belajar framework Go, ORM, middleware, dan cara membuat REST API yang siap digunakan di industri.
### Materi Level 2
#### Framework Go
Kita akan belajar framework Go yang populer seperti Gin, Fiber, atau Echo. Framework ini membantu kita membuat aplikasi web dengan cepat dan efisien.
#### ORM dan Database
Kita akan belajar menggunakan ORM seperti GORM atau SQLX, atau bahkan menggunakan raw SQL untuk interaksi database. Ini penting untuk mengelola data dengan mudah.
#### Middleware
Kita akan belajar cara membuat middleware untuk menangani autentikasi, logging, dan CORS. Middleware ini membantu kita mengontrol request yang masuk ke aplikasi.
#### Pagination, Search, dan Filter API
Kita akan belajar cara membuat API yang mendukung pagination, search, dan filter. Ini adalah pattern umum yang digunakan di industri untuk mengelola data dalam jumlah besar.
#### Autentikasi dan Otorisasi
Kita akan belajar cara membuat sistem autentikasi modern menggunakan JWT (JSON Web Token) dan RBAC (Role-Based Access Control). Ini penting untuk mengamankan API kita.
#### Dokumentasi API
Kita akan belajar cara membuat dokumentasi API otomatis menggunakan Swagger atau OpenAPI. Dokumentasi ini membantu developer lain memahami cara menggunakan API kita.
### Latihan Level 2
Setelah mempelajari materi di atas, kamu akan membuat REST API sederhana yang mencakup semua konsep yang telah dipelajari. Berikut adalah latihan yang harus kamu selesaikan:

```markdown
# Latihan Level 2: REST API Sederhana
## Tujuan
Kamu akan membuat REST API sederhana yang mencakup semua konsep yang telah dipelajari di Level 2. API ini akan digunakan untuk sistem buku tamu atau inventaris sekolah.
## Spesifikasi API
### Endpoints
- `POST /login`: Endpoint untuk login menggunakan JWT
- `GET /books`: Endpoint untuk mendapatkan daftar buku (pagination, search, filter)
- `POST /books`: Endpoint untuk menambahkan buku baru
- `PUT /books/{id}`: Endpoint untuk memperbarui informasi buku
- `DELETE /books/{id}`: Endpoint untuk menghapus buku
- `GET /books/{id}`: Endpoint untuk mendapatkan detail buku
- `GET /inventory`: Endpoint untuk mendapatkan daftar inventaris (pagination, search, filter)
- `POST /inventory`: Endpoint untuk menambahkan item inventaris baru
- `PUT /inventory/{id}`: Endpoint untuk memperbarui informasi inventaris
- `DELETE /inventory/{id}`: Endpoint untuk menghapus item inventaris
- `GET /inventory/{id}`: Endpoint untuk mendapatkan detail item inventaris
### Fitur yang Harus Diimplementasikan
- Gunakan framework Go seperti Gin, Fiber, atau Echo untuk membuat API
- Gunakan ORM seperti GORM atau SQLX, atau raw SQL untuk interaksi database
- Implementasikan middleware untuk autentikasi, logging, dan CORS
- Implementasikan pagination, search, dan filter pada endpoint `GET /books` dan `GET /inventory`
- Implementasikan sistem autentikasi menggunakan JWT
- Implementasikan RBAC untuk membedakan akses antara admin dan user
- Buat dokumentasi API otomatis menggunakan Swagger atau OpenAPI
### Penilaian
- Kode harus bersih, terstruktur, dan mudah dibaca
- Semua endpoint harus berfungsi sesuai spesifikasi
- Middleware harus berfungsi dengan baik
- Autentikasi dan otorisasi harus berfungsi dengan benar
- Dokumentasi API harus lengkap dan mudah dipahami
### Target Level 2
```markdown


