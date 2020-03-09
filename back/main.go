package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

var c redis.Conn

var schemaLoader gojsonschema.JSONLoader

type req struct {
	Number int `json:"number"`
}
type res struct {
	Result int `json:"result"`
}

type myerror struct {
	Error string `json:"error"`
	Type  int    `json:"type"`
}

type conf struct {
	DbHost  string
	DbName  int
	AppHost string
}

func increment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req{Number: 10})

}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kek", increment).Methods("GET")
	r.HandleFunc("/ready", ready).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
