package app

import (
	"net/http"
)

// Router struct
type Router struct {
	Mux *http.ServeMux
}

const BasePattern = "/api"

func (router *Router) Init() {
	router.Mux = http.NewServeMux()
	router.registerRoutes()
}

func (router *Router) registerRoutes() {
	h := &FileHandler{}
	router.Mux.HandleFunc(BasePattern+"/files", h.PostFile)
}
