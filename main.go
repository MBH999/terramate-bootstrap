package main

import (
	"fmt"
	"terramate-bootstrap/cmd/root"
	"terramate-bootstrap/internal/parseconfig"
	"terramate-bootstrap/internal/tmcommands"
	"terramate-bootstrap/internal/tmimports"
	"terramate-bootstrap/internal/tmtemplates"
)

func main() {
	root.Init()

	// process config file
	if *use_yaml_config != "" {
		config, err := parseconfig.ParseConfigFile(use_yaml_config)
		if err != nil {
			fmt.Println("error")
		}

		tmtemplates.TMGlobals(config.Backend)

		tmimports.ImportProvider()
		tmimports.ImportTerraformBlock()
		tmimports.ImportsFile()

		fmt.Println("Creating terramate structure from config file.")
		fmt.Print(config.Stacks)
		if config.Stacks.DeployEnvironmentStacks {
			tmcommands.DeployEnvironmentStacks(config.Stacks)
		}

		if config.Stacks.DeployRegionStacks {
			tmcommands.DeployRegionStacks(config.Stacks)
		}

		tmcommands.DeployResourceStacks(config.Stacks)
	}
}
