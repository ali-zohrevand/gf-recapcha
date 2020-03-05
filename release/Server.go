package main

import (
	"github.com/ali-zohrevand/gf-recapcha/release/routers"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":8080", n))

}
