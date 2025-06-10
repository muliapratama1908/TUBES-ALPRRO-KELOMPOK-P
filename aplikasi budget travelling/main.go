package main

import "fmt"

func main() {
	fmt.Print("Masukkan total budget: ") // menampilkan teks awal kepada pengguna untuk memasukkan nilai
	fmt.Scan(&totalBudget)               // menyimpan data ke dalam variabel totalbudget

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Ubah Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Lihat Semua")
		fmt.Println("5. Hitung Total & Sisa")
		fmt.Println("6. Cari Pengeluaran")
		fmt.Println("7. Urutkan Data")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahPengeluaran(&dataPengeluaran, &jumlahData)
		} else if pilih == 2 {
			ubahPengeluaran(&dataPengeluaran, jumlahData)
		} else if pilih == 3 {
			hapusPengeluaran(&dataPengeluaran, &jumlahData)
		} else if pilih == 4 {
			cetakLaporan(dataPengeluaran, jumlahData)
		} else if pilih == 5 {
			cetakLaporan(dataPengeluaran, jumlahData)
		} else if pilih == 6 {
			var cari string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&cari)
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			var metode int
			fmt.Print("Pilih metode: ")
			fmt.Scan(&metode)
			if metode == 1 {
				cariSequential(dataPengeluaran, jumlahData, cari)
			} else {
				urutKategoriInsertion(&dataPengeluaran, jumlahData)
				cariBinary(dataPengeluaran, jumlahData, cari)
			}
		} else if pilih == 7 {
			menuUrut(&dataPengeluaran, jumlahData)
		} else if pilih == 8 {
			fmt.Println("Sampai jumpa!")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
