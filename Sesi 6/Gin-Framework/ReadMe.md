# Simple Car API using Gin

Proyek ini adalah contoh sederhana API untuk manajemen data mobil menggunakan framework Gin di Golang.

## Persyaratan
- Go (Golang) harus diinstal. [Download di sini](https://golang.org/dl/)

## Penggunaan

1. Pindah ke direktori proyek:

    ```bash
    cd <Nama Folder>
    ```

2. Download depedensi dengan perintah:

    ```bash
    go mod download
    ```


3. Jalankan aplikasi dengan perintah:

    ```bash
    go run main.go
    ```

4. Untuk Menambahkan Mobil Baru, kirimkan permintaan POST ke [http://localhost:8080/cars](http://localhost:8080/cars) dengan data formulir (brand, model, dan price).

5. Untuk Memperbarui Data Mobil, kirimkan permintaan PUT ke [http://localhost:8080/cars/{carID}](http://localhost:8080/cars/{carID}) dengan data formulir (brand, model, dan price).

6. Mendapatkan Semua Data Mobil, kirimkan permintaan GET ke [http://localhost:8080/cars](http://localhost:8080/cars)

7. Untuk Mendapatkan Data Mobil Berdasarkan ID, kirimkan permintaan GET ke [http://localhost:8080/cars/{carID}](http://localhost:8080/cars/{carID})

8. Untuk Menghapus Data Mobil, kirimkan permintaan DELETE ke [http://localhost:8080/cars/{carID}](http://localhost:8080/cars/{carID})

## Contoh Penggunaan

###  Menambahkan Mobil Baru

    curl -X POST -H "Content-Type: application/json" -d '{"brand":"Toyota","model":"Camry","price":30000}' http://localhost:8080/cars
    

### Memperbarui Data Mobil

    curl -X PUT -H "Content-Type: application/json" -d '{"brand":"Honda","model":"Accord","price":35000}' http://localhost:8080/cars/{carID}
    

### Mendapatkan Semua Data Mobil

    curl http://localhost:8080/cars
    
    
### Mendapatkan Data Mobil Berdasarkan ID

    curl http://localhost:8080/cars/{carID}
    

### Menghapus Data Mobil

    curl -X DELETE http://localhost:8080/cars/{carID}
