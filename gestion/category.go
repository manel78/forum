package data

import (
	"database/sql"
	"log"
)

type Category struct {
	ID   int
	Name string
}

func CreateCategory(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", name)
	if err != nil {
		log.Println("Erreur lors de la création d'une catégorie:", err)
		return err
	}
	return nil
}
