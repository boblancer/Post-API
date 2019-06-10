package routes

import (
	"github.com/labstack/echo"
	"post-api/controllers"
)



func userRoutes(route *echo.Echo) {
	userController := &controllers.UserController{}
	route.GET("/users", userController.ListUserAll)
	route.GET("/users/:id", userController.FindByID)
	route.POST("/users", userController.Add)

}
