package dictionary

import (
	"fmt"
	"log"
	"main/database"
	"time"
)

type Entry struct {
	Definition string
	Date       string
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

func (d *Dictionary) Add(word string, definition string) {
	db := database.ConnectDB()
	defer db.Close()
	d.entries[word] = Entry{
		definition,
		time.Now().Format("02-01-2006 - 15:04:05"),
	}
	query := "INSERT INTO dictionary(word, definition, created_at) VALUES (?, ?, ?)"
	_, err := db.Exec(query, word, definition, time.Now().Format("02-01-2006 - 15:04:05"))
	if err != nil {
		log.Println("Failed to add to database : ", err)
	}
}

func (d *Dictionary) Get(word string) (map[string]Entry, error) {
	db := database.ConnectDB()
	defer db.Close()
	query := "SELECT word, definition, created_at FROM dictionary WHERE word=?"
	rows, err := db.Query(query, word)
	if err != nil {
		log.Printf("Unable to get the word = %s : %s\n", word, err)
	}
	words := make(map[string]Entry)
	for rows.Next() {
		var word, definition, created_at string
		err := rows.Scan(&word, &definition, &created_at)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan the line : ", err)
		}
		words[word] = Entry{
			Definition: definition,
			Date:       created_at,
		}
	}
	return words, nil
}

func (d *Dictionary) List() (map[string]Entry, error) {
	db := database.ConnectDB()
	defer db.Close()
	query := "SELECT word,definition, created_at FROM dictionary"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Println("There is an error : ", err)
	}
	words := make(map[string]Entry)
	for rows.Next() {
		var word, definition, created_at string
		err := rows.Scan(&word, &definition, &created_at)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan the line : ", err)
		}
		words[word] = Entry{
			Definition: definition,
			Date:       created_at,
		}
	}
	return words, nil
}

func (d *Dictionary) Remove(word string) {
	db := database.ConnectDB()
	defer db.Close()
	query := "DELETE FROM dictionary WHERE word=?"
	_, err := db.Exec(query, word)
	if err != nil {
		log.Printf("Unable to delete rows where word = %s : %s\n", word, err)
	}
}
