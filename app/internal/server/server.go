package server

import (
	"noteapp/internal/repositories/note"
	"noteapp/internal/server/routes"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	note.NoteRepository
}

func New(noteRepository note.NoteRepository) *FiberServer {
	server := &FiberServer{
		App:            fiber.New(),
		NoteRepository: noteRepository,
	}

	return server
}

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Static("/static", "./cmd/web/static")

	s.App.Get("/", func(c *fiber.Ctx) error { return routes.HomePageHandler(c, s.NoteRepository) })

	// note routes
	s.App.Get("/note/health", func(c *fiber.Ctx) error { return routes.Health(c, s.NoteRepository) })
	s.App.Post("/note", func(c *fiber.Ctx) error {
		return routes.NewNoteHandler(c, s.NoteRepository)
	})
	s.App.Delete("/note", func(c *fiber.Ctx) error { return routes.DeleteNoteHandler(c, s.NoteRepository) })
}
