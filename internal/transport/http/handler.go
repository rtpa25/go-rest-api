package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//*Stores pointer to our comments service
type Handler struct {
	Router *mux.Router
}

//*Returns a pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}

//*sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
