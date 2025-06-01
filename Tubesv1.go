package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Song mendefinisikan struktur untuk data lagu.
type Song struct {
	ID              string
	Judul           string
	Penyanyi        string
	Durasi          float64 // Durasi dalam menit
	JumlahPendengar int
}

// Konstanta untuk ukuran array statis.
const NMAXLagu = 100

// daftarLagu adalah array statis untuk menyimpan lagu.
var daftarLagu [NMAXLagu]Song
var jumlahLaguTerisi int // Melacak jumlah lagu yang sebenarnya terisi dalam array

// initDummyData menginisialisasi daftarLagu dengan beberapa data contoh.
func initDummyData() {
	dummyDataSlice := []Song{
		{ID: "L001", Judul: "Monokrom", Penyanyi: "Tulus", Durasi: 3.5, JumlahPendengar: 350},
		{ID: "L002", Judul: "Sial", Penyanyi: "Mahalini", Durasi: 4.1, JumlahPendengar: 420},
		{ID: "L003", Judul: "As It Was", Penyanyi: "Harry Styles", Durasi: 2.8, JumlahPendengar: 550},
		{ID: "L004", Judul: "Bertaut", Penyanyi: "Nadin Amizah", Durasi: 4.3, JumlahPendengar: 280},
		{ID: "L005", Judul: "Glimpse of Us", Penyanyi: "Joji", Durasi: 3.9, JumlahPendengar: 480},
		{ID: "L006", Judul: "Hati-Hati di Jalan", Penyanyi: "Tulus", Durasi: 4.0, JumlahPendengar: 320},
		{ID: "L007", Judul: "Until I Found You", Penyanyi: "Stephen Sanchez", Durasi: 3.0, JumlahPendengar: 510},
		{ID: "L008", Judul: "Komang", Penyanyi: "Raim Laode", Durasi: 3.7, JumlahPendengar: 390},
		{ID: "L009", Judul: "Anti-Hero", Penyanyi: "Taylor Swift", Durasi: 3.3, JumlahPendengar: 600},
		{ID: "L010", Judul: "Tak Ingin Usai", Penyanyi: "Keisya Levronka", Durasi: 4.6, JumlahPendengar: 250},
		{ID: "L011", Judul: "Dandelions", Penyanyi: "Ruth B.", Durasi: 3.8, JumlahPendengar: 450},
		{ID: "L012", Judul: "Runtuh", Penyanyi: "Feby Putri ft. Fiersa Besari", Durasi: 4.2, JumlahPendengar: 310},
		{ID: "L013", Judul: "Kill Bill", Penyanyi: "SZA", Durasi: 2.6, JumlahPendengar: 580},
		{ID: "L014", Judul: "To The Bone", Penyanyi: "Pamungkas", Durasi: 5.8, JumlahPendengar: 290},
		{ID: "L015", Judul: "Here With Me", Penyanyi: "d4vd", Durasi: 4.0, JumlahPendengar: 470},
	}
	jumlahLaguTerisi = 0
	for i := 0; i < len(dummyDataSlice) && jumlahLaguTerisi < NMAXLagu; i++ {
		daftarLagu[jumlahLaguTerisi] = dummyDataSlice[i]
		jumlahLaguTerisi++
	}
	insertionSortByIDAsc()
}

// generateNewID menghasilkan ID unik baru untuk lagu.
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

// findLaguByID mencari lagu berdasarkan ID (sequential search).
func findLaguByID(id string) (int, bool) {
	for i := 0; i < jumlahLaguTerisi; i++ {
		if daftarLagu[i].ID == id {
			return i, true
		}
	}
	return -1, false
}

