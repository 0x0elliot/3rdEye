package router

import (
	"log"
	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRuleRoutes() {

}

func HandleCreateRule(c* fiber.Ctx) error {
	type CreateRuleRequest struct {
		Rule string `json:"rule"`
	}

	var data CreateRuleRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	

}
