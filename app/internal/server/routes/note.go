package routes

import (
	"bytes"
	"noteapp/cmd/web"
	"noteapp/internal/repositories/note"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx, nr note.NoteRepository) error {
	status, err := nr.Health()
	if err != nil {
		c.Status(fiber.StatusOK).SendString("note service is unhealthy")
		return err
	}
	c.Status(fiber.StatusOK).SendString(status)
	return nil
}

func NewNoteHandler(c *fiber.Ctx, nr note.NoteRepository) error {
	// Parse form data
	if err := c.BodyParser(c); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(BAD_REQUEST)
	}

	// Get the name from the form data
	value := c.FormValue("value")
	author := c.FormValue("author")

	newNote := note.Note{Value: value, Author: author}

	note, err := nr.AddNote(newNote)
	if err != nil {
		return err
	}

	// Render the component
	component := web.NoteItem(note)
	buf := new(bytes.Buffer)
	component.Render(c.Context(), buf)

	c.Status(fiber.StatusOK).SendString(buf.String())
	return nil
}

func DeleteNoteHandler(c *fiber.Ctx, nr note.NoteRepository) error {
	// Parse form data
	if err := c.BodyParser(c); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(BAD_REQUEST)
	}

	// Get the id from the form data
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(GENERIC_ERR)
	}

	err = nr.DeleteNote(id)
	if err != nil {
		return err
	}

	c.SendStatus(fiber.StatusOK)
	return nil
}
