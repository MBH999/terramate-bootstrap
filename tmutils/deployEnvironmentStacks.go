package tmutils

import (
	"fmt"
	"slices"
	"strings"
	"terramate-bootstrap/types"
)

func DeployEnvironmentStacks(config types.StacksConfig) {
	fmt.Println("#################################")
	fmt.Println("Creating Stacks for Environments!")
	fmt.Println("#################################")

	for env, envDetails := range config.Environments {
		// define tags for stack
		tags := []string{env}
		tagsArray := slices.Concat(envDetails.Tags, tags)
		tagsString := strings.Join(tagsArray, ",")

		// create paths
		envStackPath := "./stacks/" + env

		// build command
		command := "terramate create --tags " + tagsString + " " + envStackPath

		errorChecks(command, envStackPath)
	}
}
