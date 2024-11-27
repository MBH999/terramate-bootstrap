package tmutils

import (
	"fmt"
	"os/exec"
)

func errorChecks(command string, path string) {
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()

	fmt.Printf("Creating Stack: %s\n", path)

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				fmt.Println("Stack already exists! Continuing...")
			} else {
				fmt.Println("Error: ", err)
			}
		}
		return
	}
	fmt.Printf("Command Output: %v", string(output))
}
