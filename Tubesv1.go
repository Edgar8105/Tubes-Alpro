package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Song struct {
	ID              string
	Judul           string
	Penyanyi        string
	Durasi          float64
	JumlahPendengar int
}

const NMAXLagu = 100

var daftarLagu [NMAXLagu]Song
var jumlahLaguTerisi int

// initDummyData menginisialisasi daftar lagu dengan 20 data contoh awal secara langsung ke array statis.
func initDummyData() {
	// Data Dummy Asli (15 lagu)
	if NMAXLagu >= 1 {
		daftarLagu[0] = Song{ID: "L001", Judul: "Monokrom", Penyanyi: "Tulus", Durasi: 3.5, JumlahPendengar: 35}
	}
	if NMAXLagu >= 2 {
		daftarLagu[1] = Song{ID: "L002", Judul: "Sial", Penyanyi: "Mahalini", Durasi: 4.1, JumlahPendengar: 42}
	}
	if NMAXLagu >= 3 {
		daftarLagu[2] = Song{ID: "L003", Judul: "As It Was", Penyanyi: "Harry Styles", Durasi: 2.8, JumlahPendengar: 55}
	}
	if NMAXLagu >= 4 {
		daftarLagu[3] = Song{ID: "L004", Judul: "Bertaut", Penyanyi: "Nadin Amizah", Durasi: 4.3, JumlahPendengar: 28}
	}
	if NMAXLagu >= 5 {
		daftarLagu[4] = Song{ID: "L005", Judul: "Glimpse of Us", Penyanyi: "Joji", Durasi: 3.9, JumlahPendengar: 48}
	}
	if NMAXLagu >= 6 {
		daftarLagu[5] = Song{ID: "L006", Judul: "Hati-Hati di Jalan", Penyanyi: "Tulus", Durasi: 4.0, JumlahPendengar: 32}
	}
	if NMAXLagu >= 7 {
		daftarLagu[6] = Song{ID: "L007", Judul: "Until I Found You", Penyanyi: "Stephen Sanchez", Durasi: 3.0, JumlahPendengar: 51}
	}
	if NMAXLagu >= 8 {
		daftarLagu[7] = Song{ID: "L008", Judul: "Komang", Penyanyi: "Raim Laode", Durasi: 3.7, JumlahPendengar: 39}
	}
	if NMAXLagu >= 9 {
		daftarLagu[8] = Song{ID: "L009", Judul: "Anti-Hero", Penyanyi: "Taylor Swift", Durasi: 3.3, JumlahPendengar: 60}
	}
	if NMAXLagu >= 10 {
		daftarLagu[9] = Song{ID: "L010", Judul: "Tak Ingin Usai", Penyanyi: "Keisya Levronka", Durasi: 4.6, JumlahPendengar: 25}
	}
	if NMAXLagu >= 11 {
		daftarLagu[10] = Song{ID: "L011", Judul: "Dandelions", Penyanyi: "Ruth B.", Durasi: 3.8, JumlahPendengar: 45}
	}
	if NMAXLagu >= 12 {
		daftarLagu[11] = Song{ID: "L012", Judul: "Runtuh", Penyanyi: "Feby Putri ft. Fiersa Besari", Durasi: 4.2, JumlahPendengar: 31}
	}
	if NMAXLagu >= 13 {
		daftarLagu[12] = Song{ID: "L013", Judul: "Kill Bill", Penyanyi: "SZA", Durasi: 2.6, JumlahPendengar: 58}
	}
	if NMAXLagu >= 14 {
		daftarLagu[13] = Song{ID: "L014", Judul: "To The Bone", Penyanyi: "Pamungkas", Durasi: 5.8, JumlahPendengar: 29}
	}
	if NMAXLagu >= 15 {
		daftarLagu[14] = Song{ID: "L015", Judul: "Here With Me", Penyanyi: "d4vd", Durasi: 4.0, JumlahPendengar: 47}
	}
	if NMAXLagu >= 16 {
		daftarLagu[15] = Song{ID: "L016", Judul: "Manusia Kuat", Penyanyi: "Tulus", Durasi: 3.2, JumlahPendengar: 40} // Lagu ke-3 Tulus
	}
	if NMAXLagu >= 17 {
		daftarLagu[16] = Song{ID: "L017", Judul: "Sisa Rasa", Penyanyi: "Mahalini", Durasi: 4.5, JumlahPendengar: 38} // Lagu ke-2 Mahalini
	}
	if NMAXLagu >= 18 {
		daftarLagu[17] = Song{ID: "L018", Judul: "Daylight", Penyanyi: "Joji", Durasi: 2.9, JumlahPendengar: 52} // Lagu ke-2 Joji
	}
	if NMAXLagu >= 19 {
		daftarLagu[18] = Song{ID: "L019", Judul: "Traitor", Penyanyi: "Olivia Rodrigo", Durasi: 3.8, JumlahPendengar: 49}
	}
	if NMAXLagu >= 20 {
		daftarLagu[19] = Song{ID: "L020", Judul: "Golden Hour", Penyanyi: "JVKE", Durasi: 3.5, JumlahPendengar: 53}
	}

	if NMAXLagu < 20 {
		jumlahLaguTerisi = NMAXLagu
		if NMAXLagu > 0 {
			fmt.Printf("Peringatan: NMAXLagu (%d) lebih kecil dari jumlah data dummy (20). Hanya %d data yang diinisialisasi.\n", NMAXLagu, NMAXLagu)
		} else if NMAXLagu == 0 {
			fmt.Println("Peringatan: NMAXLagu adalah 0. Tidak ada data dummy yang diinisialisasi.")
		}
	} else {
		jumlahLaguTerisi = 20
	}

	if jumlahLaguTerisi > 0 {
		insertionSortByIDAsc()
	}
}

