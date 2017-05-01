package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io"
)

type buildRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`

	Status string `json:"status"`
}

func (s *Server) handleNewBuild(w http.ResponseWriter, r *http.Request) {
	var req buildRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := 12345
	fmt.Fprintf(w, "New build: key=%s, name=%s, status=%s => id=%d", req.Key, req.Name, req.Status, id)
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
