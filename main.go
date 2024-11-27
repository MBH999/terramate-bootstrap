package main

import (
	"flag"
	"fmt"
	"terramate-bootstrap/fileutils"
	"terramate-bootstrap/tmimports"
	"terramate-bootstrap/tmtemplates"
	"terramate-bootstrap/tmutils"
)

func main() {
	// check if terramate is installed
	tmutils.CheckVersion()

	// flag for config file
	use_yaml_config := flag.String("config", "", "Path to a configuration file.")

	// flag for help
	flag.Usage = func() {
		fmt.Println("Usage terramate-bootstrap [options]")
		fmt.Println("Options")
		fmt.Println("   -config=./path/to/config.yaml     Pass config file to create stacks")
		fmt.Println("   -h, --help                        Show help information.")
	}

	// parse flags
	flag.Parse()

	// process config file
	if *use_yaml_config != "" {
		config, err := fileutils.ParseConfigFile(use_yaml_config)
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
			tmutils.DeployEnvironmentStacks(config.Stacks)
		}

		if config.Stacks.DeployRegionStacks {
			tmutils.DeployRegionStacks(config.Stacks)
		}

		tmutils.DeployResourceStacks(config.Stacks)
	}
}
