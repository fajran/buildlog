package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"io"
)

type buildRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`

	Status string `json:"status"`
}

type Build struct {
	Id int `json:"id"`

	Key  string `json:"key"`
	Name string `json:"name"`

	Status   string   `json:"status"`
	Started  *iso8601 `json:"started"`
	Finished *iso8601 `json:"finished"`
}

func validateBuildRequest(req *buildRequest) error {
	if req.Key == "" {
		return fmt.Errorf(`"key" is required`)
	}

	if req.Status == "" {
		req.Status = "STARTED"
	}

	return nil
}

func (s *Server) handleNewBuild(w http.ResponseWriter, r *http.Request) {
	var req buildRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = validateBuildRequest(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	b, err := s.buildlog.Create(req.Key)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	started := iso8601(time.Now())
	build := Build{
		Id:       b.Id,
		Key:      req.Key,
		Name:     req.Name,
		Status:   req.Status,
		Started:  &started,
		Finished: nil,
	}
	err = json.NewEncoder(w).Encode(build)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) handleGetBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "id=%s", id)
}

func (s *Server) handlePatchBuild(w http.ResponseWriter, r *http.Request) {
	var req buildRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (s *Server) handlePutBuild(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Printf("Part: name=%s, filename=%s, type=%s\n", p.FormName(), p.FileName(), p.Header.Get("Content-Type"))
	}
}
