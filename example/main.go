package main

import (
	"context"
	"encoding/json"
	"fmt"

	"bitbucket.org/zetaactions/pkltest"
)

func main() {
	formation, err := pkltest.LoadFromPath(context.Background(), "environment/local/main.pkl")
	if err != nil {
		panic(err)
	}

	encoded, _ := json.Marshal(formation.Kosmos)
	fmt.Println(string(encoded))

}
