package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Nome  string
	Price float64
}

func NewProduct(nome string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Nome:  nome,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/product")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Geladeira", 12.78)

	err = insertProduct(db, product)

	if err != nil {
		panic(err)
	}

	product.Price = 1500.50

	err = updateProduct(db, product)

	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id,nome,price) values (?,?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Nome, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set nome = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Nome, product.Price, product.Id)

	if err != nil {
		return err
	}

	return nil

}
