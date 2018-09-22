package routes

import (
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
	"net/http"
)

var global_Rnd *renderer.Render
func SetIndexRoutes(router *mux.Router, Rnd *renderer.Render){
	router.PathPrefix("/").HandlerFunc(index_handler)
}

func index_handler(w http.ResponseWriter, r *http.Request){
	global_Rnd.HTML(w, http.StatusOK, "index.html", nil)
}