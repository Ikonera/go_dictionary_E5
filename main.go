package main

import (
	"bufio"
	"main/api"
	"main/cli"
	"main/database"
	"main/dictionary"
	"os"
)

func main() {
	d := dictionary.New()
	r := bufio.NewReader(os.Stdin)

	database.InitDB()

	go cli.RunCli(d, r)
	api.RunServer(d)
}
