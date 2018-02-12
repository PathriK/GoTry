package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"appengine"
	"appengine/datastore"
)

type response struct {
	ErrorFlag bool        `json:"errorFlag"`
	ErrorMsg  string      `json:"errorMsg"`
	Data      interface{} `json:"data"`
}

// ListHandler gets list of MCQs for Admin display
func ListHandler(w http.ResponseWriter, r *http.Request) {
	resp := response{
		ErrorFlag: false,
		ErrorMsg:  "",
	}
	c := appengine.NewContext(r)
	q := datastore.NewQuery("MCQ").Ancestor(MCQKey(c))
	cnt, err := q.Count(c)
	if err != nil {
		writeError(w, resp, err, http.StatusInternalServerError)
		return
	}

	mcqs := make([]MCQ, 0, cnt)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		writeError(w, resp, err, http.StatusInternalServerError)
		return
	}
	resp.Data = mcqs
	writeResponse(w, resp)
}

func writeError(w http.ResponseWriter, resp response, err error, code int) {
	resp.ErrorFlag = true
	resp.ErrorMsg = err.Error()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	writeResponse(w, resp)
}

func writeResponse(w http.ResponseWriter, resp response) {
	if js, err := json.Marshal(resp); err != nil {
		if resp.ErrorFlag {
			fmt.Fprintln(w, resp.ErrorMsg)
		} else {
			fmt.Fprintln(w, "Marshal Error"+err.Error())
		}
	} else {
		w.Write(js)
	}
}
