package routes

import ("github.com/labstack/echo")

func init() {
	e := echo.New()
	userRoutes(e)

}