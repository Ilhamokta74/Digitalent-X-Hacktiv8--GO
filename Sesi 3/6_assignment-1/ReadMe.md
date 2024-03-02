# Biodata Teman CLI

Program sederhana menggunakan Go untuk menampilkan informasi tentang teman-teman di kelas berdasarkan nomor absen.

## Prasyarat

Pastikan Anda sudah menginstal Go di komputer Anda. Jika belum, Anda dapat mengunduhnya [di sini](https://golang.org/dl/).

## Cara Menjalankan

1. Buka terminal atau command prompt.

2. Pindah ke direktori yang berisi file `biodata.go`.

3. Jalankan perintah berikut:

    ```bash
    go run biodata.go <nomor_absen>
    ```

    Gantilah `<nomor_absen>` dengan nomor absen teman Anda.

    Contoh:

    ```bash
    go run biodata.go 1
    ```

4. Program akan menampilkan informasi tentang teman dengan nomor absen yang Anda tentukan.

## Contoh Output

```bash
Data Teman dengan Absen 1
===================
Nama: Andi
Alamat: Jl. Merdeka 12
Pekerjaan: Mahasiswa
Alasan Memilih Kelas Golang: Ingin belajar bahasa pemrograman yang modern dan powerful
