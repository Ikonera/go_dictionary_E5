package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Failed to open the database : ", err)
	}
	return db
}

func InitDB() {
	dbFile := "database.db"
	if _, err := os.Stat(dbFile); err == nil {
		log.Println("Database file exists !")
	} else if os.IsNotExist(err) {
		log.Println("Database file does not exist, creating one...")
		_, err := os.Create(dbFile)
		if err != nil {
			log.Println("Database file error during creation : ", err)
		} else {
			log.Println("Database file created !")
		}
	}

	db := ConnectDB()
	defer db.Close()

	query := "CREATE TABLE IF NOT EXISTS dictionary(" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL," +
		"word TEXT NOT NULL," +
		"definition TEXT NOT NULL," +
		"created_at TEXT NOT NULL" +
		")"
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Dictionary table creation failed : ", err)
	}

}
