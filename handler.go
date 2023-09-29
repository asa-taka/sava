package main

import (
	"path/filepath"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/afero"
)

type appConfig struct {
	dataDir               string
	corsAllowOriginRegexp regexp.Regexp
	corsAllowHeaders      string
}

type rootHandler struct {
	fs afero.Fs
}

func newApp(conf *appConfig) *fiber.App {
	h := &rootHandler{
		fs: afero.NewBasePathFs(afero.NewOsFs(), conf.dataDir),
	}
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool { return true },
	}))
	app.Get("/u/:email/avatar", h.getAvatar)
	app.Put("/u/:email/avatar", h.setAvatar)
	return app
}

// userDirPath is /u/:email.
func userDirPath(email string) string {
	return filepath.Join("/u", email)
}

// avatarFilePath is /u/:email/avatar.
func avatarFilePath(email string) string {
	return filepath.Join(userDirPath(email), "avatar")
}

func (h *rootHandler) getAvatar(c *fiber.Ctx) error {
	email := c.Params("email")
	b, err := afero.ReadFile(h.fs, avatarFilePath(email))
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/octet-stream")
	return c.Send(b)
}

func (h *rootHandler) setAvatar(c *fiber.Ctx) error {
	email := c.Params("email")
	if err := h.fs.MkdirAll(userDirPath(email), 0755); err != nil {
		return err
	}
	if err := afero.WriteFile(h.fs, avatarFilePath(email), c.Body(), 0644); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}
