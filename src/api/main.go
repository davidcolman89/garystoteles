package main

import (
	"github.com/davidcolman89/garystoteles/src/api/config"
	"github.com/davidcolman89/garystoteles/src/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	conf := config.NewConfig("Config")

	router := gin.Default()
	router.GET("/garystoteles", routes.GetPhrases)
	router.GET("/ping", routes.Ping)
	router.Run(conf.ServerAddr)

}