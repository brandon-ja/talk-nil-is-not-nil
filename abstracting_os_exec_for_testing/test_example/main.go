package main

import (
	"github.com/weave-lab/puffin"
	"log"
)

func main() {
	clean, err := branchIsClean(puffin.NewOsExec())
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("clean: %v\n", clean)
}

func branchIsClean(exec puffin.Exec) (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	status, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return len(status) == 0, nil
}
