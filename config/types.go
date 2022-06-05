package funtion

import "net/http"

type Server struct {
	port   string
	router *Router
}

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Middleware func(http.HandlerFunc) http.HandlerFunc
