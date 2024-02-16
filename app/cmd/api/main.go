package main

import (
	"database/sql"
	"fmt"
	"log"
	"noteapp/internal/repositories/note"
	"noteapp/internal/server"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	dburl = os.Getenv("DB_URL")
)

func main() {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	noteRepository := note.NewNoteRepository(db)

	server := server.New(noteRepository)
	server.RegisterFiberRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err = server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