// tambahLagu menangani proses penambahan lagu baru ke dalam daftarLagu.
func tambahLagu() {
	var laguBaru Song
	var input string
	var errRead error
	var parsedDurasi float64
	var errScanFloat error

	if jumlahLaguTerisi >= NMAXLagu {
		fmt.Println("Kapasitas database penuh. Tidak bisa menambah lagu.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	laguBaru.ID = generateNewID()
	laguBaru.JumlahPendengar = 1

	fmt.Printf("ID Lagu Baru (Otomatis): %-28s \n", laguBaru.ID)

	var judulValid bool = false
	for !judulValid {
		fmt.Print("Judul Lagu: ")
		input, errRead = reader.ReadString('\n')
		if errRead == nil {
			laguBaru.Judul = strings.TrimSpace(input)
			if laguBaru.Judul == "" {
				fmt.Println("Judul Lagu tidak boleh kosong. Silakan coba lagi.")
			} else {
				judulValid = true
			}
		} else {
			fmt.Println(" Kesalahan membaca input judul:", errRead)
		}
	}

	var penyanyiValid bool = false
	for !penyanyiValid {
		fmt.Print("Penyanyi: ")
		input, errRead = reader.ReadString('\n')
		if errRead == nil {
			laguBaru.Penyanyi = strings.TrimSpace(input)
			if laguBaru.Penyanyi == "" {
				fmt.Println("Nama Penyanyi tidak boleh kosong. Silakan coba lagi.")
			} else {
				penyanyiValid = true
			}
		} else {
			fmt.Println("Kesalahan membaca input penyanyi:", errRead)
		}
	}

	var durasiValid bool = false
	for !durasiValid {
		fmt.Print("Durasi (menit, contoh: 3.5): ")
		input, errRead = reader.ReadString('\n')
		if errRead == nil {
			input = strings.TrimSpace(input)
			parsedDurasi, errScanFloat = strconv.ParseFloat(input, 64)
			if errScanFloat == nil && parsedDurasi > 0 {
				laguBaru.Durasi = parsedDurasi
				durasiValid = true
			} else {
				fmt.Println("Durasi tidak valid. Harap masukkan angka positif.")
			}
		} else {
			fmt.Println("Kesalahan membaca input durasi:", errRead)
		}
	}

	daftarLagu[jumlahLaguTerisi] = laguBaru
	jumlahLaguTerisi++
	insertionSortByIDAsc()
	fmt.Println("Lagu berhasil ditambahkan!")
}

// tampilkanDataLagu menampilkan semua lagu yang ada dalam daftarLagu dengan format tabel.
func tampilkanDataLagu(data [NMAXLagu]Song, count int) {
	idWidth := 4
	judulWidth := 30
	penyanyiWidth := 38
	durasiWidth := 6
	pendengarWidth := 9

	headerLine := fmt.Sprintf("+-%s-+-%s-+-%s-+-%s-+-%s-+",
		strings.Repeat("-", idWidth),
		strings.Repeat("-", judulWidth),
		strings.Repeat("-", penyanyiWidth),
		strings.Repeat("-", durasiWidth),
		strings.Repeat("-", pendengarWidth))
	fmt.Println(headerLine)

	fmt.Printf("| %-*s | %-*s | %-*s | %*s | %*s |\n",
		idWidth, "ID",
		judulWidth, "JUDUL",
		penyanyiWidth, "PENYANYI",
		durasiWidth, "DURASI",
		pendengarWidth, "PENDENGAR")
	fmt.Println(headerLine)

	if count == 0 {
		message := "Belum ada lagu untuk ditampilkan."
		effectiveTableWidth := len(headerLine) - 2
		paddingLen := (effectiveTableWidth - len(message)) / 2
		if paddingLen < 0 {
			paddingLen = 0
		}
		rightPaddingLen := effectiveTableWidth - len(message) - paddingLen
		if rightPaddingLen < 0 {
			rightPaddingLen = 0
		}

		fmt.Printf("|%*s%s%*s|\n", paddingLen, "", message, rightPaddingLen, "")
		fmt.Println(headerLine)
		return
	}

	var displayJudul, displayPenyanyi string
	for i := 0; i < count; i++ {
		lagu := data[i]
		displayJudul = truncateString(lagu.Judul, judulWidth)
		displayPenyanyi = truncateString(lagu.Penyanyi, penyanyiWidth)

		fmt.Printf("| %-*s | %-*s | %-*s | %*.1f | %*d |\n",
			idWidth, lagu.ID,
			judulWidth, displayJudul,
			penyanyiWidth, displayPenyanyi,
			durasiWidth, lagu.Durasi,
			pendengarWidth, lagu.JumlahPendengar)
	}
	fmt.Println(headerLine)
}

// updateLagu menangani proses pembaruan data lagu yang sudah ada berdasarkan ID.
func updateLagu() {
	var idCari, input, newJudul, newPenyanyi string
	var index int
	var found bool
	var errRead, errScanFloat, errScanInt error
	var parsedDurasi float64
	var newJumlahPendengar int

	if jumlahLaguTerisi == 0 {
		fmt.Println("Database lagu kosong. Tidak ada data untuk diupdate.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Lagu yang akan diupdate: ")
	input, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca ID:", errRead)
		return
	}
	idCari = strings.TrimSpace(input)

	index, found = binarySearchByID(idCari)
	if !found {
		fmt.Printf("Lagu dengan ID '%s' tidak ditemukan.\n", idCari)
		return
	}

	fmt.Printf("Judul Baru (Lama: %s): ", truncateString(daftarLagu[index].Judul, 20))
	input, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca judul baru:", errRead)
		return
	}
	newJudul = strings.TrimSpace(input)
	if newJudul != "" {
		daftarLagu[index].Judul = newJudul
	}

	fmt.Printf("Penyanyi Baru (Lama: %s): ", truncateString(daftarLagu[index].Penyanyi, 20))
	input, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca penyanyi baru:", errRead)
		return
	}
	newPenyanyi = strings.TrimSpace(input)
	if newPenyanyi != "" {
		daftarLagu[index].Penyanyi = newPenyanyi
	}

	var durasiUpdateSelesai bool = false
	for !durasiUpdateSelesai {
		fmt.Printf("Durasi Baru (Lama: %.1f): ", daftarLagu[index].Durasi)
		input, errRead = reader.ReadString('\n')
		if errRead == nil {
			input = strings.TrimSpace(input)
			if input == "" {
				durasiUpdateSelesai = true
			} else {
				parsedDurasi, errScanFloat = strconv.ParseFloat(input, 64)
				if errScanFloat == nil && parsedDurasi > 0 {
					daftarLagu[index].Durasi = parsedDurasi
					durasiUpdateSelesai = true
				} else {
					fmt.Println("Durasi tidak valid. Harap masukkan angka positif.  ")
				}
			}
		} else {
			fmt.Println("Kesalahan membaca durasi baru:", errRead)
		}
	}

	var pendengarUpdateSelesai bool = false
	for !pendengarUpdateSelesai {
		fmt.Printf("Jumlah Pendengar Baru (Lama: %d): ", daftarLagu[index].JumlahPendengar)
		input, errRead = reader.ReadString('\n')
		if errRead == nil {
			input = strings.TrimSpace(input)
			if input == "" {
				pendengarUpdateSelesai = true
			} else {
				newJumlahPendengar, errScanInt = strconv.Atoi(input)
				if errScanInt == nil && newJumlahPendengar >= 0 {
					daftarLagu[index].JumlahPendengar = newJumlahPendengar
					pendengarUpdateSelesai = true
				} else {
					fmt.Println("Jumlah pendengar tidak valid. Harap masukkan angka non-negatif. ")
				}
			}
		} else {
			fmt.Println("Kesalahan membaca jumlah pendengar baru:", errRead)
		}
	}

	fmt.Println("Data lagu berhasil diupdate!")
}

