package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama, Alamat, Pekerjaan, AlasanGolang string
}

func main() {
	// Mendapatkan argumen dari terminal
	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argumen tidak valid:", err)
		return
	}

	// Data teman-teman di kelas
	dataTeman := []Teman{
		{
			Nama:         "Andi",
			Alamat:       "Jl. Merdeka 12",
			Pekerjaan:    "Mahasiswa",
			AlasanGolang: "Ingin belajar bahasa pemrograman yang modern dan powerful",
		},
		{
			Nama:         "Budi",
			Alamat:       "Jl. Sudirman 17",
			Pekerjaan:    "Web Developer",
			AlasanGolang: "Ingin meningkatkan skill programming untuk pengembangan web",
		},
		{
			Nama:         "Cici",
			Alamat:       "Jl. Diponegoro 21",
			Pekerjaan:    "Data Scientist",
			AlasanGolang: "Ingin mempelajari tools dan library Golang untuk data science",
		},
	}

	// Menampilkan data teman berdasarkan absen
	if absen <= 0 || absen > len(dataTeman) {
		fmt.Println("Absen tidak valid")
		return
	}

	teman := dataTeman[absen-1]
	fmt.Println("Data Teman dengan Absen", absen)
	fmt.Println("===================")
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan Memilih Kelas Golang:", teman.AlasanGolang)
}
