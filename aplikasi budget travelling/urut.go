package main

import "fmt"

func urutJumlahSelection(data *[MAX]Pengeluaran, n int) {
	for i := 0; i < n-1; i++ {
		minP := i
		for j := i + 1; j < n; j++ {
			if data[j].Jumlah < data[minP].Jumlah {
				minP = j
			}
		}
		data[i], data[minP] = data[minP], data[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jumlah (Selection Sort).")
}

func urutKategoriInsertion(data *[MAX]Pengeluaran, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Kategori > key.Kategori {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan kategori (Insertion Sort).")
}

func menuUrut(data *[MAX]Pengeluaran, n int) {
	var pilih int
	fmt.Println("1. Urutkan berdasarkan Jumlah (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Kategori (Insertion Sort)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		urutJumlahSelection(data, n)
	} else if pilih == 2 {
		urutKategoriInsertion(data, n)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
