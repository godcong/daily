package daily

import (
	"fmt"

	"os/exec"
)

func Command(name string) {
	cmd := exec.Command("git", "version")
	if whoami, err := cmd.Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(whoami))
	}
}
