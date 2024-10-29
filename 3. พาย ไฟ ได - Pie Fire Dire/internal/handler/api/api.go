package api

import (
	pb "beef/internal/pd"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	beef pb.BeefServiceClient
}

func NewHandler(beef pb.BeefServiceClient) *handler {
	return &handler{
		beef: beef,
	}
}

const idleTimeout = 5 * time.Second

func (h *handler) Init() *fiber.App {
	router := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	NewHandlerBeef(h.beef, router)

	return router
}
