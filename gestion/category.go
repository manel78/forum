package data

import (
	"database/sql"
	"log"
)

// structure categorie
type Category struct {
	ID   int
	Name string
}

// CreateCategory crée une nouvelle catégorie dans la base de données
func CreateCategory(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", name)
	if err != nil {
		log.Println("Erreur lors de la création d'une catégorie:", err)
		return err
	}
	return nil
}

// GetAllCategories récupère toutes les catégories depuis la base de données
func GetAllCategories(db *sql.DB) ([]Category, error) {
	var categories []Category
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Println("Erreur lors de la récupération de toutes les catégories:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println("Erreur lors de l'analyse des résultats de la requête pour récupérer toutes les catégories:", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		log.Println("Erreur lors de l'itération sur les résultats de la requête pour récupérer toutes les catégories:", err)
		return nil, err
	}
	return categories, nil
}

func DeleteCategory(db *sql.DB, categoryID int) error {
	_, err := db.Exec("DELETE FROM categories WHERE id = ?", categoryID)
	if err != nil {
		log.Println("Erreur lors de la suppression de la catégorie:", err)
		return err
	}
	return nil
}
