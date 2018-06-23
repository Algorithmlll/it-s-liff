package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	fmt.Println("Hello GO")

	router := mux.NewRouter()

	router.HandleFunc("/about", aboutHandler).Methods("GET")
	router.HandleFunc("/", homeGETHandler).Methods("GET")
	router.HandleFunc("/", homePOSTHandler).Methods("POST")

	if error := http.ListenAndServe(":8000", router); error != nil {
		log.Fatal(error)
	}
}

func homePOSTHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	r.ParseForm()
	email := r.PostForm.Get("email")
	lineuserid := r.PostForm.Get("lineuserid")
	linedisplayname := r.PostForm.Get("linedisplayname")

	fmt.Fprintf(w, "land sale agreement")
	fmt.Fprintf(w, email)
	fmt.Fprintf(w, lineuserid)
	fmt.Fprintf(w, linedisplayname)

}

func homeGETHandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
