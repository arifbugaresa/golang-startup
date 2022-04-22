package handler

import (
	"github.com/gin-gonic/gin"
	"golang-startup/helper"
	"golang-startup/user"
	"net/http"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	c.JSON(http.StatusOK, helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter))

}
