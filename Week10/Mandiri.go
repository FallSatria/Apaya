package main

import "fmt"

const NMAX = 100

type produk struct {
    name  string
    penjualan int
}

type dataproduct [NMAX]produk

func carimax(A dataproduct, n int) (string, int) {
    max := A[0].penjualan
    maxname := A[0].name

    for i := 1; i < n; i++ {
        if A[i].penjualan > max {
            max = A[i].penjualan
            maxname = A[i].name
        }
    }

    return maxname, max
}

func carimin(A dataproduct, n int) (string, int) {
    min := A[0].penjualan
    minname := A[0].name

    for i := 1; i < n; i++ {
        if A[i].penjualan < min {
            min = A[i].penjualan
            minname = A[i].name
        }
    }

    return minname, min
}

func ratarata(A dataproduct, n int) float64 {
    total := 0
    for i := 0; i < n; i++ {
        total += A[i].penjualan
    }

    return float64(total) / float64(n)
}

func main() {
    var n int
    var A dataproduct

    fmt.Print("Masukkan jumlah produk: ")
    fmt.Scan(&n)

        for i := 0; i < n; i++ {
        fmt.Println("\nProduk ke-", i+1)

        fmt.Print("Nama Produk : ")
        fmt.Scan(&A[i].name)

        fmt.Print("Jumlah Terjual : ")
        fmt.Scan(&A[i].penjualan)
    }

    maxname, max := carimax(A, n)
    minname, min := carimin(A, n)
    rata := ratarata(A, n)

    fmt.Println("\n===== HASIL ANALISIS PENJUALAN =====")
    fmt.Println("Produk Terlaris     :", maxname)
    fmt.Println("Penjualan Maksimum  :", max)

    fmt.Println("Produk Terendah     :", minname)
    fmt.Println("Penjualan Minimum   :", min)

    fmt.Println("Selisih Max-Min     :", max-min)
    fmt.Printf("Rata-rata Penjualan : %.0f\n", rata)
}