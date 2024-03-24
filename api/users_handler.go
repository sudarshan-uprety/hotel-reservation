package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sudarshan-uprety/hotel-reservation/db"
	"github.com/sudarshan-uprety/hotel-reservation/types"
)

type UserHandler struct {
	UserStore db.UserStore
}

func NewUserHandler(UserStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: UserStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.UserStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	c.JSON(insertedUser)
	return nil
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var updateParams types.UpdateUserParams

	if err := c.BodyParser(&updateParams); err != nil {
		return err
	}
	if errors := updateParams.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	if err := c.BodyParser(&updateParams); err != nil {
		return err
	}

	updated, err := h.UserStore.UpdateUser(c.Context(), userID, &updateParams)
	if err != nil {
		return err
	}
	return c.JSON(updated)

}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	deleted, err := h.UserStore.DeleteUser(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.JSON(deleted)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.UserStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
