package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maximepeschard/statuscode/status"
)

// GetStatus handles a request to describe a status from its code.
func GetStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s, err := status.Get(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// ListStatus handles a request to list status codes, with optional filters.
func ListStatus(w http.ResponseWriter, r *http.Request) {
	class := r.FormValue("class")

	var s []status.Status
	var err error
	if class != "" {
		s, err = status.Filter(func(s status.Status) bool {
			match, err := regexp.MatchString(`^`+class+`.*`, s.Class)
			return err == nil && match
		})
	} else {
		s, err = status.All()
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
