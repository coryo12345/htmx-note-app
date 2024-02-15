package routes

import (
	"bytes"
	"noteapp/cmd/web"
	"noteapp/internal/repositories/note"

	"github.com/gofiber/fiber/v2"
)

const (
	BAD_REQUEST = "Bad Request"
	GENERIC_ERR = "Something went wrong, please try again later"
)

func HomePageHandler(c *fiber.Ctx, nr note.NoteRepository) error {
	c.Set("Content-Type", "text/html")

	notes, err := nr.GetAllNotes()
	if err != nil {
		return err
	}

	component := web.NoteForm(notes)
	buf := new(bytes.Buffer)
	component.Render(c.Context(), buf)
	c.Status(fiber.StatusOK).SendString(buf.String())

	return nil
}
