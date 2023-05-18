package db

import (
	"database/sql"

	"github.com/chseidler/hexagonal-go-example/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {

}
