package main

const MAX = 100

type Pengeluaran struct {
    Kategori string
    Jumlah   float64
}

var dataPengeluaran [MAX]Pengeluaran
var jumlahData int
var totalBudget float64
