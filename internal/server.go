package internal

import (
	"iot/config"
	driver "iot/dependency"
	"iot/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitFiber(cfg config.Config) error {
	dep := driver.NewConnection(cfg)
	ctr := controller.NewController(cfg, dep.GetConnection())
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())
	app.Get("/api/data", ctr.GetData)
	app.Get("/api/sum", ctr.GetSumData)
	app.Get("/api/random", ctr.RandomData)
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status": "ok",
		})
	})

	err := app.Listen(":" + config.GetConfig().HTTP.Port)
	if err != nil {
		return err
	}
	return nil
}
