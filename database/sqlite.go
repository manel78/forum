package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE If NOT EXIST user (
		id INTEGER PRIMARY KEY,
		username TEXT UNIQUE,
		password TEXT
	)`)

	if err != nil {
		log.Fatal(err)
	}

	// category
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// messages
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY,
		user_id INTEGER,
		category_id INTEGER,
		content TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (category_id) REFERENCES categories(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tables créées avec succès !")
}
