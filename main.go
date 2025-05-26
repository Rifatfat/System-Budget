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
