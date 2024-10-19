package tmutils

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckVersion() {
	cmd := exec.Command("sh", "-c", "terramate --version")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking terramate version. Please verify it is installed.")
		fmt.Println("Please check the terramate documentation for details: https://terramate.io/docs/cli/installation")
		fmt.Println("Error Details: ", err)
		os.Exit(1)
	}

	fmt.Println("Current Terramate Version: " + string(output))
}