// menuHapusLagu menampilkan submenu untuk opsi penghapusan lagu.
func menuHapusLagu(reader *bufio.Reader) {
	var pilihanSubMenu int
	var inputLine string
	var errRead, errConv error

	fmt.Println("| --- Hapus Lagu ---                                    |")
	fmt.Println("| 1. Hapus Satu Lagu Berdasarkan ID                     |")
	fmt.Println("| 2. Hapus Seluruh Data Lagu                            |")
	fmt.Println("| 3. Kembali ke Menu Utama                              |")
	fmt.Println("=========================================================")
	fmt.Print("| Pilih sub-menu (1-3): ")

	inputLine, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca pilihan sub-menu:", errRead)
		return
	}
	pilihanSubMenu, errConv = strconv.Atoi(strings.TrimSpace(inputLine))

	if errConv != nil {
		fmt.Println("Pilihan tidak valid. Harap masukkan angka.")
		return
	}

	if pilihanSubMenu == 3 {
		fmt.Println("Kembali ke Menu Utama...")
		return
	}

	switch pilihanSubMenu {
	case 1:
		prosesHapusSatuLagu(reader)
	case 2:
		prosesHapusSemuaLagu(reader)
	default:
		fmt.Println("Pilihan sub-menu tidak valid.")
	}
}

