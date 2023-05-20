package main

import (
	"database/sql"

	db2 "github.com/chseidler/hexagonal-go-example/adapters/db"
	"github.com/chseidler/hexagonal-go-example/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}