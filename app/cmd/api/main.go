package main

import (
	"fmt"
	"noteapp/internal/repositories/note"
	"noteapp/internal/server"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	noteRepository := note.NewNoteRepository()

	server := server.New(noteRepository)
	server.RegisterFiberRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
