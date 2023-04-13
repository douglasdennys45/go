package controllers

import (
	goresponse "github.com/douglasdennys45/go-response"
	"github.com/douglasdennys45/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type userDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserController struct {
	Controller
	services.AddUserInterface
}

func NewAddUserController(controller Controller, addUser services.AddUserInterface) ControllerInterface {
	return &AddUserController{controller, addUser}
}

func (controller *AddUserController) perform(ctx *fiber.Ctx) goresponse.Response {
	user := new(userDTO)
	if err := ctx.BodyParser(user); err != nil {
		r := goresponse.NewResponse(nil, "", err.Error(), fiber.StatusInternalServerError, "", "POST /users")
		return r.Response()
	}
	if err := controller.Execute(user.Name, user.Email, user.Password); err != nil {
		r := goresponse.NewResponse(nil, "", err.Error(), fiber.StatusInternalServerError, "", "POST /users")
		return r.Response()
	}
	r := goresponse.NewResponse(nil, "", "User created successfully", fiber.StatusInternalServerError, "", "POST /users")
	return r.Response()
}
