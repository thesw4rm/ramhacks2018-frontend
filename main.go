package main

import (
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
	"ramhacks2018/routes"
)

func main(){
	mainRouter := mux.NewRouter().StrictSlash(false)
	opts := renderer.Options{
		ParseGlobPattern: "./views/*.html",
	}
	Rnd := renderer.New(opts)
	indexRouter := mainRouter.PathPrefix("/").Subrouter().StrictSlash(false)
	routes.SetIndexRoutes(indexRouter, Rnd)

}
