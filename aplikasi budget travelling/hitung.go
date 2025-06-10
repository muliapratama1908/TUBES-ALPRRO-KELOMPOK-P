package main

import "fmt"

func hitungTotal(data [MAX]Pengeluaran, n int) float64 {
    var total float64
    for i := 0; i < n; i++ {
        total += data[i].Jumlah
    }
    return total
}

func cetakLaporan(data [MAX]Pengeluaran, n int) {
    total := hitungTotal(data, n)
    sisa := totalBudget - total

    fmt.Println("\n--- Laporan ---")
    for i := 0; i < n; i++ {
        fmt.Printf("%d. %s - Rp%.2f\n", i+1, data[i].Kategori, data[i].Jumlah)
    }
    fmt.Printf("Total Pengeluaran: Rp%.2f\n", total)
    fmt.Printf("Sisa Budget: Rp%.2f\n", sisa)
}
