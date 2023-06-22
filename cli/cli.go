package cli

import (
	"bufio"
	"fmt"
	"main/dictionary"
	"strings"
)

func AddCommand(d *dictionary.Dictionary, r *bufio.Reader) {
	fmt.Print("Enter a new word : ")
	word, _ := r.ReadString('\n')
	word = strings.TrimSpace(word)
	fmt.Print("Enter the word definition : ")
	definition, _ := r.ReadString('\n')
	definition = strings.TrimSpace(definition)
	d.Add(word, definition)
}

func GetCommand(d *dictionary.Dictionary, r *bufio.Reader) {
	fmt.Print("Enter a word to search : ")
	word, _ := r.ReadString('\n')
	word = strings.TrimSpace(word)
	words, _ := d.Get(word)
	for word, entry := range words {
		fmt.Printf("[%s] %-10s %s\n", entry.Date, word, entry.Definition)
	}
}

func RemoveCommand(d *dictionary.Dictionary, r *bufio.Reader) {
	fmt.Print("Enter a word to remove : ")
	word, exists := r.ReadString('\n')
	if exists != nil {
		fmt.Printf("The word %s doesn't exists.", word)
	}
	word = strings.TrimSpace(word)
	d.Remove(word)
}

func ListCommand(d *dictionary.Dictionary) {
	words, _ := d.List()
	for word, entry := range words {
		fmt.Printf("[%s] %-10s %s\n", entry.Date, word, entry.Definition)
	}
}

func RunCli(d *dictionary.Dictionary, r *bufio.Reader) {
	for {
		fmt.Print("\nAvailable commands -> [add/get/remove/list/exit] $ ")
		command, _ := r.ReadString('\n')
		command = strings.TrimSpace(command)
		switch command {
		case "add":
			AddCommand(d, r)
		case "get":
			GetCommand(d, r)
		case "list":
			ListCommand(d)
		case "remove":
			RemoveCommand(d, r)
		case "exit":
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
