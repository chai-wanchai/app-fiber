package controller

import (
	"iot/config"
	driver "iot/dependency"
	"iot/internal/dto"
	"iot/internal/service"

	"github.com/gofiber/fiber/v2"
)

type IoTController interface {
	GetData(ctx *fiber.Ctx) error
	GetSumData(ctx *fiber.Ctx) error
	RandomData(ctx *fiber.Ctx) error
}

type controller struct {
	Config      config.Config
	Connections driver.Connections
	Service     service.IoTService
}

func NewController(config config.Config, conn driver.Connections) IoTController {
	svc := service.NewIoTService(conn.SqlORM)
	return controller{
		Config:      config,
		Connections: conn,
		Service:     svc,
	}
}

func (ctr controller) GetData(ctx *fiber.Ctx) error {
	q := dto.QueryPayload{}
	if err := ctx.QueryParser(&q); err != nil {
		return err
	}
	data, err := ctr.Service.GetData(ctx.UserContext(), q)
	if err != nil {
		return err
	}
	return ctx.JSON(data)
}
func (ctr controller) GetSumData(ctx *fiber.Ctx) error {
	q := dto.QueryPayload{}
	if err := ctx.QueryParser(&q); err != nil {
		return err
	}
	data, err := ctr.Service.GetSumData(ctx.UserContext(), q)
	if err != nil {
		return err
	}
	return ctx.JSON(data)
}

func (ctr controller) RandomData(ctx *fiber.Ctx) error {
	err := ctr.Service.RandomData(ctx.UserContext())
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]string{
		"status": "sucess",
	})
}