// prosesHapusSatuLagu menangani logika untuk menghapus satu lagu berdasarkan ID.
func prosesHapusSatuLagu(reader *bufio.Reader) {
	var idCari, konfirmasi string
	var index int
	var found bool
	var errRead error

	if jumlahLaguTerisi == 0 {
		fmt.Println("Database lagu kosong. Tidak ada data untuk dihapus.")
		return
	}

	fmt.Print("Masukkan ID Lagu yang akan dihapus: ")
	idCari, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca ID:", errRead)
		return
	}
	idCari = strings.TrimSpace(idCari)

	index, found = findLaguByID(idCari)
	if !found {
		fmt.Printf("Lagu dengan ID '%s' tidak ditemukan.\n", idCari)
		return
	}

	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("| Lagu yang akan dihapus:                               |")
	fmt.Printf("| ID         : %-40s |\n", daftarLagu[index].ID)
	fmt.Printf("| Judul      : %-40s |\n", truncateString(daftarLagu[index].Judul, 40))
	fmt.Printf("| Penyanyi   : %-40s |\n", truncateString(daftarLagu[index].Penyanyi, 40))
	fmt.Printf("| Durasi     : %-40.1f |\n", daftarLagu[index].Durasi)
	fmt.Printf("| Pendengar  : %-40d |\n", daftarLagu[index].JumlahPendengar)
	fmt.Println("|-------------------------------------------------------|")

	fmt.Print("Apakah Anda yakin ingin menghapus lagu ini? (Y/N): ")
	konfirmasi, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca konfirmasi:", errRead)
		return
	}
	konfirmasi = strings.TrimSpace(strings.ToUpper(konfirmasi))

	if konfirmasi == "Y" {
		for i := index; i < jumlahLaguTerisi-1; i++ {
			daftarLagu[i] = daftarLagu[i+1]
		}
		jumlahLaguTerisi--
		if jumlahLaguTerisi >= 0 && jumlahLaguTerisi < NMAXLagu {
			daftarLagu[jumlahLaguTerisi] = Song{}
		}
		if jumlahLaguTerisi > 0 {
			insertionSortByIDAsc()
		}
		fmt.Println("Lagu berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

// prosesHapusSemuaLagu menangani logika untuk menghapus semua lagu dalam daftar.
func prosesHapusSemuaLagu(reader *bufio.Reader) {
	var konfirmasi string
	var errRead error

	if jumlahLaguTerisi == 0 {
		fmt.Println("Database lagu sudah kosong. Tidak ada data untuk dihapus.")
		return
	}

	fmt.Print("Apakah Anda yakin ingin menghapus SELURUH data lagu? (Y/N): ")
	konfirmasi, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca konfirmasi:", errRead)
		return
	}
	konfirmasi = strings.TrimSpace(strings.ToUpper(konfirmasi))

	if konfirmasi == "Y" {
		for i := 0; i < jumlahLaguTerisi; i++ {
			daftarLagu[i] = Song{}
		}
		jumlahLaguTerisi = 0
		fmt.Println("Seluruh data lagu berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan seluruh data dibatalkan.")
	}
}

// menuCariLagu menampilkan sub-menu untuk fitur pencarian lagu dan memproses pilihan pengguna.
func menuCariLagu(reader *bufio.Reader) {
	var sub int
	var errConv, errRead error
	var inputLine, inputTrimmed, id, nama string
	var index int
	var found bool
	var hasilPencarian [NMAXLagu]Song
	var jumlahHasil int

	fmt.Println("| --- Cari Lagu ---                                     |")
	fmt.Println("| 1. Cari berdasarkan ID (Binary Search)                |")
	fmt.Println("| 2. Cari berdasarkan Nama Penyanyi (Sequential Search) |")
	fmt.Println("| 3. Kembali ke Menu Utama                              |")
	fmt.Println("=========================================================")
	fmt.Print("Pilih sub-menu (1-3): ")

	inputLine, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca sub-menu:", errRead)
		return
	}
	inputTrimmed = strings.TrimSpace(inputLine)
	sub, errConv = strconv.Atoi(inputTrimmed)

	if errConv != nil {
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama.")
		return
	}

	if sub == 3 {
		fmt.Println("Kembali ke Menu Utama...")
		return
	}

	switch sub {
	case 1:
		fmt.Print("Masukkan ID lagu: ")
		id, errRead = reader.ReadString('\n')
		if errRead != nil {
			fmt.Println("Kesalahan membaca ID:", errRead)
			return
		}
		id = strings.TrimSpace(id)
		index, found = binarySearchByID(id)
		if found {
			var tempArrayForDisplay [NMAXLagu]Song
			tempArrayForDisplay[0] = daftarLagu[index]
			tampilkanDataLagu(tempArrayForDisplay, 1)
		} else {
			fmt.Println("Lagu tidak ditemukan.")
		}
	case 2:
		fmt.Print("Masukkan nama penyanyi (boleh sebagian): ")
		nama, errRead = reader.ReadString('\n')
		if errRead != nil {
			fmt.Println("Kesalahan membaca nama penyanyi:", errRead)
			return
		}
		nama = strings.TrimSpace(nama)
		jumlahHasil = sequentialSearchByPenyanyi(nama, &hasilPencarian)
		if jumlahHasil == 0 {
			fmt.Println("Tidak ditemukan lagu dari penyanyi tersebut.")
		} else {
			tampilkanDataLagu(hasilPencarian, jumlahHasil)
		}
	default:
		fmt.Println("Pilihan sub-menu tidak valid.")
	}
}

