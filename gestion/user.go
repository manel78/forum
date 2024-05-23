package data

import (
	"database/sql"
	"log"
)

// User représente la structure d'un utilisateur
type User struct {
	ID       int
	Username string
	Password string
}

// CreateUser crée un nouvel utilisateur dans la base de données
func CreateUser(db *sql.DB, username, password string) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Println("Erreur lors de la création d'un utilisateur:", err)
		return err
	}
	return nil
}

// GetUserByID récupère les informations d'un utilisateur à partir de son ID
func GetUserByID(db *sql.DB, userID int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println("Erreur lors de la récupération d'un utilisateur par ID:", err)
		return nil, err
	}
	return &user, nil
}

// UpdateUserPassword met à jour le mot de passe d'un utilisateur
func UpdateUserPassword(db *sql.DB, userID int, newPassword string) error {
	_, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", newPassword, userID)
	if err != nil {
		log.Println("Erreur lors de la mise à jour du mot de passe de l'utilisateur:", err)
		return err
	}
	return nil
}

// DeleteUser supprime un utilisateur de la base de données
func DeleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		log.Println("Erreur lors de la suppression de l'utilisateur:", err)
		return err
	}
	return nil
}
