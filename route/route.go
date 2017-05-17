package route

import (
	"github.com/labstack/echo"
	"MyServer/handler"
)

var e = echo.New()

func Init()  {
	e.POST("/user/register", handler.RegisterHandler)
	e.POST("/user/login", handler.LoginHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
