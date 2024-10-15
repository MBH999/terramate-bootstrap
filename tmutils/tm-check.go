package tmutils

import (
  "fmt"
  "os/exec"
)

func CheckVersion() (string, error) {  
 cmd := exec.Command("sh", "-c", "terramate --version")

  output, err := cmd.Output()

  if err != nil {
    fmt.Println("Error: ", err)
    return "", err
  }

  return string(output), nil
}
