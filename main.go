package main

import (
	"log"

	"chatsapi/internal/http"
)

func main() {
	app := http.Http()

	log.Fatal(app.Listen(":80"))
}
