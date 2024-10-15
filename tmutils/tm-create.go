package tmutils

import (
  "fmt"
  "os/exec"
)

func TerramateCreate( regions []string, environments []string)  {
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
        fmt.Println("Full Error: %s\n", string(output))
        continue
      }

      fmt.Printf("Command Output: %v", string(output))
    }
  } 
}
