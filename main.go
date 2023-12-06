package main

import (
	"fmt"
	"log"

	"github.com/kubestaff/golearning/db"
	"github.com/kubestaff/golearning/home"
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

	s.Handle("/", home.HandleHome)
	s.Handle("/me10", user.HandleMe10 )
	s.Handle("/setting", setting.HandleReadSetting)
	
	s.Start()

}
	



