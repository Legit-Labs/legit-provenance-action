package main

import (
	"fmt"

	"github.com/legit-labs/legit-provenance-action/pkg/legit_provenance_action"
)

func main() {
	env := legit_provenance_action.GetEnv()
	fmt.Printf("%v\n", env)
}