// menuUrutkanLagu menampilkan sub-menu untuk fitur pengurutan lagu dan memproses pilihan pengguna.
func menuUrutkanLagu(reader *bufio.Reader) {
	var sub int
	var errConv, errRead error
	var inputLine, inputTrimmed string

	fmt.Println("| --- Urutkan Lagu ---                                  |")
	fmt.Println("| 1. Berdasarkan ID (Ascending - Insertion Sort)        |")
	fmt.Println("| 2. Berdasarkan ID (Descending - Selection Sort)       |")
	fmt.Println("| 3. Berdasarkan Judul (A-Z - Insertion Sort)           |")
	fmt.Println("| 4. Berdasarkan Judul (Z-A - Selection Sort)           |")
	fmt.Println("| 5. Berdasarkan Penyanyi (A-Z - Insertion Sort)        |")
	fmt.Println("| 6. Berdasarkan Penyanyi (Z-A - Selection Sort)        |")
	fmt.Println("| 7. Berdasarkan Durasi (Terpanjang - Selection Sort)   |")
	fmt.Println("| 8. Berdasarkan Durasi (Terpendek - Insertion Sort)    |")
	fmt.Println("| 9. Berdasarkan Pendengar (Terbanyak - Selection Sort) |")
	fmt.Println("|10. Berdasarkan Pendengar (Tersedikit - Insertion Sort)|")
	fmt.Println("|11. Kembali ke Menu Utama                              |")
	fmt.Println("=========================================================")
	fmt.Print("Pilih sub-menu (1–11): ")

	inputLine, errRead = reader.ReadString('\n')
	if errRead != nil {
		fmt.Println("Kesalahan membaca sub-menu:", errRead)
		return
	}
	inputTrimmed = strings.TrimSpace(inputLine)
	sub, errConv = strconv.Atoi(inputTrimmed)

	if errConv != nil {
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama.")
		return
	}
	if sub == 11 {
		fmt.Println("Kembali ke Menu Utama...")
		return
	}

	sorted := true
	switch sub {
	case 1:
		insertionSortByIDAsc()
	case 2:
		selectionSortByIDDesc()
	case 3:
		insertionSortByJudulAsc()
	case 4:
		selectionSortByJudulDesc()
	case 5:
		insertionSortByPenyanyiAsc()
	case 6:
		selectionSortByPenyanyiDesc()
	case 7:
		selectionSortByDurasiDesc()
	case 8:
		insertionSortByDurasiAsc()
	case 9:
		selectionSortByPendengarDesc()
	case 10:
		insertionSortByPendengarAsc()
	default:
		fmt.Println("Pilihan sub-menu tidak valid.")
		sorted = false
	}

	if sorted {
		tampilkanDataLagu(daftarLagu, jumlahLaguTerisi)
	}
}

