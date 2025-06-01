# 🎵 Database Lagu Favorit Mahasiswa

Proyek ini adalah aplikasi terminal berbasis bahasa Go (Golang) untuk mengelola data lagu favorit mahasiswa. Aplikasi ini menggunakan struktur data array statis dan menyediakan fitur CRUD (Create, Read, Update, Delete), pencarian, pengurutan, dan perhitungan statistik sederhana.

## 📦 Fitur Utama

- ✅ Tambah lagu baru (dengan ID otomatis)
- 📄 Tampilkan daftar semua lagu dalam bentuk tabel
- ✏️ Update informasi lagu (kecuali ID)
- ❌ Hapus lagu dari database
- 🔍 Cari lagu berdasarkan ID (binary search) atau nama penyanyi (sequential search)
- 🔃 Urutkan lagu berdasarkan berbagai kriteria:
  - ID (Asc/Desc)
  - Judul lagu (A-Z/Z-A)
  - Penyanyi (A-Z/Z-A)
  - Durasi lagu (terpanjang/terpendek)
  - Jumlah pendengar (terbanyak/tersedikit)
- 📊 Hitung rata-rata jumlah pendengar dari seluruh lagu

## 🛠️ Teknologi

- Bahasa: [Go (Golang)](https://golang.org/)
- Pendekatan: CLI berbasis teks
- Algoritma:
  - Insertion Sort & Selection Sort
  - Binary Search & Sequential Search

## 🚀 Cara Menjalankan

1. **Clone repository ini:**

```bash
git clone https://github.com/username/nama-repo.git
cd nama-repo
