## Raw SQL

## Setup
* Jalankan di terminal (sesuaikan nama project kalian) `go mod init namaproject`
* download driver mysql. jalankan `go get -u github.com/go-sql-driver/mysql`
* buat file main.go
* mulai coding

## DB
* buat dulu databasenya
* buat tablenya

## Create ENV
* buat file dengan nama `.env`
* tulis configurasi connection string kalian disana
```
#contoh: 
export CONNECTION_DB='root:password@tcp(127.0.0.1:3306)/db_name'
```
* buat file dengan nama `.gitignore`
* tambahkan `.env` di gitignore tsb

## Menjalankan app golang dengan env
Untuk menjalankan app golang yang memanfaatkan env, maka kita harus load file env terlebih dahulu.
<br>
note: jalankan di terminal bash atau linux.

```
source .env
```
lalu, ketika file .env sudah di load, maka jalankan project golang kalan
```
go run main.go
```