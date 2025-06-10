package main

import "fmt"

func tambahPengeluaran(data *[MAX]Pengeluaran, n *int) {
    if *n >= MAX {
        fmt.Println("Data penuh!")
        return
    }

    var p Pengeluaran
    fmt.Print("Masukkan kategori: ")
    fmt.Scan(&p.Kategori)
    fmt.Print("Masukkan jumlah: ")
    fmt.Scan(&p.Jumlah)

    data[*n] = p
    *n++

    fmt.Println("Pengeluaran berhasil ditambahkan.")
}

func ubahPengeluaran(data *[MAX]Pengeluaran, n int) {
    var idx int
    fmt.Print("Masukkan nomor data yang mau diubah: ")
    fmt.Scan(&idx)

    if idx < 1 || idx > n {
        fmt.Println("Nomor tidak valid.")
        return
    }

    fmt.Print("Kategori baru: ")
    fmt.Scan(&data[idx-1].Kategori)
    fmt.Print("Jumlah baru: ")
    fmt.Scan(&data[idx-1].Jumlah)
    fmt.Println("Data berhasil diubah.")
}

func hapusPengeluaran(data *[MAX]Pengeluaran, n *int) {
    var idx int
    fmt.Print("Masukkan nomor data yang mau dihapus: ")
    fmt.Scan(&idx)

    if idx < 1 || idx > *n {
        fmt.Println("Nomor tidak valid.")
        return
    }

    for i := idx - 1; i < *n-1; i++ {
        data[i] = data[i+1]
    }
    *n--
    fmt.Println("Data berhasil dihapus.")
}
