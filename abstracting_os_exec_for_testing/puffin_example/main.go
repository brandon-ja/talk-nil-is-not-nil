package main

import (
	"fmt"
	"github.com/weave-lab/puffin"
)

func main() {
	cmd := puffin.NewOsExec().Command("go", "test", "-test.v", "-test.run", `^\QTest_branchIsClean\E$`)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(out))
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
