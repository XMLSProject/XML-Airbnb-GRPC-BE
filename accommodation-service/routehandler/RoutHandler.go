package routehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}
func Cao(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	a, _ := json.Marshal("cao")
	w.Write(a)

}

// func Routing(router *mux.Router, handler *handler.UserHandler) {
// 	router.HandleFunc("/users/{id}", handler.Get).Methods("GET")
// 	router.HandleFunc("/users", handler.Register).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/auth", handler.Login).Methods("POST")
// 	router.HandleFunc("/api", home).Methods("GET")
// 	router.HandleFunc("/login", handler.Loginn).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/home", handler.Home).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/api/cao", Cao).Methods("GET", "OPTIONS")
// }
