package daily

import (
	"fmt"
	"os"

	"os/exec"
)

func Command(name string) {
	cmd := exec.Command("git", "version")
	cmd.Env = os.Environ()
	if whoami, err := cmd.Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(whoami))
	}
}
