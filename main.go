package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"terramate-bootstrap/tmimports"
	"terramate-bootstrap/tmutils"
	"terramate-bootstrap/types"
	"terramate-bootstrap/userinput"

	"gopkg.in/yaml.v3"
)

func main() {
	tmutils.CheckVersion()

	tmimports.ImportProvider()
	tmimports.ImportTerraformBlock()
	tmimports.ImportsFile()

	use_yaml_config := flag.String("config", "", "Path to a configuration file.")
	caf_landing_zone := flag.Bool("clz", false, "Deploy CAF Landing Zone Structure.")
	environments := flag.Bool("env", false, "Deploy Environment Stucture")

	flag.Usage = func() {
		fmt.Println("Usage terramate-bootstrap [options]")
		fmt.Println("Options")
		fmt.Println("   -clz          Deploy CAF Landing Zone Structure")
		fmt.Println("   -env          Deploy Environments Structure")
		fmt.Println("   -h, --help    Show help information.")
	}

	flag.Parse()

	if *use_yaml_config != "" {
		file, err := os.Open(*use_yaml_config)
		if err != nil {
			log.Fatalf("Error opening configuration file: %v", err)
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading configuration file: %v", err)
		}

		var config types.Config
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("Error parsing configuration file %v", err)
		}

		// fmt.Printf("Parsed Config: %+v\n", config)

		tmutils.CreateTMStructureFromConfig(config)
	}

	regions := userinput.Regions()

	if *caf_landing_zone {
		fmt.Println("Deploying CAF Landing Zone Structure")

		tmutils.TerramateCreateClz(regions)
	}

	if *environments {
		environments := userinput.Environments()

		tmutils.TerramateCreateEnv(regions, environments)
	}
}
