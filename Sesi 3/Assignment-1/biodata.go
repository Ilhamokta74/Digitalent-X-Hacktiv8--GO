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
	// Mengambil argumen dari baris perintah
	args := os.Args

	// Pastikan argumen yang diberikan sesuai dengan ekspektasi
	if len(args) != 2 {
		fmt.Println("Gunakan: go run biodata.go <nomor_absen>")
		return
	}

	// Konversi argumen nomor absen ke dalam bentuk integer
	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat")
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
		{
			Nama:         "Dewi",
			Alamat:       "Jl. Gatotkaca 8",
			Pekerjaan:    "UI/UX Designer",
			AlasanGolang: "Menggunakan Golang untuk pengembangan antarmuka pengguna yang efisien",
		},
		{
			Nama:         "Eko",
			Alamat:       "Jl. Pahlawan 45",
			Pekerjaan:    "Embedded Systems Engineer",
			AlasanGolang: "Golang cocok untuk pengembangan sistem terbenam yang handal",
		},
		{
			Nama:         "Fani",
			Alamat:       "Jl. Kapten Ismail 3",
			Pekerjaan:    "Game Developer",
			AlasanGolang: "Menyukai kinerja tinggi Golang untuk pengembangan game",
		},
		{
			Nama:         "Gita",
			Alamat:       "Jl. Pemuda 27",
			Pekerjaan:    "Network Engineer",
			AlasanGolang: "Menggunakan Golang untuk pengembangan aplikasi jaringan yang efisien",
		},
	}

	// Menampilkan data teman berdasarkan absen
	if absen <= 0 || absen > len(dataTeman) {
		fmt.Println("Absen tidak valid, Program dihentikan.")
		return
	}

	teman := dataTeman[absen-1]
	fmt.Println("Data Teman dengan Absen : ", absen)
	fmt.Println("===================")
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan Memilih Kelas Golang:", teman.AlasanGolang)
}
