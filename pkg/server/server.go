package server

import (
	"log"
	"net/http"

	"github.com/fajran/buildlog/pkg/buildlog"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string

	buildlog *buildlog.BuildLog
}

func NewServer(address string, buildlog *buildlog.BuildLog) *Server {
	return &Server{
		Addr:     address,
		buildlog: buildlog,
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
