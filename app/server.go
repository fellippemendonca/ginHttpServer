package app

import (
	"fmt"
	"github.com/fellippemendonca/ginHttpServer/app/routes"
	"net/http"
)

func Init() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")	
	router := routes.Init()
	http.ListenAndServe(":8080", router)
}
