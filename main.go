package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("==== NOW PRINTING ALL AVAILABLE ENVIRONMENT VARIABLES ====")
	for _, env := range os.Environ() {
		log.Println(env)
	}
	log.Println("==== FINISHED PRINTING AVAILABLE ENVIRONMENT VARIABLES ====")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/hello", helloWorldHandler)

	log.Printf("Server starting on port %v\n", os.Getenv("PORT"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
