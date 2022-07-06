package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SarthakRanaTiket/projectname/router"
)


func main(){
	r := router.Router()
	fmt.Println("Starring server on PORT : 8080...")

	log.Fatal(http.ListenAndServe(":8080",r))
}