package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/maxdevelopment/go_hello/chat"
)

const (
	serverPort = ":8000"
	wsPath = "/ws"
)

func main() {
	wsServer := chat.WsServer()
	go wsServer.Listen()
	go wsServer.BroadcastTest()

	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc(wsPath, func(writer http.ResponseWriter, request *http.Request) {
		chat.StartWs(wsServer, writer, request)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(serverPort, router))
}

func GetIndex(writer http.ResponseWriter, request *http.Request) {
	indexFile, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Write(indexFile)
}