// tambahLagu menambahkan lagu baru ke array.
func tambahLagu() {
	var laguBaru Song
	var input string
	var errRead error
	var parsedDurasi float64
	var errScanFloat error

	if jumlahLaguTerisi >= NMAXLagu {
		fmt.Println("Kapasitas database penuh. Tidak bisa menambah lagu.  ")
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
				fmt.Println("Judul Lagu tidak boleh kosong. Silakan coba lagi.    ")
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
				fmt.Println("Nama Penyanyi tidak boleh kosong. Silakan coba lagi. ")
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
				fmt.Println("Durasi tidak valid. Harap masukkan angka positif.  ")
				if errScanFloat != nil {
					// fmt.Println("| Detail kesalahan konversi durasi:", errScanFloat)
				}
			}
		} else {
			fmt.Println("Kesalahan membaca input durasi:", errRead)
		}
	}

	daftarLagu[jumlahLaguTerisi] = laguBaru
	jumlahLaguTerisi++
	insertionSortByIDAsc()
	fmt.Println("Lagu berhasil ditambahkan!                           ")
}

// updateLagu memperbarui data lagu yang ada.
func updateLagu() {
	var idCari, input, newJudul, newPenyanyi string
	var index int
	var found bool
	var errRead, errScanFloat, errScanInt error
	var parsedDurasi float64
	var newJumlahPendengar int

	fmt.Println("| --- Update Data Lagu (ID Tidak Dapat Diubah) ---   |")
	if jumlahLaguTerisi == 0 {
		fmt.Println("| Database lagu kosong. Tidak ada data untuk diupdate. |")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("| Masukkan ID Lagu yang akan diupdate: ")
	input, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca ID:", errRead); return }
	idCari = strings.TrimSpace(input)

	index, found = binarySearchByID(idCari)
	if !found {
		fmt.Printf("| Lagu dengan ID '%s' tidak ditemukan.             |\n", idCari)
		return
	}
	
	fmt.Println("|------------------------------------------------------|")
	fmt.Println("| Data Lagu Saat Ini (ID Tidak Dapat Diubah):          |")
	fmt.Printf("| ID                : %-33s |\n", daftarLagu[index].ID)
	fmt.Printf("| Judul Lama        : %-33s |\n", truncateString(daftarLagu[index].Judul, 33))
	fmt.Printf("| Penyanyi Lama     : %-33s |\n", truncateString(daftarLagu[index].Penyanyi, 33))
	fmt.Printf("| Durasi Lama       : %-33.1f |\n", daftarLagu[index].Durasi)
	fmt.Printf("| Jml Pendengar Lama: %-30d |\n", daftarLagu[index].JumlahPendengar)
	fmt.Println("|------------------------------------------------------|")
	fmt.Println("| Masukkan data baru (kosongkan jika tidak ingin diubah): |")

	fmt.Printf("| Judul Baru (Lama: %s): ", truncateString(daftarLagu[index].Judul, 20)) // Tampilkan judul lama dipotong
	input, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca judul baru:", errRead); return }
	newJudul = strings.TrimSpace(input)
	if newJudul != "" {
		daftarLagu[index].Judul = newJudul
	}

	fmt.Printf("| Penyanyi Baru (Lama: %s): ", truncateString(daftarLagu[index].Penyanyi, 20)) // Tampilkan penyanyi lama dipotong
	input, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca penyanyi baru:", errRead); return }
	newPenyanyi = strings.TrimSpace(input)
	if newPenyanyi != "" {
		daftarLagu[index].Penyanyi = newPenyanyi
	}

	var durasiUpdateSelesai bool = false
	for !durasiUpdateSelesai {
		fmt.Printf("| Durasi Baru (Lama: %.1f): ", daftarLagu[index].Durasi)
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
					fmt.Println("| Durasi tidak valid. Harap masukkan angka positif.  |")
				}
			}
		} else {
			fmt.Println("| Kesalahan membaca durasi baru:", errRead)
		}
	}

	var pendengarUpdateSelesai bool = false
	for !pendengarUpdateSelesai {
		fmt.Printf("| Jumlah Pendengar Baru (Lama: %d): ", daftarLagu[index].JumlahPendengar)
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
					fmt.Println("| Jumlah pendengar tidak valid. Harap masukkan angka non-negatif. |")
				}
			}
		} else {
			fmt.Println("| Kesalahan membaca jumlah pendengar baru:", errRead)
		}
	}

	fmt.Println("|------------------------------------------------------|")
	fmt.Println("| Data lagu berhasil diupdate!                         |")
}

