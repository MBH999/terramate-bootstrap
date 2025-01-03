package tmcommands

import (
	"os/exec"

	"github.com/charmbracelet/log"
)

func errorChecks(command string, path string) {
	cmd := exec.Command("sh", "-c", command)

	_, err := cmd.CombinedOutput()

	log.Info("Creating Stack:", "path", path)

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				log.Warn("Stack already exists! Continuing...")
			} else {
				log.Error("Error: ", err)
			}
		}
		return
	}
}
