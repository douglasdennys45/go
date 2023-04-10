package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type DBTransaction interface {
	OpenTransaction(context.Context) (*sql.Tx, error)
	Commit(*sql.Tx) error
	Rollback(*sql.Tx) error
	GetTX() *sql.Tx
	GetContext() context.Context
}

type DBTransactionController struct {
	ControllerInterface
	DBTransaction
}

func NewDBTransactionController(controller ControllerInterface, transaction DBTransaction) ControllerInterface {
	return &DBTransactionController{controller, transaction}
}

func (controller *DBTransactionController) perform(ctx *fiber.Ctx) Response {
	tx, err := controller.OpenTransaction(context.Background())
	if err != nil {
		return Response{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
			Path:    "POST /users",
		}
	}
	response := controller.ControllerInterface.perform(ctx)
	if response.Status == fiber.StatusInternalServerError {
		err := controller.Rollback(tx)
		if err != nil {
			return Response{
				Message: err.Error(),
				Status:  fiber.StatusInternalServerError,
				Path:    "POST /users",
			}
		}
		return response
	}
	err = controller.Commit(tx)
	if err != nil {
		return Response{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
			Path:    "POST /users",
		}
	}
	return response
}
