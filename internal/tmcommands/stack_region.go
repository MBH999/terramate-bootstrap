package tmcommands

import (
	"fmt"
	"slices"
	"strings"
	"terramate-bootstrap/types"
)

func DeployRegionStacks(config types.StacksConfig) {
	fmt.Println("#################################")
	fmt.Println("Creating Stacks for Regions!")
	fmt.Println("#################################")

	for env, envDetails := range config.Environments {
		for region, regionDetails := range config.Regions {
			// define tags for stack
			tags := []string{env, region}
			tagsArray := slices.Concat(envDetails.Tags, regionDetails.Tags, tags)
			tagsString := strings.Join(tagsArray, ",")

			// create paths
			regionStackPath := "./stacks/" + env + "/" + region

			// build command
			command := "terramate create --tags " + tagsString + " " + regionStackPath

			errorChecks(command, regionStackPath)
		}
	}
	fmt.Println("------------------------------------------------------------------------------")
}
