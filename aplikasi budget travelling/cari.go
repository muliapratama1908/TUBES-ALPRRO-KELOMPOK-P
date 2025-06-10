package main

import "fmt"

func cariSequential(data [MAX]Pengeluaran, n int, key string) {
	found := false
	for i := 0; i < n; i++ {
		if data[i].Kategori == key {
			fmt.Printf("Ditemukan: %s - Rp%.2f\n", data[i].Kategori, data[i].Jumlah)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariBinary(data [MAX]Pengeluaran, n int, key string) {
	kiri := 0
	kanan := n - 1
	found := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah].Kategori == key {
			fmt.Printf("Ditemukan: %s - Rp%.2f\n", data[tengah].Kategori, data[tengah].Jumlah)
			found = true
			break
		} else if data[tengah].Kategori < key {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan (pastikan sudah diurutkan).")
	}
}
