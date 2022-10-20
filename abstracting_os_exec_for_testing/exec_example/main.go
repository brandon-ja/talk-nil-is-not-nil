package main

import (
	"log"
	"os/exec"
)

func main() {
	clean, err := branchIsClean()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("clean: %v\n", clean)
}

// START_FUNC OMIT
func branchIsClean() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	status, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return len(status) == 0, nil
}

// END_FUNC OMIT
