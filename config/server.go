package funtion

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func CreateServer(port string) *Server {
	router := &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
	fmt.Println("Server listening in the port ", port)
	return &Server{port: port, router: router}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {

	allowedMethoes := []string{"GET", "POST", "DELETE", "PUT"}

	if !sliceContain(method, allowedMethoes) {
		_error := fmt.Sprintf("Method %s is not allowed", method)
		log.Println(errors.New(_error))
		os.Exit(3)
	}

	_, exist := s.router.rules[method]
	// that validate if the main routes existe GET POST PUT DELETE if not that should be created
	if !exist {
		s.router.rules[method] = make(map[string]http.HandlerFunc)
	}
	// We add the second level
	s.router.rules[method][path] = handler

}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	handler, methodExist := r.rules[request.Method][path]
	if !methodExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
