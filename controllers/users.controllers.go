package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"post-api/services"
)

type ErrorObj struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields"`
}
type UserController struct {
}
//route.GET("/users", user.ListUserAll)
//route.GET("/users/:id", user.FindByID)
//route.POST("/users",user.Add)


func (UserController) ListUserAll(c echo.Context) error {
	fmt.Println("print all user")
	userService := &services.UserService{}
	users, err := userService.FindAll()
	if err != nil {
		return c.JSON(http.StatusOK, "I don't have value from database")
	}
	return c.JSON(http.StatusOK, users)
}

func (UserController) FindByID(c echo.Context) error {
	fmt.Println("find by id")
	return c.JSON(http.StatusOK, nil)
}

func (UserController) Add(c echo.Context) error {
	fmt.Println("add user")
	return c.JSON(http.StatusOK, nil)
}