// Tubes Alpro
// Kelompok Topik 20 (Budget Traveling)
// Nama anggota kelompok :
// Tiara Fitriannisha (103012400015)
// Muh Khuba Adla - 1030400327
// Nurendra Rifat Wicaksono - 103012400279

package main

import (
	"fmt"
)

type pengguna struct {
	nama string
	id   int
}
type Pengeluaran struct {
	akomodasi    int
	transportasi int
	makanan      int
	hiburan      int
}

const PAX int = 10

const (
	Reset = "\033[0m"
	sp    = "\033[38;5;218m" // Font warna soft pink
	bold  = "\033[1m"
)

type tabPengeluaran [PAX]Pengeluaran
type tabPengguna [PAX]pengguna

func main() {
	var pax, opsi, budget int
	var arrayUser tabPengguna
	var arrayPengeluaran tabPengeluaran
	var pilihan int

	fmt.Printf(sp + "Masukkan Jumlah Pengguna : " + Reset)
	fmt.Scan(&pax)
	fmt.Printf(sp + "Masukkan Budget Perjalanan : " + Reset)
	fmt.Scan(&budget)
	inputData(&arrayUser, &pax)
	for {
		tampilanMenu()
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			tambahPengeluaran(arrayUser, &arrayPengeluaran, pax)
		case 2:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			editPengeluaran(arrayUser, arrayPengeluaran, pax)
		case 3:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			hapusPengeluaran(arrayUser, &arrayPengeluaran, pax)
		case 4:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			lihatSemua(arrayUser, arrayPengeluaran, pax)
		case 5:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			fmt.Println(sp + "╔════════════════════════════╗" + Reset)
			fmt.Println(sp + "║     PILIH METODE CARI      ║" + Reset)
			fmt.Println(sp + "╠════════════════════════════╣" + Reset)
			fmt.Println(sp + "║ 1 ║ Sequential Search      ║" + Reset)
			fmt.Println(sp + "║ 2 ║ Binary Search          ║" + Reset)
			fmt.Println(sp + "╚════════════════════════════╝" + Reset)
			fmt.Print(sp + bold + "Masukkan pilihan (1/2): " + Reset + Reset)
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				cariPengeluaranSeq(arrayUser, arrayPengeluaran, pax)
			case 2:
				cariPengeluaranBinary(arrayUser, arrayPengeluaran, pax)
			}
		case 6:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			urutkanPengeluaran(&arrayUser, &arrayPengeluaran, pax)
		case 7:
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			laporanAkhir(arrayUser, arrayPengeluaran, pax, budget)
		}
		if opsi >= 8 || opsi < 1 {
			fmt.Println("⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆˚✿˖° 𐙚 ₊ ⊹ ♡⋆")
			fmt.Println(sp + bold + "terimaksih ya!" + Reset + Reset + Reset)
			break
		}

	}

}
func inputData(A *tabPengguna, n *int) { // sudah rapih
	for i := 0; i < *n; i++ {
		fmt.Print(sp + "Masukkan Nama Pengguna: " + Reset)
		fmt.Scan(&A[i].nama)

		var id int
		for {
			fmt.Print(sp + "Masukkan ID Pengguna : " + Reset)
			fmt.Scan(&id)

			duplikat := false
			for j := 0; j < i; j++ {
				if A[j].id == id {
					duplikat = true
					fmt.Println(sp + " ID sudah dipakai! Coba lagi." + Reset)
					break
				}
			}

			if !duplikat {
				A[i].id = id
				break
			}
		}
	}
}

