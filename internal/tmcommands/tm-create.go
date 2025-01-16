package tmcommands

import (
	"fmt"
	"os/exec"
)

func TerramateCreateEnv(regions []string, environments []string) {
	for i := 0; i < len(environments); i++ {
		for j := 0; j < len(regions); j++ {
			path := "./stacks/environments/" + environments[i] + "/" + regions[j]

			fmt.Println("Creating Stack: " + path)

			command := "terramate create --tags " + environments[i] + "," + regions[j] + " " + path
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

func TerramateCreateClz(regions []string) {
	landing_zones := []string{"identity", "management", "connectivity"}

	for i := 0; i < len(landing_zones); i++ {
		for j := 0; j < len(regions); j++ {

			path := "./stacks/landing_zones/" + landing_zones[i] + "/" + regions[j]

			fmt.Println("Creating Landing Zone Stack: " + path)

			command := "terramate create --tags landing_zones," + landing_zones[i] + "," + regions[j] + " " + path
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
