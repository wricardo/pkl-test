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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formationHandler(w http.ResponseWriter, r *http.Request) {
	formation := r.URL.Path[len("/formation/"):]

	fmt.Fprintf(w, "Formation: %s", formation)
	f, err := pkltest.LoadFromPath(context.Background(), fmt.Sprintf("%s/%s/main.pkl", "environment", formation))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	encoded, _ := json.Marshal(f)

	fmt.Fprintf(w, "%s", string(encoded))

}
