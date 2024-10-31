package tmglobals

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed templates/state.globals.tm.hcl.tmpl
var embeddedFiles embed.FS

func TMGlobals() {
	// Define the variables for the template
	vars := map[string]interface{}{
		"ResourceGroupName": "rg-test",
	}

	// Read the embedded template file
	templateContent, err := embeddedFiles.ReadFile("templates/state.globals.tm.hcl.tmpl")
	if err != nil {
		fmt.Printf("Error reading embedded template file: %v\n", err)
		return
	}

	// Parse the template content
	tmpl, err := template.New("stateTemplate").Parse(string(templateContent))
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file
	file, err := os.Create("./state.globals.tm.hcl")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the provided variables
	if err := tmpl.Execute(file, vars); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Println("File generated successfully as state.globals.tm.hcl")
}
