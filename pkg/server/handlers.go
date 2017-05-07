package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Build struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

type Log struct {
	Id      int `json:"id"`
	BuildId int `json:"buildId"`

	Type                 string `json:"type"`
	ContentType          string `json:"contentType"`
	ContentTypeParameter string `json:"contentTypeParameter"`
	Size                 int64  `json:"size"`
}

func (s *Server) handleNewBuild(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	key := qs.Get("key")
	if key == "" {
		http.Error(w, `"key" is required`, 400)
		return
	}

	id, err := s.buildlog.Create(key)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b := Build{
		Id:  id,
		Key: key,
	}
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) handleGetBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	build, err := s.buildlog.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if build == nil {
		http.Error(w, "Not found", 404)
		return
	}

	b := Build{
		Id:  build.Id,
		Key: build.Key,
	}

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) handlePostLog(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	t := qs.Get("type")
	if t == "" {
		http.Error(w, `"type" is required`, 400)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	build, err := s.buildlog.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if build == nil {
		http.Error(w, "Not found", 404)
		return
	}

	ct := r.Header["Content-Type"][0]

	lid, err := build.Log(t, ct, r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	l := Log{
		Id: lid,
	}

	err = json.NewEncoder(w).Encode(l)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) handleGetLogMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}
	buildId, err := strconv.Atoi(vars["buildId"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	build, err := s.buildlog.Get(buildId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if build == nil {
		http.Error(w, "Not found", 404)
		return
	}

	data, err := build.GetLog(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if data == nil {
		http.Error(w, "Not found", 404)
		return
	}

	jsonData := Log{
		Id:                   data.Id,
		BuildId:              data.Build.Id,
		Type:                 data.Type,
		ContentType:          data.ContentType,
		ContentTypeParameter: data.ContentTypeParameter,
		Size:                 data.Size,
	}
	err = json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) handleGetLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}
	buildId, err := strconv.Atoi(vars["buildId"])
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	build, err := s.buildlog.Get(buildId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if build == nil {
		http.Error(w, "Not found", 404)
		return
	}

	data, err := build.GetLog(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if data == nil {
		http.Error(w, "Not found", 404)
		return
	}

	dr, err := data.Read()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ct := createContentType(data.ContentType, data.ContentTypeParameter)
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", data.Size))
	io.CopyN(w, dr, data.Size)
}

func createContentType(contentType, parameter string) string {
	var b bytes.Buffer
	b.WriteString(contentType)
	if parameter != "" {
		b.WriteString("; ")
		b.WriteString(parameter)
	}
	return b.String()
}
