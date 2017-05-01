package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
}

func NewServer(address string) *Server {
	return &Server{
		Addr: address,
	}
}

func (s *Server) Start() {
	r := mux.NewRouter()
	r.NewRoute().Methods("POST").Path("/v1/builds").
		Headers("Content-Type", "application/json").
		HandlerFunc(s.handleNewBuild)
	r.NewRoute().Methods("PATCH").Path("/v1/builds/{id}").
		Headers("Content-Type", "application/json").
		HandlerFunc(s.handlePatchBuild)
	r.NewRoute().Methods("GET").Path("/v1/builds/{id}").
		HandlerFunc(s.handleGetBuild)
	r.NewRoute().Methods("PUT").Path("/v1/builds/{id}").
		HandlerFunc(s.handlePutBuild)

	ss := &http.Server{
		Addr:    s.Addr,
		Handler: r,
	}
	log.Fatal(ss.ListenAndServe())
}
