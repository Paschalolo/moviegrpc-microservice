package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"movieexample.com/movie/internal/controller/movie"
)

// Handler definees a movie handler
type Handler struct {
	ctrl *movie.Controller
}

// New create a new movies http handler
func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl}
}

// GetMovieDetails handles GET movie requests
func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("repository get error :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Println("response ecoded err ", err)
		return
	}

}
