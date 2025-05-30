package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ollama/ollama/api"
)

const OLLAMA_BASE_URL = "http://localhost:11434"

func listModels(w http.ResponseWriter, r *http.Request) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create client: %v", err), http.StatusInternalServerError)
		return
	}

	list, err := client.List(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list models: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	for _, model := range list.Models {
		fmt.Fprintf(w, "<li>%s</li>", model.Name)
	}
}

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/api/models", listModels)

	fmt.Println("starting server on port 8020")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatalln(err)
	}
}
