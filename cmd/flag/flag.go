package flag

import (
	"flag"
	"fmt"
)

type ConfigFlags struct {
	ConfigPath string
}

func ParseFlags() *ConfigFlags {
	config := flag.String("config", "", "Path to a configuration file.")

	flag.Usage = func() {
		fmt.Println("Usage terramate-bootstrap [options]")
		fmt.Println("Options")
		fmt.Println("   -config=./path/to/config.yaml     Pass config file to create stacks")
		fmt.Println("   -h, --help                        Show help information.")
	}

	flag.Parse()

	return &ConfigFlags{
		*config,
	}
}
