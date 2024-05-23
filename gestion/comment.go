package data

import (
	"database/sql"
	"log"
)

// structure commentaire
type Comment struct {
	ID      int
	PostID  int
	Content string
}

// cree un nvx commentaire
func CreateComment(db *sql.DB, postID int, content string) error {
	_, err := db.Exec("INSERT INTO comments (post_id, content) VALUES (?, ?)", postID, content)
	if err != nil {
		log.Println("Erreur lors de la création d'un commentaire:", err)
		return err
	}
	return nil
}

// recupere tous les commentaires associées au post id
func GetCommentsByPostID(db *sql.DB, postID int) ([]Comment, error) {
	var comments []Comment
	rows, err := db.Query("SELECT id, post_id, content FROM comments WHERE post_id = ?", postID)
	if err != nil {
		log.Println("Erreur lors de la récupération des commentaires pour un post:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content); err != nil {
			log.Println("Erreur lors de l'analyse des résultats de la requête pour récupérer les commentaires pour un post:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		log.Println("Erreur lors de l'itération sur les résultats de la requête pour récupérer les commentaires pour un post:", err)
		return nil, err
	}
	return comments, nil
}

// met à jour le contenu d'un commentaire
func UpdateCommentContent(db *sql.DB, commentID int, newContent string) error {
	_, err := db.Exec("UPDATE comments SET content = ? WHERE id = ?", newContent, commentID)
	if err != nil {
		log.Println("Erreur lors de la mise à jour du contenu du commentaire:", err)
		return err
	}
	return nil
}

// supprime un commentaire de la base de données
func DeleteComment(db *sql.DB, commentID int) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		log.Println("Erreur lors de la suppression du commentaire:", err)
		return err
	}
	return nil
}
