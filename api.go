package main

import (
    "log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
    "strconv"
    "github.com/bornehill/mygoapi/datastore"
    "time"
)

var (
	menus datastore.MenuData
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	menus = &datastore.Menus{}
	menus.Initialize()
}

func main() {
	r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api menu")
	})
    //api.HandleFunc("/", get).Methods(http.MethodGet)
    api.HandleFunc("/", post).Methods(http.MethodPost)
    api.HandleFunc("/", put).Methods(http.MethodPut)
    api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)
    api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
    api.HandleFunc("/menu/{parentId}", getMenuView).Methods(http.MethodGet)
    log.Fatal(http.ListenAndServe(":5021", r))
}