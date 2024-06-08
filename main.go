package main

import (
	"Hospital/routers"
	"Hospital/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
