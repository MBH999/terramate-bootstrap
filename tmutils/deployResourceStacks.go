package tmutils

import (
	"fmt"
	"slices"
	"strings"
	"terramate-bootstrap/types"
)

func DeployResourceStacks(config types.StacksConfig) {
	fmt.Println("#################################")
	fmt.Println("Creating Stacks for Resources!")
	fmt.Println("#################################")

	for env, envDetails := range config.Environments {
		for region, regionDetails := range config.Regions {
			for resource, resourceDetails := range config.ResourceTypes {
				// define tags for stack
				tags := []string{env, region, resource}
				tagsArray := slices.Concat(envDetails.Tags, regionDetails.Tags, resourceDetails.Tags, tags)
				tagsString := strings.Join(tagsArray, ",")

				if resourceDetails.ExcludeEnvironments == nil {
					continue
				}

				for _, excludedEnv := range resourceDetails.ExcludeEnvironments {
					if excludedEnv == env {
						continue
					}

					// create paths
					resourceStackPath := "./stacks/" + env + "/" + region + "/" + resource

					// build command
					command := "terramate create --tags " + tagsString + " " + resourceStackPath

					errorChecks(command, resourceStackPath)
				}
			}
		}
	}
}
