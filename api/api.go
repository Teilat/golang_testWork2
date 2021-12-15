package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"src/golang_testWork2/api/handler"
	"src/golang_testWork2/vault"
)

func Init(cache *vault.Vault)  {
	handlers := api.New(cache)
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./web/static/")))

	router.Path("/view").Queries("key", "{key}").HandlerFunc(handlers.HandlerView())
	router.Path("/view").HandlerFunc(handlers.HandlerView())

	router.Path("/add").Queries("key", "{key}", "value", "{value}", "duration", "{duration}").HandlerFunc(handlers.HandlerAdd())
	router.Path("/add").Queries("key", "{key}", "value", "{value}").HandlerFunc(handlers.HandlerAdd())
	router.Path("/add").HandlerFunc(handlers.HandlerAdd())

	router.Path("/flush").HandlerFunc(handlers.HandlerFlush())

	router.Path("/remove").Queries("key", "{key}").HandlerFunc(handlers.HandlerRemove())
	router.Path("/remove").HandlerFunc(handlers.HandlerRemove())

	log.Println("started")
	log.Fatal(http.ListenAndServe(":80", router))
}
