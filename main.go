package main

import (
	"crud-go/config"
	"crud-go/handlers"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/kategori/create", handlers.CreateKategori)
	http.HandleFunc("/produk/create", handlers.CreateProduk)
	http.HandleFunc("/produk/read", handlers.GetProduk)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server berjalan di http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
