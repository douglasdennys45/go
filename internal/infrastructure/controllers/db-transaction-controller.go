package controllers

import "github.com/gofiber/fiber/v2"

type DBTransaction interface {
	Begin() error
	Commit() error
	Rollback() error
}

type DBTransactionController struct {
	ControllerInterface
	DBTransaction
}

func (controller *DBTransactionController) perform(ctx *fiber.Ctx) Response {
	if err := controller.Begin(); err != nil {
		return Response{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
			Path:    "POST /users",
		}
	}
	response := controller.ControllerInterface.perform(ctx)
	if response.Status == fiber.StatusInternalServerError {
		if err := controller.Rollback(); err != nil {
			return Response{
				Message: err.Error(),
				Status:  fiber.StatusInternalServerError,
				Path:    "POST /users",
			}
		}
		return response
	}
	if err := controller.Commit(); err != nil {
		return Response{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
			Path:    "POST /users",
		}
	}
	return response
}