// hitungRataRataPendengar menghitung dan mengembalikan rata-rata jumlah pendengar dari semua lagu.
func hitungRataRataPendengar() float64 {
	if jumlahLaguTerisi == 0 {
		return 0.0
	}
	totalPendengar := 0
	for i := 0; i < jumlahLaguTerisi; i++ {
		totalPendengar += daftarLagu[i].JumlahPendengar
	}
	return float64(totalPendengar) / float64(jumlahLaguTerisi)
}

// generateNewID menghasilkan ID unik baru untuk lagu berikutnya berdasarkan ID numerik tertinggi yang ada.
func generateNewID() string {
	if jumlahLaguTerisi == 0 {
		return "L001"
	}
	maxNum := 0
	for i := 0; i < jumlahLaguTerisi; i++ {
		s := daftarLagu[i]
		if strings.HasPrefix(s.ID, "L") {
			currentNumStr := strings.TrimPrefix(s.ID, "L")
			currentNum, errParse := strconv.Atoi(currentNumStr)
			if errParse == nil && currentNum > maxNum {
				maxNum = currentNum
			}
		}
	}
	return fmt.Sprintf("L%03d", maxNum+1)
}

// findLaguByID mencari lagu dalam daftarLagu berdasarkan ID secara sequential search.
func findLaguByID(id string) (int, bool) {
	for i := 0; i < jumlahLaguTerisi; i++ {
		if daftarLagu[i].ID == id {
			return i, true
		}
	}
	return -1, false
}

