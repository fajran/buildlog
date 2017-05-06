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
		HandlerFunc(s.handleNewBuild)
	r.NewRoute().Methods("GET").Path("/v1/builds/{id}").
		HandlerFunc(s.handleGetBuild)
	r.NewRoute().Methods("POST").Path("/v1/builds/{id}/logs").
		HandlerFunc(s.handlePostLog)
	r.NewRoute().Methods("GET").Path("/v1/builds/{buildId}/logs/{id}/metadata").
		HandlerFunc(s.handleGetLogMetadata)

	ss := &http.Server{
		Addr:    s.Addr,
		Handler: r,
	}
	log.Fatal(ss.ListenAndServe())
}
