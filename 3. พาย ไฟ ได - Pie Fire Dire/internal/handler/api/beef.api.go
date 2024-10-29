package api

import (
	pb "beef/internal/pd"

	"github.com/gofiber/fiber/v2"
)

type handlerBeef struct {
	client pb.BeefServiceClient
}

func NewHandlerBeef(client pb.BeefServiceClient, router *fiber.App) {
	h := handlerBeef{
		client: client,
	}

	router.Get("/beef/summary", h.handlerFindBeef)
}

func (h *handlerBeef) handlerFindBeef(c *fiber.Ctx) error {

	res, err := h.client.FindBeef(c.Context(), &pb.FindBeefRequest{})

	if err != nil {
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
