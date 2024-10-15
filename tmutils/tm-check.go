package tmutils

import (
	"fmt"
	"os/exec"
)

func CheckVersion() {
	cmd := exec.Command("sh", "-c", "terramate --version")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Current Terramate Version: " + string(output))
}
