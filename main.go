package main

import (
	"log"

	"github.com/P47H4N/shafar_foundation_api/cmd"
	"github.com/P47H4N/shafar_foundation_api/internals/database"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	cnf, err := cmd.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config.")
		return
	}
	db, err := database.InitDB(cnf)
	if err != nil {
		log.Fatalln("Database Connection Failed.", err)
		return
	}
	err = cmd.Start(router, db)
	if err != nil {
		log.Fatalln("Unable to Start.", err)
		return
	}
	router.Run(":" + cnf.Port)
}
