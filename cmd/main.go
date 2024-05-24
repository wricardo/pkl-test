package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pkltest "bitbucket.org/zetaactions/pkltest"
)

func main() {
	http.HandleFunc("/formation/", formationHandler)

	// File server to serve static files from the "environment" directory
	environmentDir := "."
	fs := http.FileServer(http.Dir(environmentDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formationHandler(w http.ResponseWriter, r *http.Request) {
	formation := r.URL.Path[len("/formation/"):]

	// fmt.Fprintf(w, "Formation: %s", formation)
	f, err := pkltest.LoadFromPath(context.Background(), fmt.Sprintf("%s/%s/main.pkl", "environment", formation))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	encoded, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(encoded))

}
