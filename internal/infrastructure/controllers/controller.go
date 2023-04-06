package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"requestId,omitempty"`
	Message   string      `json:"message,omitempty"`
	Status    int         `json:"status,omitempty"`
	Error     string      `json:"error,omitempty"`
	Path      string      `json:"path,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}

type ControllerInterface interface {
	perform(ctx *fiber.Ctx) Response
	Handle(ControllerInterface) Response
}

type Controller struct {
	Ctx *fiber.Ctx
}

func (controller *Controller) Handle(ctrl ControllerInterface) Response {
	return ctrl.perform(controller.Ctx)
}
