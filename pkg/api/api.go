package api

import (
	"encoding/json"
	"go-microservice-example/pkg/db/models"
	"log"
	"net/http"
	"strconv"

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

	//routes for our service
	r.Route("/comments", func(r chi.Router) {
		r.Post("/", createComment)
		r.Get("/", getComments)
		r.Get("/{commentID}", getCommentByID)
		r.Put("/{commentID}", updateCommentByID)
	})

	//test route to make sure everything works
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

// -- Requests and Responses

type CommentRequest struct {
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
}

type CommentResponse struct {
	Success bool            `json:"success"`
	Error   string          `json:"error"`
	Comment *models.Comment `json:"comment"`
}

type CommentsResponse struct {
	Success  bool              `json:"success"`
	Error    string            `json:"error"`
	Comments []*models.Comment `json:"comments"`
}

//-- UTILS --

func handleErr(w http.ResponseWriter, err error) {
	res := &CommentResponse{
		Success: false,
		Error:   err.Error(),
		Comment: nil,
	}
	err = json.NewEncoder(w).Encode(res)
	//if there's an error with encoding handle it
	if err != nil {
		log.Printf("error sending response %v\n", err)
	}
	//return a bad request and exist the function
	w.WriteHeader(http.StatusBadRequest)
}

func handleDBFromContextErr(w http.ResponseWriter) {
	res := &CommentResponse{
		Success: false,
		Error:   "could not get the DB from context",
		Comment: nil,
	}
	err := json.NewEncoder(w).Encode(res)
	//if there's an error with encoding handle it
	if err != nil {
		log.Printf("error sending response %v\n", err)
	}
	//return a bad request and exist the function
	w.WriteHeader(http.StatusBadRequest)
}

// -- handle routes

func createComment(w http.ResponseWriter, r *http.Request) {
	//get the request body and decode it
	req := &CommentRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	//if there's an error with decoding the information
	//send a response with an error
	if err != nil {
		handleErr(w, err)
		return
	}
	//get the db from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	//if we can't get the db let's handle the error
	//and send an adequate response
	if !ok {
		handleDBFromContextErr(w)
		return
	}
	//if we can get the db then
	comment, err := models.CreateComment(pgdb, &models.Comment{
		Comment: req.Comment,
		UserID:  req.UserID,
	})
	if err != nil {
		handleErr(w, err)
		return
	}
	//everything is good
	//let's return a positive response
	res := &CommentResponse{
		Success: true,
		Error:   "",
		Comment: comment,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating comment %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	//get db from ctx
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		handleDBFromContextErr(w)
		return
	}
	//call models package to access the database and return the comments
	comments, err := models.GetComments(pgdb)
	if err != nil {
		handleErr(w, err)
		return
	}
	//positive response
	res := &CommentsResponse{
		Success:  true,
		Error:    "",
		Comments: comments,
	}
	//encode the positive response to json and send it back
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getCommentByID(w http.ResponseWriter, r *http.Request) {
	//get the id from the URL parameter
	//alternatively you could use a URL query
	commentID := chi.URLParam(r, "commentID")

	//get the db from ctx
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		handleDBFromContextErr(w)
		return
	}

	//get the comment from the DB
	comment, err := models.GetComment(pgdb, commentID)
	if err != nil {
		handleErr(w, err)
		return
	}

	//if the retrieval from the db was successful send the data
	res := &CommentResponse{
		Success: true,
		Error:   "",
		Comment: comment,
	}
	//encode the positive response to json and send it back
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateCommentByID(w http.ResponseWriter, r *http.Request) {
	//get the data from the request
	req := &CommentRequest{}
	//decode the data
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		handleErr(w, err)
		return
	}
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		handleDBFromContextErr(w)
		return
	}
	//get the commentID to know what comment to modify
	commentID := chi.URLParam(r, "commentID")
	//we get a string but we need to send an int so we convert it
	intCommentID, err := strconv.ParseInt(commentID, 10, 64)
	if err != nil {
		handleErr(w, err)
		return
	}

	//update the comment
	comment, err := models.UpdateComment(pgdb, &models.Comment{
		ID:      intCommentID,
		Comment: req.Comment,
		UserID:  req.UserID,
	})
	if err != nil {
		handleErr(w, err)
	}
	//return successful response
	res := &CommentResponse{
		Success: true,
		Error:   "",
		Comment: comment,
	}
	//send the encoded response to responsewriter
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//send a 200 response
	w.WriteHeader(http.StatusOK)
}
