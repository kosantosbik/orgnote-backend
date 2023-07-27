package handlers

import (
	"moonbrain/app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// UploadFiles godoc
// @Summary      Upload files
// @Description  Upload files.
// @Tags         files
// @Accept       json
// @Produce      json
// @Param        files   formData      []string  true  "files"
// @Success      200  {object}  any
// @Failure      400  {object}  HttpError[any]
// @Failure      404  {object}  HttpError[any]
// @Failure      500  {object}  HttpError[any]
// @Router       /files/upload  [post]
func (h FilesHandlers) UploadFiles(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Error().Err(err).Msg("files handler: upload files: could not get multipart form")
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Can't parse multipart form data", nil))
	}
	files := form.File["files"]
	err = h.fileService.UploadFiles(files)
	if err != nil {
		log.Error().Err(err).Msg("files handler: upload files: could not upload files")
		return c.Status(http.StatusInternalServerError).JSON(NewHttpError[any]("Can't upload files", nil))
	}
	return c.Status(http.StatusOK).JSON(nil)

}

type FilesHandlers struct {
	fileService *services.FileService
}

func RegisterFileHandler(app fiber.Router, fileService *services.FileService, authMiddleware func(*fiber.Ctx) error) {
	fileHandlers := &FilesHandlers{
		fileService: fileService,
	}
	// TODO: master add middleware for max size.
	app.Post("/files/upload", authMiddleware, fileHandlers.UploadFiles)
	// app.Delete("/files/:id", noteHandlers.GetNote)
}