// hapusLagu menghapus lagu dari array.
func hapusLagu() {
	var idCari, konfirmasi string
	var index int
	var found bool
	var errRead error

	fmt.Println("| --- Hapus Lagu ---                                 |")
	if jumlahLaguTerisi == 0 {
		fmt.Println("| Database lagu kosong. Tidak ada data untuk dihapus.  |")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("| Masukkan ID Lagu yang akan dihapus: ")
	idCari, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca ID:", errRead); return }
	idCari = strings.TrimSpace(idCari)

	index, found = findLaguByID(idCari)
	if !found {
		fmt.Printf("| Lagu dengan ID '%s' tidak ditemukan.               |\n", idCari)
		return
	}

	fmt.Println("|------------------------------------------------------|")
	fmt.Println("| Lagu yang akan dihapus:                             |")
	fmt.Printf("| ID         : %-37s |\n", daftarLagu[index].ID)
	fmt.Printf("| Judul      : %-37s |\n", truncateString(daftarLagu[index].Judul, 37))
	fmt.Printf("| Penyanyi   : %-37s |\n", truncateString(daftarLagu[index].Penyanyi,37))
	fmt.Printf("| Durasi     : %-37.1f |\n", daftarLagu[index].Durasi)
	fmt.Printf("| Pendengar  : %-37d |\n", daftarLagu[index].JumlahPendengar)
	fmt.Println("|------------------------------------------------------|")

	fmt.Print("| Apakah Anda yakin ingin menghapus lagu ini? (Y/N): ")
	konfirmasi, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca konfirmasi:", errRead); return }
	konfirmasi = strings.TrimSpace(strings.ToUpper(konfirmasi))

	if konfirmasi == "Y" {
		for i := index; i < jumlahLaguTerisi-1; i++ {
			daftarLagu[i] = daftarLagu[i+1]
		}
		jumlahLaguTerisi--
		daftarLagu[jumlahLaguTerisi] = Song{}
		insertionSortByIDAsc()
		fmt.Println("| Lagu berhasil dihapus!                             |")
	} else {
		fmt.Println("| Penghapusan dibatalkan.                            |")
	}
}

// binarySearchByID mencari lagu berdasarkan ID.
func binarySearchByID(id string) (int, bool) {
	low, high := 0, jumlahLaguTerisi-1
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

// sequentialSearchByPenyanyi mencari lagu berdasarkan nama penyanyi.
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

// --- Fungsi-Fungsi Pengurutan (beroperasi pada array global daftarLagu) ---
func insertionSortByIDAsc() { /* Implementasi sama */
	var i, j int; var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]; j = i - 1
		for j >= 0 && daftarLagu[j].ID > key.ID { daftarLagu[j+1] = daftarLagu[j]; j-- }
		daftarLagu[j+1] = key
	}
}
func selectionSortByIDDesc() { /* Implementasi sama */
	var i, j, idxExtreme int; n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ { if daftarLagu[j].ID > daftarLagu[idxExtreme].ID { idxExtreme = j } }
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}
func insertionSortByJudulAsc() { /* Implementasi sama */
	var i, j int; var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]; j = i - 1
		for j >= 0 && strings.ToLower(daftarLagu[j].Judul) > strings.ToLower(key.Judul) { daftarLagu[j+1] = daftarLagu[j]; j-- }
		daftarLagu[j+1] = key
	}
}
func selectionSortByJudulDesc() { /* Implementasi sama */
	var i, j, idxExtreme int; n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ { if strings.ToLower(daftarLagu[j].Judul) > strings.ToLower(daftarLagu[idxExtreme].Judul) { idxExtreme = j } }
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}
func insertionSortByPenyanyiAsc() { /* Implementasi sama */
	var i, j int; var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]; j = i - 1
		for j >= 0 && strings.ToLower(daftarLagu[j].Penyanyi) > strings.ToLower(key.Penyanyi) { daftarLagu[j+1] = daftarLagu[j]; j-- }
		daftarLagu[j+1] = key
	}
}
func selectionSortByPenyanyiDesc() { /* Implementasi sama */
	var i, j, idxExtreme int; n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ { if strings.ToLower(daftarLagu[j].Penyanyi) > strings.ToLower(daftarLagu[idxExtreme].Penyanyi) { idxExtreme = j } }
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}
func insertionSortByDurasiAsc() { /* Implementasi sama */
	var i, j int; var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]; j = i - 1
		for j >= 0 && daftarLagu[j].Durasi > key.Durasi { daftarLagu[j+1] = daftarLagu[j]; j-- }
		daftarLagu[j+1] = key
	}
}
func selectionSortByDurasiDesc() { /* Implementasi sama */
	var i, j, idxExtreme int; n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ { if daftarLagu[j].Durasi > daftarLagu[idxExtreme].Durasi { idxExtreme = j } }
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}
func insertionSortByPendengarAsc() { /* Implementasi sama */
	var i, j int; var key Song
	for i = 1; i < jumlahLaguTerisi; i++ {
		key = daftarLagu[i]; j = i - 1
		for j >= 0 && daftarLagu[j].JumlahPendengar > key.JumlahPendengar { daftarLagu[j+1] = daftarLagu[j]; j-- }
		daftarLagu[j+1] = key
	}
}
func selectionSortByPendengarDesc() { /* Implementasi sama */
	var i, j, idxExtreme int; n := jumlahLaguTerisi
	for i = 0; i < n-1; i++ {
		idxExtreme = i
		for j = i + 1; j < n; j++ { if daftarLagu[j].JumlahPendengar > daftarLagu[idxExtreme].JumlahPendengar { idxExtreme = j } }
		daftarLagu[i], daftarLagu[idxExtreme] = daftarLagu[idxExtreme], daftarLagu[i]
	}
}

