package routes

import (
	"github.com/boblancer/Post-API/routes"
	"github.com/labstack/echo"
	"github.com/boblancer/Post-API/controllers"
)

func userRoutes(route *echo.Echo) {
	user := &controllers.UserController{}
	err = user.ListUserAll()
	//route.GET("/users", user.ListUserAll)
	//route.GET("/users/:id", user.FindByID)
	//route.POST("/users",user.Add)

}