func tampilanMenu() { // sudah rapih

	fmt.Println(sp + "█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█" + Reset)
	fmt.Println(sp + "█                      G O   B U D G E T                      █" + Reset)
	fmt.Println(sp + "█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█" + Reset)
	fmt.Println(sp + "             SELAMAT DATANG DI APLIKASI GO BUDGET             " + Reset)
	fmt.Println(sp + "               Aplikasi Pencatatan Keuangan Anda              " + Reset)
	fmt.Println(sp + "▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀" + Reset)
	fmt.Println(sp + "█                                                             █" + Reset)
	fmt.Println(sp + "█    ➤  1. Tambah Pengeluaran                                 █" + Reset)
	fmt.Println(sp + "█    ➤  2. Ubah Pengeluaran                                   █" + Reset)
	fmt.Println(sp + "█    ➤  3. Hapus Pengeluaran                                  █" + Reset)
	fmt.Println(sp + "█    ➤  4. Lihat Semua Pengeluaran                            █" + Reset)
	fmt.Println(sp + "█    ➤  5. Cari Pengeluaran                                   █" + Reset)
	fmt.Println(sp + "█    ➤  6. Urutkan Pengeluaran                                █" + Reset)
	fmt.Println(sp + "█    ➤  7. Laporan Akhir                                      █" + Reset)
	fmt.Println(sp + "█    ➤  8. Exit                                               █" + Reset)
	fmt.Println(sp + "█                                                             █" + Reset)
	fmt.Println(sp + "█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█" + Reset)
	fmt.Println(sp + bold + " Masukan angka (1-8) lalu ENTER untuk memulai " + Reset + Reset)
}
func tambahPengeluaran(A tabPengguna, C *tabPengeluaran, n int) { // sudah rapih
	var i, IDUser, jumlah int
	var pilihan int
	var kategori string
	var hotelPilihan int
	var harga int
	var malam int
	var total int
	var hargaTiket int
	var negara, negaraBaru int

	fmt.Println(" Tambah Pengeluaran ")
	fmt.Printf("Masukkan ID : ")
	fmt.Scan(&IDUser)

	for i = 0; i < n; i++ {
		if A[i].id == IDUser {
			fmt.Println(sp + "╔════╦═══════════════════╗" + Reset)
			fmt.Println(sp + "║ No ║     Kategori      ║" + Reset)
			fmt.Println(sp + "╠════╬═══════════════════╣" + Reset)
			fmt.Println(sp + "║ 1  ║ Akomodasi         ║" + Reset)
			fmt.Println(sp + "║ 2  ║ Transportasi      ║" + Reset)
			fmt.Println(sp + "║ 3  ║ Makanan           ║" + Reset)
			fmt.Println(sp + "║ 4  ║ Hiburan           ║" + Reset)
			fmt.Println(sp + "╚════╩═══════════════════╝" + Reset)
			fmt.Printf(sp + bold + "Masukkan kategori (1-4): " + Reset + Reset)
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println(sp + "╔══════════════════════════════════════╗" + Reset)
				fmt.Println(sp + "║              Pilih Hotel             ║" + Reset)
				fmt.Println(sp + "╠══════════════════════════════════════╣" + Reset)
				fmt.Println(sp + "║ 1. Hotel Bintang 3  (Rp300.000/mlm)  ║" + Reset)
				fmt.Println(sp + "║ 2. Hotel Bintang 4  (Rp500.000/mlm)  ║" + Reset)
				fmt.Println(sp + "║ 3. Hotel Bintang 5  (Rp800.000/mlm)  ║" + Reset)
				fmt.Println(sp + "╚══════════════════════════════════════╝" + Reset)
				fmt.Printf(sp + bold + "Masukan Pilihan : " + Reset + Reset)
				fmt.Scan(&hotelPilihan)

				switch hotelPilihan {
				case 1:
					harga = 300000
					kategori = "Hotel Bintang 3"
				case 2:
					harga = 500000
					kategori = "Hotel Bintang 4"
				case 3:
					harga = 800000
					kategori = "Hotel Bintang 5"
				default:
					fmt.Println(sp + bold + "Pilihan hotel tidak valid." + Reset + Reset)
					return
				}

				fmt.Printf(sp + "Masukkan jumlah malam menginap: " + Reset)
				fmt.Scan(&malam)

				total = harga * malam
				C[i].akomodasi = C[i].akomodasi + total
				fmt.Print(sp+"Berhasil menambahkan Akomodasi ", kategori, " Selama ", malam, " malam"+Reset)

			case 2:
				fmt.Println(sp + "╔════╦════════════════════╦══════════════════════╗" + Reset)
				fmt.Println(sp + "║ No ║   Negara Tujuan    ║     Harga Tiket      ║" + Reset)
				fmt.Println(sp + "╠════╬════════════════════╬══════════════════════╣" + Reset)
				fmt.Println(sp + "║ 1  ║ Belanda            ║ Rp10.800.000         ║" + Reset)
				fmt.Println(sp + "║ 2  ║ Italia             ║ Rp11.500.000         ║" + Reset)
				fmt.Println(sp + "║ 3  ║ Jepang             ║ Rp7.000.000          ║" + Reset)
				fmt.Println(sp + "║ 4  ║ Jerman             ║ Rp11.000.000         ║" + Reset)
				fmt.Println(sp + "║ 5  ║ Korea              ║ Rp6.500.000          ║" + Reset)
				fmt.Println(sp + "║ 6  ║ Prancis            ║ Rp12.000.000         ║" + Reset)
				fmt.Println(sp + "║ 7  ║ Spanyol            ║ Rp10.500.000         ║" + Reset)
				fmt.Println(sp + "║ 8  ║ Thailand           ║ Rp3.500.000          ║" + Reset)
				fmt.Println(sp + "╚════╩════════════════════╩══════════════════════╝" + Reset)
				fmt.Printf(sp + "Masukkan nomor negara tujuan: " + Reset)
				fmt.Scan(&negara)
				switch negara {
				case 1:
					hargaTiket = 10800000
					kategori = "Ke Belanda"
				case 2:
					hargaTiket = 11500000
					kategori = "Ke Itali"
				case 3:
					hargaTiket = 7000000
					kategori = "Ke Jepang"
				case 4:
					hargaTiket = 11000000
					kategori = "Ke Jerman"
				case 5:
					hargaTiket = 6500000
					kategori = "Ke Korea"
				case 6:
					hargaTiket = 12000000
					kategori = "Ke Prancis"
				case 7:
					hargaTiket = 10500000
					kategori = "Ke Spanyol"
				case 8:
					hargaTiket = 3500000
					kategori = "Ke Thailand"
				default:
					fmt.Println(sp + "Negara belum terdaftar, masukkan biaya manual." + Reset)
					fmt.Printf(sp + "Masukkan nama negara: " + Reset)
					fmt.Scan(&negaraBaru)
					fmt.Print(sp + "Masukkan biaya tiket :" + Reset)
					fmt.Scan(&hargaTiket)
				}

				C[i].transportasi = C[i].transportasi + hargaTiket
				fmt.Print(sp+"Berhasil menambahkan Transportasi ", kategori, " Seharga Rp ", hargaTiket)

			case 3:
				fmt.Printf(sp + "Masukan jumlah : ")
				fmt.Scan(&jumlah)
				C[i].makanan += jumlah
				kategori = "makanan"
				fmt.Print(sp+"Berhasil menambahkan  ", kategori, " Seharga Rp ", jumlah)
			case 4:
				fmt.Printf(sp + "Masukan jumlah : ")
				fmt.Scan(&jumlah)
				C[i].hiburan += jumlah
				kategori = "hiburan"
				fmt.Print(sp+"Berhasil menambahkan  ", kategori, " Seharga Rp ", jumlah)
			default:
				fmt.Println(sp + "Kategori tidak valid!")
				return
			}
			return
		}
		fmt.Println(sp + "ID tidak ditemukan.")
	}
}
func editPengeluaran(A tabPengguna, B tabPengeluaran, n int) { // sudah aman
	var pilihan, jumlah, IDUser, i int
	var harga, hargaTiket int
	var hotelPilihan int
	var kategori string
	var negaraBaru string
	var malam int
	var found bool
	var negara int

	fmt.Printf(sp + "Masukkan ID pengguna: ")
	fmt.Scan(&IDUser)

	found = false

	for i = 0; i < n; i++ {
		if IDUser == A[i].id {
			found = true
			fmt.Println(sp + "╔════╦═══════════════════╗")
			fmt.Println(sp + "║ No ║     Kategori      ║")
			fmt.Println(sp + "╠════╬═══════════════════╣")
			fmt.Println(sp + "║ 1  ║ Akomodasi         ║")
			fmt.Println(sp + "║ 2  ║ Transportasi      ║")
			fmt.Println(sp + "║ 3  ║ Makanan           ║")
			fmt.Println(sp + "║ 4  ║ Hiburan           ║")
			fmt.Println(sp + "╚════╩═══════════════════╝")
			fmt.Printf(sp + "Masukkan kategori (1-4): ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println(sp + "╔══════════════════════════════════════╗")
				fmt.Println(sp + "║              Pilih Hotel             ║")
				fmt.Println(sp + "╠══════════════════════════════════════╣")
				fmt.Println(sp + "║ 1. Hotel Bintang 3  (Rp300.000/mlm)  ║")
				fmt.Println(sp + "║ 2. Hotel Bintang 4  (Rp500.000/mlm)  ║")
				fmt.Println(sp + "║ 3. Hotel Bintang 5  (Rp800.000/mlm)  ║")
				fmt.Println(sp + "╚══════════════════════════════════════╝")
				fmt.Printf(sp + "Masukan Pilihan : ")
				fmt.Scan(&hotelPilihan)

				fmt.Printf(sp + "Masukan jumlah malam : ")
				fmt.Scan(&malam)
				B[i].akomodasi = 0

				switch hotelPilihan {
				case 1:
					harga = 300000
					kategori = "Hotel Bintang 3"
				case 2:
					harga = 500000
					kategori = "Hotel Bintang 4"
				case 3:
					harga = 800000
					kategori = "Hotel Bintang 5"
				default:
					fmt.Println(sp + "Pilihan hotel tidak valid.")
					return
				}
				B[i].akomodasi = B[i].akomodasi + harga
				fmt.Printf(sp+"Berhasil menambahkan akomodasi %s (%d malam) seharga Rp%d", kategori, malam, harga)

			case 2:
				fmt.Println(sp + "╔════╦════════════════════╦══════════════════════╗")
				fmt.Println(sp + "║ No ║   Negara Tujuan    ║     Harga Tiket      ║")
				fmt.Println(sp + "╠════╬════════════════════╬══════════════════════╣")
				fmt.Println(sp + "║ 1  ║ Belanda            ║ Rp10.800.000         ║")
				fmt.Println(sp + "║ 2  ║ Italia             ║ Rp11.500.000         ║")
				fmt.Println(sp + "║ 3  ║ Jepang             ║ Rp7.000.000          ║")
				fmt.Println(sp + "║ 4  ║ Jerman             ║ Rp11.000.000         ║")
				fmt.Println(sp + "║ 5  ║ Korea              ║ Rp6.500.000          ║")
				fmt.Println(sp + "║ 6  ║ Prancis            ║ Rp12.000.000         ║")
				fmt.Println(sp + "║ 7  ║ Spanyol            ║ Rp10.500.000         ║")
				fmt.Println(sp + "║ 8  ║ Thailand           ║ Rp3.500.000          ║")
				fmt.Println(sp + "╚════╩════════════════════╩══════════════════════╝")
				fmt.Printf(sp + "Masukkan nomor negara tujuan (1-8): ")
				fmt.Scan(&negara)
				B[i].transportasi = 0

				switch negara {
				case 1:
					hargaTiket = 10800000
					kategori = "Ke Belanda"
				case 2:
					hargaTiket = 11500000
					kategori = "Ke Italia"
				case 3:
					hargaTiket = 7000000
					kategori = "Ke Jepang"
				case 4:
					hargaTiket = 11000000
					kategori = "Ke Jerman"
				case 5:
					hargaTiket = 6500000
					kategori = "Ke Korea"
				case 6:
					hargaTiket = 12000000
					kategori = "Ke Prancis"
				case 7:
					hargaTiket = 10500000
					kategori = "Ke Spanyol"
				case 8:
					hargaTiket = 3500000
					kategori = "Ke Thailand"
				default:
					fmt.Println(sp + "Negara belum terdaftar, masukkan biaya manual!")
					fmt.Printf(sp + "Masukkan nama negara: ")
					fmt.Scan(&negaraBaru)
					fmt.Print(sp + "Masukkan biaya tiket: ")
					fmt.Scan(&hargaTiket)

					kategori = "Ke " + negaraBaru
				}
				B[i].transportasi += hargaTiket
				fmt.Printf(sp+"Berhasil menambahkan transportasi %s seharga Rp %d\n", kategori, hargaTiket)

			case 3:
				fmt.Print(sp + "Masukkan jumlah pengeluaran makanan: ")
				fmt.Scan(&jumlah)
				B[i].makanan = 0
				B[i].makanan += jumlah
				fmt.Printf(sp+"Berhasil menambahkan pengeluaran makanan sebesar Rp %d\n", jumlah)

			case 4:
				fmt.Print(sp + "Masukkan jumlah pengeluaran hiburan: ")
				fmt.Scan(&jumlah)
				B[i].hiburan = 0
				B[i].hiburan += jumlah
				fmt.Printf(sp+"Berhasil menambahkan pengeluaran hiburan sebesar Rp %d\n", jumlah)

			default:
				fmt.Println(sp + "Kategori tidak valid.")
			}
			return
		}
		if !found {
			fmt.Println(sp + "Id Pengguna Tidak Ditemukan ")
		}
	}
}
func hapusPengeluaran(A tabPengguna, C *tabPengeluaran, n int) { // sudah rapih
	var i, IDUser int
	var pilihan int
	var kategori string

	fmt.Println()
	fmt.Println(sp + "Hapus Pengeluaran ")
	fmt.Printf(sp + "Masukkan ID : ")
	fmt.Scan(&IDUser)

	if A[i].id == IDUser { // sudah ada batasan jika salah input
		fmt.Print(sp + "Masukan Pilihan : ")
		fmt.Scan(&pilihan)
		for i = 0; i < n; i++ {
			if A[i].id == IDUser {
				fmt.Println(sp + "╔════╦═══════════════════╗")
				fmt.Println(sp + "║ No ║     Kategori      ║")
				fmt.Println(sp + "╠════╬═══════════════════╣")
				fmt.Println(sp + "║ 1  ║ Akomodasi         ║")
				fmt.Println(sp + "║ 2  ║ Transportasi      ║")
				fmt.Println(sp + "║ 3  ║ Makanan           ║")
				fmt.Println(sp + "║ 4  ║ Hiburan           ║")
				fmt.Println(sp + "╚════╩═══════════════════╝")
				fmt.Printf(sp + "Masukkan kategori (1-4): ")
				fmt.Scan(&pilihan)
				switch pilihan {
				case 1:
					C[i].akomodasi = 0
					kategori = "akomodasi"
				case 2:
					C[i].transportasi = 0
					kategori = "transportasi"
				case 3:
					C[i].makanan = 0
					kategori = "makanan"
				case 4:
					C[i].hiburan = 0
					kategori = "hiburan"
				default:
					fmt.Println("Kategori tidak valid!")
					return
				}
				fmt.Printf(sp+"Berhasil menghapus untuk %s di kategori %s\n", A[i].nama, kategori)
				return
			}
		}
	} else {
		fmt.Print(sp + "id tidak di temukan")
	}
}
func lihatSemua(A tabPengguna, B tabPengeluaran, n int) { // sudah rapih
	fmt.Println(sp + "╔════════╦════════════════╦════════════╦══════════════╦══════════╦══════════╗")
	fmt.Println(sp + "║   ID   ║      Nama      ║ Akomodasi  ║ Transportasi ║ Makanan  ║ Hiburan  ║")
	fmt.Println(sp + "╠════════╬════════════════╬════════════╬══════════════╬══════════╬══════════╣")

	for i := 0; i < n; i++ {
		fmt.Printf("║  %-4d ║ %-14s ║ %-10d ║ %-12d ║ %-8d ║ %-8d ║\n",
			A[i].id, A[i].nama, B[i].akomodasi, B[i].transportasi, B[i].makanan, B[i].hiburan)
	}

	fmt.Println(sp + "╚════════╩════════════════╩════════════╩══════════════╩══════════╩══════════╝")
}
func laporanAkhir(A tabPengguna, B tabPengeluaran, pax int, budget int) {
	var total int
	var selisih int
	var presentase float64
	var keterangan string
	var rekomendasi string

	fmt.Println()
	fmt.Println(sp + "══════════════════════════════════════════════════════════════════════════════")
	fmt.Println(sp + "                               LAPORAN AKHIR")
	fmt.Println(sp + "══════════════════════════════════════════════════════════════════════════════")
	fmt.Println(sp + "╔════╦════════════════╦════════════════════╦════════════╦════════════╦══════════╦══════════════════════════╗")
	fmt.Println(sp + "║ ID ║ Nama           ║ Total Pengeluaran  ║ Budget     ║ Selisih    ║ Persen   ║ Keterangan               ║")
	fmt.Println(sp + "╠════╬════════════════╬════════════════════╬════════════╬════════════╬══════════╬══════════════════════════╣")

	for i := 0; i < pax; i++ {
		total = B[i].akomodasi + B[i].transportasi + B[i].makanan + B[i].hiburan
		selisih = budget - total
		presentase = float64(total) / float64(budget) * 100

		switch {
		case presentase < 70:
			keterangan = "Masih jauh dari limit"
		case presentase >= 70 && presentase < 90:
			keterangan = "Cukup besar, hati-hati"
		case presentase >= 90 && presentase < 100:
			keterangan = "Hampir habis, atur baik-baik"
		case presentase >= 100:
			keterangan = "Habis! Perlu evaluasi"
		}

		fmt.Printf("║ %-2d ║ %-14s ║ %-18d ║ %-10d ║ %-10d ║ %-8.2f%% ║ %-24s ║\n",
			A[i].id, A[i].nama, total, budget, selisih, presentase, keterangan)
	}

	fmt.Println(sp + "╚════╩════════════════╩════════════════════╩════════════╩════════════╩══════════╩══════════════════════════╝")

	// Tabel rekomendasi personal
	fmt.Println()
	fmt.Println(sp + "══════════════════════════════════════════════════════════════════════════════")
	fmt.Println(sp + "                         REKOMENDASI PERSONAL UNTUK PENGGUNA")
	fmt.Println(sp + "══════════════════════════════════════════════════════════════════════════════")
	fmt.Println(sp + "╔════╦════════════════╦══════════════════════════════════════════════════════╗")
	fmt.Println(sp + "║ ID ║ Nama           ║ Rekomendasi                                          ║")
	fmt.Println(sp + "╠════╬════════════════╬══════════════════════════════════════════════════════╣")

	for i := 0; i < pax; i++ {
		total = B[i].akomodasi + B[i].transportasi + B[i].makanan + B[i].hiburan
		presentase = float64(total) / float64(budget) * 100

		switch {
		case presentase < 70:
			rekomendasi = "Coba beli oleh-oleh unik atau ikut tur lokal dadakan "
		case presentase >= 70 && presentase < 90:
			rekomendasi = "Pertimbangkan makan malam hemat atau cari diskon oleh-oleh "
		case presentase >= 90 && presentase < 100:
			rekomendasi = "Fokus ke destinasi gratis seperti taman kota atau street art "
		case presentase >= 100:
			rekomendasi = "Catat pengeluaran dan kurangi Belanja, utamakan kebutuhan pokok "
		}

		fmt.Printf("║ %-2d ║ %-14s ║ %-60s ║\n", A[i].id, A[i].nama, rekomendasi)
	}

	fmt.Println(sp + "╚════╩════════════════╩══════════════════════════════════════════════════════╝")
}

