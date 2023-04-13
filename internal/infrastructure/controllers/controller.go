package controllers

import (
	goresponse "github.com/douglasdennys45/go-response"
	"github.com/gofiber/fiber/v2"
)

type ControllerInterface interface {
	perform(ctx *fiber.Ctx) goresponse.Response
	Handle(ControllerInterface) goresponse.Response
}

type Controller struct {
	Ctx *fiber.Ctx
}

func (controller *Controller) Handle(ctrl ControllerInterface) goresponse.Response {
	return ctrl.perform(controller.Ctx)
}
