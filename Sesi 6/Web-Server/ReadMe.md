# Simple Go Web Server

Program ini adalah contoh sederhana dari server web menggunakan bahasa pemrograman Go (Golang).

## Persyaratan
- Go (Golang) harus diinstal. [Download di sini](https://golang.org/dl/)

## Penggunaan

1. Pindah ke direktori proyek:

    ```bash
    cd <Nama Folder>
    ```

2. Jalankan aplikasi dengan perintah:

    ```bash
    go run main.go
    ```

3. Buka browser dan akses [http://localhost:8080](http://localhost:8080) untuk melihat pesan sambutan "Hallo World".

4. Untuk melihat daftar karyawan, akses [http://localhost:8080/employees](http://localhost:8080/employees).

5. Untuk menambahkan karyawan baru, kirimkan permintaan POST ke [http://localhost:8080/employee](http://localhost:8080/employee) dengan data formulir (nama, usia, dan divisi).

## Contoh Penggunaan

### Menambahkan Karyawan Baru

```bash
curl -X POST -d "name=John&age=30&division=Sales" http://localhost:8080/employee
