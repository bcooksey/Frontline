package api

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
    v1Api := router.PathPrefix("/api/v1").Subrouter()

    v1Api.HandleFunc("/research/attempts", buyAttempts).Methods("POST")
    v1Api.HandleFunc("/research/attempt/{category}", attemptResearch).Methods("DELETE")
}