func cariPengeluaranSeq(A tabPengguna, B tabPengeluaran, n int) {
	var i int
	var maxakomodasi, maxtransportasi, maxmakanan, maxhiburan int

	for i = 0; i < n-1; i++ {
		if B[i].akomodasi > B[i+1].akomodasi {
			maxakomodasi = B[i].akomodasi
		} else {
			maxakomodasi = B[i+1].akomodasi
		}
	}
	i = 0
	for maxakomodasi != B[i].akomodasi {
		i++
	}

	fmt.Printf(sp+"Pengeluaran terbanyak pada kategori akomodasi sebesar %d oleh %s\n", maxakomodasi, A[i].nama)
	for i = 0; i < n-1; i++ {
		if B[i].transportasi > B[i+1].transportasi {
			maxtransportasi = B[i].transportasi
		} else {
			maxtransportasi = B[i+1].transportasi
		}
	}

	i = 0

	for maxtransportasi != B[i].transportasi {
		i++
	}

	fmt.Printf(sp+"Pengeluaran terbanyak pada kategori transportasi sebesar %d oleh %s\n", maxtransportasi, A[i].nama)
	for i = 0; i < n-1; i++ {
		if B[i].makanan > B[i+1].makanan {
			maxmakanan = B[i].makanan
		} else {
			maxmakanan = B[i+1].makanan
		}
	}

	i = 0

	for maxmakanan != B[i].makanan {
		i++
	}

	fmt.Printf(sp+"Pengeluaran terbanyak pada kategori makanan sebesar %d oleh %s\n", maxmakanan, A[i].nama)
	for i = 0; i < n-1; i++ {
		if B[i].hiburan > B[i+1].hiburan {
			maxhiburan = B[i].hiburan
		} else {
			maxhiburan = B[i+1].hiburan
		}
	}

	i = 0

	for maxhiburan != B[i].hiburan {
		i++
	}

	fmt.Printf(sp+"Pengeluaran terbanyak pada hiburan akomodasi sebesar %d oleh %s\n", maxhiburan, A[i].nama)
}

