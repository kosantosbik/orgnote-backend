package handlers

import (
	"moonbrain/app/models"
	"moonbrain/app/services"
	"net/http"

	_ "moonbrain/app/docs"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type NoteHandlers struct {
	noteService *services.NoteService
}

// TODO: master wait when swago will support generics :(

type SuccessGetNotesResponse struct {
	Notes []models.Note `json:"notes"`
}

// GetNote godoc
// @Summary      Get note
// @Description  get note by id
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Note ID"
// @Success      200  {object}  HttpResponse[models.PublicNote, any]
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/{id}  [get]
func (h *NoteHandlers) GetNote(c *fiber.Ctx) error {
	noteID := c.Params("id")

	ctxUser := c.Locals("user")

	var userID string

	if ctxUser != nil {
		userID = ctxUser.(*models.User).ID.Hex()
	}

	notes, err := h.noteService.GetNote(noteID, userID)
	if err != nil {
		log.Info().Err(err).Msg("note handler: get note: get by id")
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Couldn't get note, something went wrong", nil))
	}
	if notes == nil {
		return c.Status(http.StatusNotFound).JSON(NewHttpResponse[any, any](nil, nil))
	}
	return c.Status(http.StatusOK).JSON(NewHttpResponse[*models.PublicNote, any](notes, nil))
}

// DeleteNotes godoc
// @Summary      Delete notes
// @Description  Mark notes as deleted by provided list of ids
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        ids   body     []string  true  "List of ids of deleted notes"
// @Success      200  {object}  HttpResponse[any, any]
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes [delete]
func (h *NoteHandlers) DeleteNotes(c *fiber.Ctx) error {
	notesIDs := []string{}
	err := c.BodyParser(&notesIDs)
	if err != nil {
		log.Info().Err(err).Msg("note handler: delete notes: body parser")
		return c.Status(http.StatusBadRequest).JSON(NewHttpError[any]("Couldn't parse body, something went wrong", nil))
	}
	h.noteService.DeleteNotes(notesIDs)
	return nil
}

// GetNote godoc
// @Summary      Get notes
// @Description  Get all notes with optional filter
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        userId       query  string  false  "User ID"
// @Param        searchText   query  string  false  "Search text"
// @Param        limit        query  int  true  "Limit for pagination"
// @Param        offset       query  int  true  "Offset for pagination"
// @Success      200  {object}  HttpResponse[[]models.PublicNote, models.Pagination]
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/  [get]
func (h *NoteHandlers) GetNotes(c *fiber.Ctx) error {
	defaultLimit := int64(10)
	defaultOffset := int64(0)

	filter := new(models.NoteFilter)

	if err := c.QueryParser(filter); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewHttpError("Incorrect input query", err))
	}

	ctxUser := c.Locals("user")

	includePrivateNotes := filter.UserID != nil && ctxUser != nil && ctxUser.(*models.User).ID.Hex() == *filter.UserID

	if filter.Limit == nil {
		filter.Limit = &defaultLimit
	}

	if filter.Offset == nil {
		filter.Offset = &defaultOffset
	}

	paginatedNotes, err := h.noteService.GetNotes(includePrivateNotes, *filter)
	if err != nil {
		log.Info().Err(err).Msgf("note handler: get notes: get %v", err)
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Couldn't get notes, something went wrong", nil))
	}
	return c.Status(http.StatusOK).JSON(
		NewHttpResponse(paginatedNotes.Data, models.Pagination{
			Limit:  paginatedNotes.Limit,
			Offset: paginatedNotes.Offset,
			Total:  paginatedNotes.Total,
		}))
}

// CreateNote godoc
// @Summary      Create note
// @Description  Create note
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        note       body  CreatingNote  true  "Note model"
// @Success      200  {object}  any
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/  [post]
func (h *NoteHandlers) CreateNote(c *fiber.Ctx) error {
	note := new(CreatingNote)

	if err := c.BodyParser(note); err != nil {
		log.Info().Err(err).Msg("note handler: post note: parse body")
		return c.Status(fiber.StatusInternalServerError).JSON(NewHttpError("Can't parse body", err))
	}

	author := c.Locals("user").(*models.User)
	n := mapCreatingNoteToNote(*note)
	n.AuthorID = author.ID.Hex()
	err := h.noteService.CreateNote(n)

	if err != nil {
		log.Info().Err(err).Msgf("note handler: post note: create %v", err)
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Can't create note", nil))
	}
	return c.Status(http.StatusOK).JSON(nil)
}

// UpserNotes godoc
// @Summary      Upsert notes
// @Description  Bulk update or insert notes
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        notes body []CreatingNote true "List of created notes"
// @Success      200  {object}  any
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/bulk-upsert  [put]
func (h *NoteHandlers) UpsertNotes(c *fiber.Ctx) error {
	notesForCreate := []CreatingNote{}

	if err := c.BodyParser(&notesForCreate); err != nil {
		log.Error().Err(err).Msg("note handler: upsert notes: parse body")
		return c.Status(http.StatusBadRequest).JSON(NewHttpError[any]("Couldn't parse body, something went wrong", nil))
	}

	user := c.Locals("user").(*models.User)
	notes := mapCreatingNotesToNotes(notesForCreate)
	err := h.noteService.BulkCreateOrUpdate(user.ID.Hex(), notes)
	if err != nil {
		log.Warn().Msgf("note handlers: save notes: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Can't create notes", nil))
	}
	return c.Status(http.StatusOK).JSON(nil)
}

// GetNoteGraph godoc
// @Summary      Get notes graph
// @Description  Return graph model with links between connected notes
// @Tags         notes
// @Accept       json
// @Produce      json
// @Success      200  {object}  handlers.HttpResponse[models.NoteGraph, any]
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/graph  [get]
func (h *NoteHandlers) GetNoteGraph(c *fiber.Ctx) error {
	ctxUser := c.Locals("user")

	if ctxUser == nil {
		return c.Status(http.StatusNotFound).Send(nil)
	}

	graph, err := h.noteService.GetNoteGraph(ctxUser.(*models.User).ID.Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Couldn't get note graph", nil))
	}

	return c.Status(http.StatusOK).JSON(NewHttpResponse[*models.NoteGraph, any](graph, nil))
}

// SyncNotes godoc
// @Summary      Syncronize notes
// @Description  Syncronize notes with specific timestamp
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        timestamp   path      string  true  "Timestamp of the last syncronization"
// @Success      200  {object}
// @Failure      400  {object}  HttpResponse[SyncNotesData, any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /notes/sync  [post]
// func (h *NodeHandlers) SyncNotes(c *fiber.Ctx) error {
// 	ctxUser := c.Locals("user")
// 	return c.Status(http.StatusOK).JSON(NewHttpResponse[]())
// }

func RegisterNoteHandler(app fiber.Router, noteService *services.NoteService, authMiddleware func(*fiber.Ctx) error) {
	noteHandlers := &NoteHandlers{
		noteService: noteService,
	}
	app.Get("/notes/graph", authMiddleware, noteHandlers.GetNoteGraph)
	app.Get("/notes/:id", noteHandlers.GetNote)
	app.Get("/notes", noteHandlers.GetNotes)
	// app.Post("/sync", noteHandelrs.SyncNotes)
	app.Post("/notes", authMiddleware, noteHandlers.CreateNote)
	app.Put("/notes/bulk-upsert", authMiddleware, noteHandlers.UpsertNotes)
	app.Delete("/notes", authMiddleware, noteHandlers.DeleteNotes)
}
