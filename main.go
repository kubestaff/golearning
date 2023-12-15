package main

import (
	"log"

	"github.com/kubestaff/golearning/db"
	"github.com/kubestaff/golearning/home"
	"github.com/kubestaff/golearning/setting"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

func main() {
	dbConn, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Migrate(dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}

	opts := server.Options{
		Port: 34567,
	}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	settingHandler := setting.Handler{
		DbConnection: dbConn,
	}

	userHandler := user.Handler{
		DbConnection: dbConn,
	}

	homeHandler := home.Handler{
		DbConnection: dbConn,
	}

	s.Handle("/", homeHandler.HandleHome)
	s.Handle("/me", userHandler.HandleMe)
	s.Handle("/setting", settingHandler.HandleReadSetting)

	s.Start()
}