func cariPengeluaranBinary(A tabPengguna, B tabPengeluaran, n int) {
	var kategori string
	var target int
	var dataPengeluaran [100]int
	var urutUser [100]string

	fmt.Println(sp + "╔════╦═══════════════════╗")
	fmt.Println(sp + "║ No ║     Kategori      ║")
	fmt.Println(sp + "╠════╬═══════════════════╣")
	fmt.Println(sp + "║ 1  ║ Akomodasi         ║")
	fmt.Println(sp + "║ 2  ║ Transportasi      ║")
	fmt.Println(sp + "║ 3  ║ Makanan           ║")
	fmt.Println(sp + "║ 4  ║ Hiburan           ║")
	fmt.Println(sp + "╚════╩═══════════════════╝")
	fmt.Print(sp + "Pilih kategori yang ingin dicari (1-4): ")
	fmt.Scan(&kategori)

	switch kategori {
	case "1":
		fmt.Print(sp + "Masukkan total pengeluaran akomodasi yang dicari: ")
		fmt.Scan(&target)
		for i := 0; i < n; i++ {
			dataPengeluaran[i] = B[i].akomodasi
			urutUser[i] = A[i].nama
		}
	case "2":
		fmt.Print(sp + "Masukkan total pengeluaran transportasi yang dicari: ")
		fmt.Scan(&target)
		for i := 0; i < n; i++ {
			dataPengeluaran[i] = B[i].transportasi
			urutUser[i] = A[i].nama
		}
	case "3":
		fmt.Print(sp + "Masukkan total pengeluaran makanan yang dicari: ")
		fmt.Scan(&target)
		for i := 0; i < n; i++ {
			dataPengeluaran[i] = B[i].makanan
			urutUser[i] = A[i].nama
		}
	case "4":
		fmt.Print(sp + "Masukkan total pengeluaran hiburan yang dicari: ")
		fmt.Scan(&target)
		for i := 0; i < n; i++ {
			dataPengeluaran[i] = B[i].hiburan
			urutUser[i] = A[i].nama
		}
	default:
		fmt.Println(sp + "Kategori tidak valid.")
		return
	}

	// Sorting dulu
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if dataPengeluaran[i] > dataPengeluaran[j] {
				dataPengeluaran[i], dataPengeluaran[j] = dataPengeluaran[j], dataPengeluaran[i]
				urutUser[i], urutUser[j] = urutUser[j], urutUser[i]
			}
		}
	}

	// Binary Search
	low := 0
	high := n - 1
	mid := -1
	found := false

	for low <= high {
		mid = (low + high) / 2
		if dataPengeluaran[mid] == target {
			found = true
			break
		} else if dataPengeluaran[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		fmt.Println(sp + "Pengeluaran ditemukan pada beberapa user:")
		// Cek ke kiri
		i := mid
		for i >= 0 && dataPengeluaran[i] == target {
			fmt.Printf(sp+"Nama: %s | Total: %d\n", urutUser[i], dataPengeluaran[i])
			i--
		}
		// Cek ke kanan
		i = mid + 1
		for i < n && dataPengeluaran[i] == target {
			fmt.Printf(sp+"Nama: %s | Total: %d\n", urutUser[i], dataPengeluaran[i])
			i++
		}
	} else {
		fmt.Println(sp + "Tidak ada user dengan pengeluaran tersebut.")
	}
}

