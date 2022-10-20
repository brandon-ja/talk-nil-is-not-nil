package main

import (
	"log"

	"github.com/weave-lab/puffin"
)

func main() {
	clean, err := branchIsClean(puffin.NewOsExec())
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("clean: %v\n", clean)
}

// START_FUNC OMIT
func branchIsClean(exec puffin.Exec) (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	status, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return len(status) == 0, nil
}

// END_FUNC OMIT
