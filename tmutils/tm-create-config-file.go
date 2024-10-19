package tmutils

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"terramate-bootstrap/types"
)

func CreateTMStructureFromConfig(config types.Config) {
	fmt.Println("Creating terramate structure from config file.")

	var paths []string

	// Loop through regions object
	for region, regionDetails := range config.Config.Regions {
		// Loop through environments object
		for env, envDetails := range config.Config.Environments {
			// Loop through resourceTypes object
			for resource, resourceDetails := range config.Config.ResourceTypes {

				genericTags := []string{region, env, resource}

				// Join all tags together
				tagsArray := slices.Concat(envDetails.Tags, regionDetails.Tags, resourceDetails.Tags, genericTags)

				tagsString := strings.Join(tagsArray, ",")

				// Set variable for the base path
				basePath := "./stacks/" + env + "/" + region
				var fullPath string

				// If exclude environments is set on the resourceDetails object.
				if resourceDetails.ExcludeEnvironments != nil {
					// Loop through the ExcludeEnvironments
					for i := 0; i < len(resourceDetails.ExcludeEnvironments); i++ {
						if resourceDetails.ExcludeEnvironments[i] == env {
							// fullPath := basePath
							// fmt.Println(fullPath)
							// paths = append(paths, fullPath)
							break
						} else {
							fullPath = basePath + "/" + resource
							paths = append(paths, fullPath)
						}
					}
				} else {
					fullPath = basePath
					paths = append(paths, fullPath)
				}

				if fullPath != "" {
					command := "terramate create --tags " + tagsString + " " + fullPath
					fmt.Println(command)
					cmd := exec.Command("sh", "-c", command)

					output, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Println("Error:", err)
						fmt.Printf("Full Error: %s\n", string(output))
						continue
					}

					fmt.Printf("Command Output: %v", string(output))

				}
			}
		}
	}

	fmt.Println(paths)
	os.Exit(0)
}
