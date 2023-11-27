package main

import (
	"fmt"
	
	"github.com/kubestaff/golearning/home"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

func main() {
    provider := user.Provider{}
	err := provider.SaveUsers()
	if err != nil {
		fmt.Println(err)
		return
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
	
	s.Start()
}
	



