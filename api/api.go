package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"main/dictionary"
	"net/http"
)

func ListHandler(res http.ResponseWriter, req *http.Request, d *dictionary.Dictionary) {
	words, _ := d.List()
	for word, entry := range words {
		_, err := fmt.Fprintf(res, "[%s] %-10s %s\n", entry.Date, word, entry.Definition)
		if err != nil {
			return
		}
	}
}

func GetHandler(res http.ResponseWriter, req *http.Request, d *dictionary.Dictionary) {
	wordParam := chi.URLParam(req, "word")
	words, _ := d.Get(wordParam)
	for word, entry := range words {
		_, err := fmt.Fprintf(res, "[%s] %-10s %s\n", entry.Date, word, entry.Definition)
		if err != nil {
			return
		}
	}
}

func AddHandler(res http.ResponseWriter, req *http.Request, d *dictionary.Dictionary) {
	var entryData struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}
	err := json.NewDecoder(req.Body).Decode(&entryData)
	if err != nil {
		fmt.Printf("Error occured -> %s", err)
	}

	d.Add(entryData.Word, entryData.Definition)
}

func RemoveHandler(res http.ResponseWriter, req *http.Request, d *dictionary.Dictionary) {
	var bodyData struct {
		Word string `json:"word"`
	}
	err := json.NewDecoder(req.Body).Decode(&bodyData)
	if err != nil {
		fmt.Printf("Error occured -> %s", err)
	}
	d.Remove(bodyData.Word)
}

func RunServer(d *dictionary.Dictionary) {
	router := chi.NewRouter()

	router.Get("/list", func(res http.ResponseWriter, req *http.Request) {
		ListHandler(res, req, d)
	})
	router.Get("/get/{word}", func(res http.ResponseWriter, req *http.Request) {
		GetHandler(res, req, d)
	})
	router.Post("/add", func(res http.ResponseWriter, req *http.Request) {
		AddHandler(res, req, d)
	})
	router.Delete("/remove", func(res http.ResponseWriter, req *http.Request) {
		RemoveHandler(res, req, d)
	})

	err := http.ListenAndServe(":9000", router)
	if err != nil {
		fmt.Printf("Error occured -> %s\n", err)
	}
}
