package main

import (
	"flag"
	"fmt"
	"terramate-bootstrap/tmutils"
	"terramate-bootstrap/userinput"
)

func main() {
	tmutils.CheckVersion()

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
