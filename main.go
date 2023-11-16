package main

import (
	"github.com/kubestaff/golearning/home"
	"github.com/kubestaff/golearning/setting"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

func main() {	
	opts := server.Options{
		Port: 34567,
	}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	s.Handle("/", home.HandleHome)
	s.Handle("/me", user.HandleMe)
	s.Handle("/setting", setting.HandleReadSetting)

	s.Start()
}