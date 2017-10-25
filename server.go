package main

import (
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	c "github.com/korovaisdead/go-simple-membership/config"
	"github.com/korovaisdead/go-simple-membership/controllers"
	"net/http"
	"os"
)

var (
	config *c.Configuration
)

func main() {
	config, err := c.GetConfig()
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}

	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/register", controllers.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/authenticate", controllers.LoginHandler).Methods(http.MethodPost)

	if err := http.ListenAndServe(config.Web.Port, handlers.LoggingHandler(os.Stdout, router)); err != nil {
		panic(err)
	}

	select {}
}