// hitungRataRataPendengar menghitung rata-rata jumlah pendengar.
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

// main adalah fungsi utama program.
func main() {
	var pilihanMenu int
	var errConv, errRead error
	var inputLine, inputTrimmed string

	initDummyData()
	mainReader := bufio.NewReader(os.Stdin)

	var programBerjalan bool = true
	for programBerjalan {
		fmt.Println("======================================================")
		fmt.Println("|        DATABASE LAGU FAVORIT MAHASISWA             |")
		fmt.Println("======================================================")
		tampilkanMenu()
		fmt.Print("Pilih menu (1-8): ")

		inputLine, errRead = mainReader.ReadString('\n')
		if errRead != nil {
			fmt.Println("| Kesalahan membaca pilihan menu:", errRead)
			errConv = fmt.Errorf("kesalahan baca") // Set errConv agar masuk ke blok tekan enter
		}
		inputTrimmed = strings.TrimSpace(inputLine)
		if errRead == nil { // Hanya coba konversi jika tidak ada error baca
			pilihanMenu, errConv = strconv.Atoi(inputTrimmed)
		}


		fmt.Println("======================================================")
		if errConv != nil {
			fmt.Println("| ! Input tidak valid. Harap masukkan angka.         |")
		} else {
			switch pilihanMenu {
			case 1:
				tambahLagu()
			case 2:
				tampilkanDataLagu(daftarLagu, jumlahLaguTerisi)
			case 3:
				updateLagu()
			case 4:
				hapusLagu()
			case 5:
				menuCariLagu(mainReader)
			case 6:
				menuUrutkanLagu(mainReader)
			case 7:
				fmt.Println("| --- Rata-rata Jumlah Pendengar ---                 |")
				rataRata := hitungRataRataPendengar()
				rataRataStr := fmt.Sprintf("%.2f", rataRata)
				// Lebar total untuk baris ini sekitar 50 karakter. Konten "Rata-rata ... :" ~28 char.
				// Sisa untuk angka dan padding. Angka bisa bervariasi.
				// fmt.Printf("| Rata-rata pendengar per lagu: %-22.2f |\n", rataRata) // Sebelumnya
				// Membuat perataan kanan untuk angka dengan padding dinamis
				fmt.Printf("| Rata-rata pendengar per lagu: %22s |\n", rataRataStr)


			case 8:
				fmt.Println("| Terima kasih telah menggunakan program ini.        |")
				fmt.Println("======================================================")
				programBerjalan = false
			default:
				fmt.Println("| ! Pilihan tidak valid. Silakan pilih (1-8).        |")
			}
		}

		if programBerjalan { // Hanya minta Enter jika program belum akan keluar
			fmt.Println("======================================================")
			fmt.Print("\nTekan Enter untuk kembali ke menu...")
			_, errRead = mainReader.ReadString('\n')
			if errRead != nil {
				// fmt.Println("| Kesalahan saat menunggu Enter:", errRead) // Opsional
			}
		}
	}
}

