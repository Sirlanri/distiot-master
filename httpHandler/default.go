package httphandler

import (
	"github.com/Sirlanri/distiot-master/server/config"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func IrisInit() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	app.Use(crs)
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	master := app.Party("/master").AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost)
	{

	}
	var portStr = config.Config.HttpPort
	app.Run(iris.Addr(":" + portStr))
}
