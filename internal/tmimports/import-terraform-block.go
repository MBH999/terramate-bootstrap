package tmimports

import (
	"embed"
	"log"
	"os"
	"path/filepath"
)

// Step 1: Embed the file using embed.FS
//
//go:embed assets/azurerm/_generated_azurerm_terraform.tm.hcl
var terraformBlock embed.FS

func ImportTerraformBlock() {
	// Step 2: Define the destination directory and file paths
	destinationDir := "./imports/azurerm"
	destinationFile := filepath.Join(destinationDir, "_generated_azurerm_terraform.tm.hcl")

	// Step 3: Read the embedded file content from embed.FS as []byte
	azurermTerraformFile, err := terraformBlock.ReadFile("assets/azurerm/_generated_azurerm_terraform.tm.hcl")
	if err != nil {
		log.Fatalf("Failed to read embedded file: %v", err)
	}

	// Step 4: Create the destination directory if it doesn't exist
	err = os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Step 5: Write the embedded file content to the destination file
	err = os.WriteFile(destinationFile, azurermTerraformFile, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Printf("File successfully written to %s", destinationFile)
}
