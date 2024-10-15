package userinput

import (
  "fmt"
  "strings"
)


func Regions() []string {
  var regions string
  fmt.Println("What regions are you deploying to (comma seperated list)?: ")
  fmt.Scanln(&regions)

  return strings.Split(regions, ",")
}

func Environments() []string {
  var environments string
  fmt.Println("What environments are you deploying? (e.g. prd,dev,uat - comma seperated list)")
  fmt.Scanln(&environments)
  return strings.Split(environments, ",")
}
