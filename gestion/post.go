package data

import (
	"database/sql"
	"log"
)

// Post représente la structure d'une publication
type Post struct {
	ID      int
	Title   string
	Content string
}

// CreatePost crée une nouvelle publication dans la base de données
func CreatePost(db *sql.DB, title, content string) error {
	_, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		log.Println("Erreur lors de la création d'une publication:", err)
		return err
	}
	return nil
}

// GetPostByID récupère les informations d'une publication à partir de son ID
func GetPostByID(db *sql.DB, postID int) (*Post, error) {
	var post Post
	err := db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		log.Println("Erreur lors de la récupération d'une publication par ID:", err)
		return nil, err
	}
	return &post, nil
}

// UpdatePostContent met à jour le contenu d'une publication
func UpdatePostContent(db *sql.DB, postID int, newContent string) error {
	_, err := db.Exec("UPDATE posts SET content = ? WHERE id = ?", newContent, postID)
	if err != nil {
		log.Println("Erreur lors de la mise à jour du contenu de la publication:", err)
		return err
	}
	return nil
}

// DeletePost supprime une publication de la base de données
func DeletePost(db *sql.DB, postID int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", postID)
	if err != nil {
		log.Println("Erreur lors de la suppression de la publication:", err)
		return err
	}
	return nil
}
