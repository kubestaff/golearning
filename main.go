package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kubestaff/golearning/db"
	"github.com/kubestaff/golearning/user"
)

func main() {
	dbConn, err := db.CreateDataBase()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
	}))

	userHandler := user.Handler{
		DbConnection: dbConn,
	}

	r.GET("/users", userHandler.HandleUsers)
	r.DELETE("/users", userHandler.HandleDeleteUser)
	r.POST("/users", userHandler.HandleChangeUser)

	//r.GET("/setting", settingsHandler.HandleReadSetting)

	r.Run(":34567")
}
