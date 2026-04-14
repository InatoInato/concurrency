package main

import (
	"concurrency/internal/db"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Querier interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type Server struct {
	queries Querier
}


func main() {
	connStr := "postgres://go_user:golang_123@db:5432/golang?sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{
		queries: db.New(conn),
	}

	r := chi.NewRouter()

	r.Post("/users", server.createUser)
	r.Get("/users/{id}", server.getUser)
	r.Put("/users/{id}", server.updateUser)
	r.Delete("/users/{id}", server.deleteUser)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req db.CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := s.queries.CreateUser(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	user, err := s.queries.GetUserByID(context.Background(), int32(id))
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var req db.UpdateUserParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	req.ID = int32(id)

	user, err := s.queries.UpdateUser(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	err := s.queries.DeleteUser(context.Background(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}