func urutkanPengeluaran(daftarPengguna *tabPengguna, daftarPengeluaran *tabPengeluaran, jumlahPengguna int) {
	var i, j, idxMin int
	var penggunaSementara pengguna
	var pengeluaranSementara Pengeluaran
	var totalSekarang, totalBanding, totalMin, totalSebelumnya int

	var pilihan int
	fmt.Println("╔════╦══════════════════════════════════════════╗")
	fmt.Println("║ No ║              Jenis Pengurutan            ║")
	fmt.Println("╠════╬══════════════════════════════════════════╣")
	fmt.Println("║ 1  ║ Minimal ke Maksimal (Selection Sort)     ║")
	fmt.Println("║ 2  ║ Maksimal ke Minimal (Insertion Sort)     ║")
	fmt.Println("╚════╩══════════════════════════════════════════╝")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		// Selection Sort (Ascending)
		for i = 0; i < jumlahPengguna-1; i++ {
			idxMin = i
			totalMin = daftarPengeluaran[i].akomodasi + daftarPengeluaran[i].transportasi + daftarPengeluaran[i].makanan + daftarPengeluaran[i].hiburan

			for j = i + 1; j < jumlahPengguna; j++ {
				totalBanding = daftarPengeluaran[j].akomodasi + daftarPengeluaran[j].transportasi + daftarPengeluaran[j].makanan + daftarPengeluaran[j].hiburan

				if totalBanding < totalMin {
					idxMin = j
					totalMin = totalBanding
				}
			}

			if idxMin != i {
				// Tukar pengguna
				penggunaSementara = daftarPengguna[i]
				daftarPengguna[i] = daftarPengguna[idxMin]
				daftarPengguna[idxMin] = penggunaSementara

				// Tukar pengeluaran
				pengeluaranSementara = daftarPengeluaran[i]
				daftarPengeluaran[i] = daftarPengeluaran[idxMin]
				daftarPengeluaran[idxMin] = pengeluaranSementara
			}
		}
	case 2:
		// Insertion Sort (Descending)
		for i = 1; i < jumlahPengguna; i++ {
			j = i
			for j > 0 {
				totalSekarang = daftarPengeluaran[j].akomodasi + daftarPengeluaran[j].transportasi + daftarPengeluaran[j].makanan + daftarPengeluaran[j].hiburan
				totalSebelumnya = daftarPengeluaran[j-1].akomodasi + daftarPengeluaran[j-1].transportasi + daftarPengeluaran[j-1].makanan + daftarPengeluaran[j-1].hiburan

				if totalSekarang > totalSebelumnya {
					// Tukar pengguna
					penggunaSementara = daftarPengguna[j]
					daftarPengguna[j] = daftarPengguna[j-1]
					daftarPengguna[j-1] = penggunaSementara

					// Tukar pengeluaran
					pengeluaranSementara = daftarPengeluaran[j]
					daftarPengeluaran[j] = daftarPengeluaran[j-1]
					daftarPengeluaran[j-1] = pengeluaranSementara

					j--
				} else {
					break
				}
			}
		}
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println()

	fmt.Println("╔════╦════════════════╦════════════════════╗")
	fmt.Println("║ ID ║ Nama           ║ Total Pengeluaran  ║")
	fmt.Println("╠════╬════════════════╬════════════════════╣")
	for i = 0; i < jumlahPengguna; i++ {
		total := (*daftarPengeluaran)[i].akomodasi + (*daftarPengeluaran)[i].transportasi + (*daftarPengeluaran)[i].makanan + (*daftarPengeluaran)[i].hiburan
		fmt.Printf("║ %-2d ║ %-14s ║ %-18d ║\n", (*daftarPengguna)[i].id, (*daftarPengguna)[i].nama, total)
	}
	fmt.Println("╚════╩════════════════╩════════════════════╝")

}
