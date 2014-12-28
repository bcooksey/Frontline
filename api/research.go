package api

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func buyAttempts(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Buying...")
}

func attemptResearch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]
    fmt.Fprintf(w, "Researching %s... using %s attempts.", category, r.URL.Query().Get("numAttempts"))
}
