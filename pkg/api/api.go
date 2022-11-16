package api

import (
	"encoding/json"
	"go-microservice-example/pkg/db/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

//start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/comments", func(r chi.Router) {
		r.Get("/", getComments)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("comments"))
}

type CreateCommentRequest struct {
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
}

type CommentResponse struct {
	Success bool            `json:"success"`
	Error   string          `json:"error"`
	Comment *models.Comment `json:"comment"`
}

func createComment(w http.ResponseWriter, r *http.Request) {
	req := &CreateCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
}