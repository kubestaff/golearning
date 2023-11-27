package main

import (
	"log"

	"github.com/kubestaff/golearning/db"
	"github.com/kubestaff/golearning/setting"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
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

	opts := server.Options{
		Port: 34567,
	}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	settingsHandler := setting.Handler{
		DbConnection: dbConn,
	}

	userHandler := user.Handler{
		DbConnection: dbConn,
	}

	s.HandleJSON("/users", userHandler.HandleUsers)
	s.HandleJSON("/user-delete", userHandler.HandleDeleteUser)
	s.HandleJSON("/user-change", userHandler.HandleChangeUser)

	s.Handle("/setting", settingsHandler.HandleReadSetting)

	s.Start()
}