// tampilkanMenu menampilkan opsi menu utama.
func tampilkanMenu() {
	fmt.Println("| 1. Tambah Lagu Baru                                |")
	fmt.Println("| 2. Tampilkan Semua Lagu                            |")
	fmt.Println("| 3. Update Data Lagu                                |")
	fmt.Println("| 4. Hapus Lagu                                      |")
	fmt.Println("| 5. Cari Lagu                                       |")
	fmt.Println("| 6. Urutkan Lagu                                    |")
	fmt.Println("| 7. Hitung Rata-rata Jumlah Pendengar               |")
	fmt.Println("| 8. Keluar                                          |")
	fmt.Println("======================================================")
}

// menuCariLagu menampilkan sub-menu untuk pencarian.
func menuCariLagu(reader *bufio.Reader) {
	var sub int
	var errConv, errRead error
	var inputLine, inputTrimmed, id, nama string
	var index int
	var found bool
	var hasilPencarian [NMAXLagu]Song
	var jumlahHasil int

	fmt.Println("| --- Cari Lagu ---                                  |")
	fmt.Println("| 1. Cari berdasarkan ID (Binary Search)             |")
	fmt.Println("| 2. Cari berdasarkan Nama Penyanyi (Sequential Search)|")
	fmt.Println("======================================================")
	fmt.Print("| Pilih sub-menu (1-2): ")

	inputLine, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca sub-menu:", errRead); return }
	inputTrimmed = strings.TrimSpace(inputLine)
	sub, errConv = strconv.Atoi(inputTrimmed)

	if errConv != nil {
		fmt.Println("| Pilihan tidak valid. Kembali ke menu utama.        |")
		return
	}

	switch sub {
	case 1:
		fmt.Print("| Masukkan ID lagu: ")
		id, errRead = reader.ReadString('\n')
		if errRead != nil { fmt.Println("| Kesalahan membaca ID:", errRead); return }
		id = strings.TrimSpace(id)
		index, found = binarySearchByID(id)
		if found {
			var tempArrayForDisplay [NMAXLagu]Song // Array sementara untuk satu lagu
			tempArrayForDisplay[0] = daftarLagu[index]
			tampilkanDataLagu(tempArrayForDisplay, 1)
		} else {
			fmt.Println("| Lagu tidak ditemukan.                              |")
		}
	case 2:
		fmt.Print("| Masukkan nama penyanyi (boleh sebagian): ")
		nama, errRead = reader.ReadString('\n')
		if errRead != nil { fmt.Println("| Kesalahan membaca nama penyanyi:", errRead); return }
		nama = strings.TrimSpace(nama)
		jumlahHasil = sequentialSearchByPenyanyi(nama, &hasilPencarian)
		if jumlahHasil == 0 {
			fmt.Println("| Tidak ditemukan lagu dari penyanyi tersebut.         |")
		} else {
			tampilkanDataLagu(hasilPencarian, jumlahHasil)
		}
	default:
		fmt.Println("| Pilihan sub-menu tidak valid.                      |")
	}
}

