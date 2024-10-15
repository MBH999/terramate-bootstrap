package main

import (
  "terramate-bootstrap/tmutils"
  "terramate-bootstrap/userinput"
  "fmt"
)

func main() {
  versionCheck, err := tmutils.CheckVersion()
  fmt.Println(versionCheck, err)

  regions := userinput.Regions()
  // regionList := strings.Split(regions, ",")
  fmt.Println(regions)

  environments := userinput.Environments()
  fmt.Println(environments)

  tmutils.TerramateCreate(regions, environments)
}
