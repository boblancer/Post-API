package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type UserController struct {
}


func (UserController) ListUserAll(c echo.Context) error {
	fmt.Println("print all user")
	return c.JSON(http.StatusOK, nil)
}