// binarySearchByID mencari lagu dalam daftarLagu berdasarkan ID menggunakan binary search.
func binarySearchByID(id string) (int, bool) {
	low := 0
	high := jumlahLaguTerisi - 1
	var mid int
	var ditemukan bool = false
	var hasilIndex int = -1

	for low <= high && !ditemukan {
		mid = low + (high-low)/2
		if daftarLagu[mid].ID == id {
			ditemukan = true
			hasilIndex = mid
		} else if daftarLagu[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return hasilIndex, ditemukan
}

// sequentialSearchByPenyanyi mencari lagu berdasarkan nama penyanyi secara sequential search.
func sequentialSearchByPenyanyi(nama string, hasilPencarian *[NMAXLagu]Song) int {
	namaLower := strings.ToLower(nama)
	count := 0
	for i := 0; i < jumlahLaguTerisi && count < NMAXLagu; i++ {
		if strings.Contains(strings.ToLower(daftarLagu[i].Penyanyi), namaLower) {
			hasilPencarian[count] = daftarLagu[i]
			count++
		}
	}
	return count
}

// insertionSortByIDAsc mengurutkan daftarLagu berdasarkan ID secara menaik (ascending) menggunakan insertion sort.
func insertionSortByIDAsc() {
	var i, j int
	var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]
		j = i - 1
		for j >= 0 && daftarLagu[j].ID > key.ID {
			daftarLagu[j+1] = daftarLagu[j]
			j--
		}
		daftarLagu[j+1] = key
	}
}

// selectionSortByIDDesc mengurutkan daftarLagu berdasarkan ID secara menurun (descending) menggunakan selection sort.
func selectionSortByIDDesc() {
	var i, j, idxExtreme int
	n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ {
			if daftarLagu[j].ID > daftarLagu[idxExtreme].ID {
				idxExtreme = j
			}
		}
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// insertionSortByJudulAsc mengurutkan daftarLagu berdasarkan judul secara menaik (A-Z, case-insensitive) menggunakan insertion sort.
func insertionSortByJudulAsc() {
	var i, j int
	var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]
		j = i - 1
		for j >= 0 && strings.ToLower(daftarLagu[j].Judul) > strings.ToLower(key.Judul) {
			daftarLagu[j+1] = daftarLagu[j]
			j--
		}
		daftarLagu[j+1] = key
	}
}

// selectionSortByJudulDesc mengurutkan daftarLagu berdasarkan judul secara menurun (Z-A, case-insensitive) menggunakan selection sort.
func selectionSortByJudulDesc() {
	var i, j, idxExtreme int
	n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ {
			if strings.ToLower(daftarLagu[j].Judul) > strings.ToLower(daftarLagu[idxExtreme].Judul) {
				idxExtreme = j
			}
		}
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// insertionSortByPenyanyiAsc mengurutkan daftarLagu berdasarkan nama penyanyi secara menaik (A-Z, case-insensitive) menggunakan insertion sort.
func insertionSortByPenyanyiAsc() {
	var i, j int
	var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]
		j = i - 1
		for j >= 0 && strings.ToLower(daftarLagu[j].Penyanyi) > strings.ToLower(key.Penyanyi) {
			daftarLagu[j+1] = daftarLagu[j]
			j--
		}
		daftarLagu[j+1] = key
	}
}

// selectionSortByPenyanyiDesc mengurutkan daftarLagu berdasarkan nama penyanyi secara menurun (Z-A, case-insensitive) menggunakan selection sort.
func selectionSortByPenyanyiDesc() {
	var i, j, idxExtreme int
	n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ {
			if strings.ToLower(daftarLagu[j].Penyanyi) > strings.ToLower(daftarLagu[idxExtreme].Penyanyi) {
				idxExtreme = j
			}
		}
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// insertionSortByDurasiAsc mengurutkan daftarLagu berdasarkan durasi secara menaik (terpendek) menggunakan insertion sort.
func insertionSortByDurasiAsc() {
	var i, j int
	var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]
		j = i - 1
		for j >= 0 && daftarLagu[j].Durasi > key.Durasi {
			daftarLagu[j+1] = daftarLagu[j]
			j--
		}
		daftarLagu[j+1] = key
	}
}

// selectionSortByDurasiDesc mengurutkan daftarLagu berdasarkan durasi secara menurun (terpanjang) menggunakan selection sort.
func selectionSortByDurasiDesc() {
	var i, j, idxExtreme int
	n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ {
			if daftarLagu[j].Durasi > daftarLagu[idxExtreme].Durasi {
				idxExtreme = j
			}
		}
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// insertionSortByPendengarAsc mengurutkan daftarLagu berdasarkan jumlah pendengar secara menaik (tersedikit) menggunakan insertion sort.
func insertionSortByPendengarAsc() {
	var i, j int
	var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]
		j = i - 1
		for j >= 0 && daftarLagu[j].JumlahPendengar > key.JumlahPendengar {
			daftarLagu[j+1] = daftarLagu[j]
			j--
		}
		daftarLagu[j+1] = key
	}
}

