package tmutils

import (
	"fmt"
	"slices"
	"strings"
	"terramate-bootstrap/types"

	"github.com/charmbracelet/log"
)

func DeployResourceStacks(config types.StacksConfig) {
	fmt.Println("#################################")
	fmt.Println("Creating Stacks for Resources!")
	fmt.Println("#################################")

	for env, envDetails := range config.Environments {
		for region, regionDetails := range config.Regions {
			for resource, resourceDetails := range config.ResourceTypes {
				// Check if the current environment is excluded for this resource
				if resourceDetails.ExcludeEnvironments != nil {
					// Skip if the environment is in the excluded list
					if slices.Contains(resourceDetails.ExcludeEnvironments, env) {
						log.Errorf("Skipping %s for environment %s as it is excluded.", resource, env)
						continue
					}
				}

				// Define tags for the stack
				tags := []string{env, region, resource}
				tagsArray := slices.Concat(envDetails.Tags, regionDetails.Tags, resourceDetails.Tags, tags)
				tagsString := strings.Join(tagsArray, ",")

				// Create paths
				resourceStackPath := "./stacks/" + env + "/" + region + "/" + resource

				// Build the command
				command := "terramate create --tags " + tagsString + " " + resourceStackPath

				// Execute the command with error checking
				errorChecks(command, resourceStackPath)
			}
		}
	}
	fmt.Println("------------------------------------------------------------------------------")
}
