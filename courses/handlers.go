package courses

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func readAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get all items"))
}

func readOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get one item"))
}

func create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create"))
}

func update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update"))
}

func delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete"))
}
