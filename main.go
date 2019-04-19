package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yashdalfthegray/color-info/handlers"
	"github.com/yashdalfthegray/color-info/utils"

	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.NewStatusHandler(utils.GetServerID())).Methods("GET")
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	loggedRouter := gHandlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", loggedRouter)
	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
