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
	tmutils.CheckVersion()

	use_yaml_config := flag.String("config", "", "Path to a configuration file.")

	flag.Usage = func() {
		fmt.Println("Usage terramate-bootstrap [options]")
		fmt.Println("Options")
		fmt.Println("   -config=./path/to/config.yaml     Pass config file to create stacks")
		fmt.Println("   -h, --help                        Show help information.")
	}

	flag.Parse()

	if *use_yaml_config != "" {
		config, err := fileutils.ParseConfigFile(use_yaml_config)
		if err != nil {
			fmt.Println("error")
		}

		tmtemplates.TMGlobals(config.Backend)

		tmimports.ImportProvider()
		tmimports.ImportTerraformBlock()
		tmimports.ImportsFile()

		tmutils.CreateTMStructureFromConfig(config.Stacks)
	}
}