// menuUrutkanLagu menampilkan sub-menu untuk pengurutan.
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
	fmt.Print("| Pilih sub-menu (1–11): ")

	inputLine, errRead = reader.ReadString('\n')
	if errRead != nil { fmt.Println("| Kesalahan membaca sub-menu:", errRead); return }
	inputTrimmed = strings.TrimSpace(inputLine)
	sub, errConv = strconv.Atoi(inputTrimmed)

	if errConv != nil {
		fmt.Println("| Pilihan tidak valid. Kembali ke menu utama.        |")
		return
	}
	if sub == 11 { return }

	sorted := true
	switch sub {
	case 1: insertionSortByIDAsc()
	case 2: selectionSortByIDDesc()
	case 3: insertionSortByJudulAsc()
	case 4: selectionSortByJudulDesc()
	case 5: insertionSortByPenyanyiAsc()
	case 6: selectionSortByPenyanyiDesc()
	case 7: selectionSortByDurasiDesc()
	case 8: insertionSortByDurasiAsc()
	case 9: selectionSortByPendengarDesc()
	case 10: insertionSortByPendengarAsc()
	default:
		fmt.Println("| Pilihan sub-menu tidak valid.                      |")
		sorted = false
	}

	if sorted {
		tampilkanDataLagu(daftarLagu, jumlahLaguTerisi)
	}
}

// truncateString memotong string jika lebih panjang dari maxLength dan menambahkan "..."
func truncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		if maxLength-3 > 0 {
			return s[:maxLength-3] + "..."
		}
		return s[:maxLength] // Jika maxLength terlalu kecil untuk "..."
	}
	return s
}


// tampilkanDataLagu adalah fungsi generik untuk menampilkan data lagu dari array.
// Menerima array (disalin), jumlah data valid, dan judul kustom untuk tabel.
// Perhatikan: Menerima array [NMAXLagu]Song akan menyalin seluruh array.
// Jika NMAXLagu besar, ini bisa tidak efisien. Alternatifnya adalah pointer.
func tampilkanDataLagu(data [NMAXLagu]Song, count int) {
	// Definisi lebar kolom yang konsisten
	idWidth := 4
	judulWidth := 30
	penyanyiWidth := 38 // Cukup untuk nama penyanyi yang sangat panjang
	durasiWidth := 6    // Format "XX.X"
	pendengarWidth := 9 // Untuk angka besar

	// Membuat format string untuk garis horizontal
	// Panjang setiap segmen garis = lebarKontenKolom + 2 (untuk spasi padding)
	headerLine := fmt.Sprintf("+-%s-+-%s-+-%s-+-%s-+-%s-+",
		strings.Repeat("-", idWidth),
		strings.Repeat("-", judulWidth),
		strings.Repeat("-", penyanyiWidth),
		strings.Repeat("-", durasiWidth),
		strings.Repeat("-", pendengarWidth))
	fmt.Println(headerLine)

	// Header kolom
	fmt.Printf("| %-*s | %-*s | %-*s | %*s | %*s |\n",
		idWidth, "ID",
		judulWidth, "JUDUL",
		penyanyiWidth, "PENYANYI",
		durasiWidth, "DURASI",    // Rata kanan
		pendengarWidth, "PENDENGAR") // Rata kanan
	fmt.Println(headerLine)

	if count == 0 {
		message := "Belum ada lagu untuk ditampilkan."
		// Hitung total lebar konten dalam tabel (tanpa pipa terluar) untuk centering
		// (width1+2) + (width2+2) + ... + (widthN+2) + (N-1)*1 (untuk pipa internal)
		// Atau lebih mudah: len(headerLine) - 2
		effectiveTableWidth := len(headerLine) - 2
		paddingLen := (effectiveTableWidth - len(message)) / 2
		if paddingLen < 0 { paddingLen = 0 }
		rightPaddingLen := effectiveTableWidth - len(message) - paddingLen
		if rightPaddingLen < 0 { rightPaddingLen = 0 }

		fmt.Printf("|%*s%s%*s|\n", paddingLen, "", message, rightPaddingLen, "")
		fmt.Println(headerLine)
		return
	}

	// Data lagu
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