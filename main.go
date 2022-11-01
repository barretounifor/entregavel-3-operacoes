package main

import (
	"net/http"

	httpFunctions "entregavel-3/httpFunctions"

	"github.com/gorilla/mux"
)

func main() {
	httpHandler := mux.NewRouter()

	httpHandler.HandleFunc("/sum", httpFunctions.HandleSum)
	httpHandler.HandleFunc("/sub", httpFunctions.HandleSub)
	httpHandler.HandleFunc("/div", httpFunctions.HandleDiv)
	httpHandler.HandleFunc("/mult", httpFunctions.HandleMult)

	http.ListenAndServe(":3000", httpHandler)

}
