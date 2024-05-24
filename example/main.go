package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"bitbucket.org/zetaactions/pkltest"
)

func main() {
	environment := "local"
	formation := "default"

	// Get the first command-line argument
	if len(os.Args) >= 2 {
		args := os.Args[1:]
		// Set the environment variable to the first argument
		environment = args[0]
		formation = args[1]
	}
	formation, err := pkltest.LoadFromPath(context.Background(), fmt.Sprintf("environment/%s/%s/main.pkl", environment, formation))
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(formation)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(encoded))

	fmt.Sprint(formation.Zasper.Redises[0].Host)
}
