package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	fun "github.com/ingjjaa2/goolangBaseServer/config"
)

func main() {
	fmt.Println("=========================================================")
	server := fun.CreateServer(":4400")
	server.Handle("GET", "/", server.AddMiddleware(HandleHome, CheckAuth()))
	server.Listen()
	fmt.Println("=========================================================")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 20)
	fmt.Fprintf(w, "Hello from the very best server!")
}

func CheckAuth() fun.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Checking Authentication")
			f(w, r)
		}
	}
}
