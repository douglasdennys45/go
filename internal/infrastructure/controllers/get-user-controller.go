package controllers

import (
	goresponse "github.com/douglasdennys45/go-response"
	"github.com/douglasdennys45/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type GetUserController struct {
	Controller
	services.GetUserInterface
}

func NewGetUserController(controller Controller, getUser services.GetUserInterface) ControllerInterface {
	return &GetUserController{controller, getUser}
}

func (controller *GetUserController) perform(ctx *fiber.Ctx) goresponse.Response {
	id := ctx.Params("id")
	users, err := controller.Execute(id)
	if err != nil {
		r := goresponse.NewResponse(nil, "", err.Error(), fiber.StatusNotFound, "", "POST /users/"+id)
		return r.Response()
	}
	r := goresponse.NewResponse(users, "", "User found successfully", fiber.StatusOK, "", "POST /users/"+id)
	return r.Response()
}