// selectionSortByPendengarDesc mengurutkan daftarLagu berdasarkan jumlah pendengar secara menurun (terbanyak) menggunakan selection sort.
func selectionSortByPendengarDesc() {
	var i, j, idxExtreme int
	n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ {
			if daftarLagu[j].JumlahPendengar > daftarLagu[idxExtreme].JumlahPendengar {
				idxExtreme = j
			}
		}
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// truncateString memotong string s jika panjangnya melebihi maxLength dan menambahkan "...".
func truncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		if maxLength-3 > 0 {
			return s[:maxLength-3] + "..."
		}
		return s[:maxLength]
	}
	return s
}

// tampilkanMenu menampilkan opsi menu utama kepada pengguna.
func tampilkanMenu() {
	fmt.Println("| 1. Tambah Lagu Baru                                   |")
	fmt.Println("| 2. Tampilkan Semua Lagu                               |")
	fmt.Println("| 3. Update Data Lagu                                   |")
	fmt.Println("| 4. Hapus Lagu                                         |")
	fmt.Println("| 5. Cari Lagu                                          |")
	fmt.Println("| 6. Urutkan Lagu                                       |")
	fmt.Println("| 7. Hitung Rata-rata Jumlah Pendengar                  |")
	fmt.Println("| 8. Keluar                                             |")
	fmt.Println("=========================================================")
}

// main adalah fungsi utama yang menjalankan loop program dan interaksi dengan pengguna.
func main() {
	var pilihanMenu int
	var errConv, errRead error
	var inputLine, inputTrimmed string

	initDummyData()
	mainReader := bufio.NewReader(os.Stdin)

	var programBerjalan bool = true
	for programBerjalan {
		fmt.Println("=========================================================")
		fmt.Println("|           DATABASE LAGU FAVORIT MAHASISWA             |")
		fmt.Println("=========================================================")
		tampilkanMenu()
		fmt.Print("Pilih menu (1-8): ")

		inputLine, errRead = mainReader.ReadString('\n')
		if errRead != nil {
			fmt.Println("| Kesalahan membaca pilihan menu:", errRead)
			errConv = fmt.Errorf("kesalahan baca")
		}
		inputTrimmed = strings.TrimSpace(inputLine)
		if errRead == nil {
			pilihanMenu, errConv = strconv.Atoi(inputTrimmed)
		}

		fmt.Println("=========================================================")
		if errConv != nil {
			fmt.Println("| ! Input tidak valid. Harap masukkan angka.            |")
		} else {
			switch pilihanMenu {
			case 1:
				tambahLagu()
			case 2:
				tampilkanDataLagu(daftarLagu, jumlahLaguTerisi)
			case 3:
				updateLagu()
			case 4:
				menuHapusLagu(mainReader)
			case 5:
				menuCariLagu(mainReader)
			case 6:
				menuUrutkanLagu(mainReader)
			case 7:
				rataRata := hitungRataRataPendengar()
				rataRataStr := fmt.Sprintf("%.2f", rataRata)
				fmt.Printf("| Rata-rata pendengar per lagu: %22s  |\n", rataRataStr)
			case 8:
				fmt.Println("| Terima kasih telah menggunakan program ini.           |")
				fmt.Println("=========================================================")
				programBerjalan = false
			default:
				fmt.Println("! Pilihan tidak valid. Silakan pilih (1-8).")
			}
		}

		if programBerjalan {
			fmt.Println("=========================================================")
			fmt.Print("\nTekan Enter untuk kembali ke menu...")
			var enterKeyInput string
			var errEnterKey error
			enterKeyInput, errEnterKey = mainReader.ReadString('\n')
			if errEnterKey != nil {
				fmt.Println("\nError kecil saat membaca 'Enter', program lanjut.")
			}
			if len(enterKeyInput) < 0 {
				return
			}
		}
	}
}