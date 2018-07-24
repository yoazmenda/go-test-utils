package fileutil

import (
	"fmt"
	"os"
)

// this module is here to check that only necessary modules (the ones that are actually imported are downloaded from the repository and not the whole repository)

func Experiment() int {
	fmt.Println("You Should not use me")
	os.Exit(-1)
	return -1
}