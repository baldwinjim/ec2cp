package main

import "fmt"

func main() {
  var srcAmi, srcRegion string
  var err error
  fmt.Print("Enter Source AMI:  ")
  _, err := fmt.Scanln(&srcAmi)
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Print("Enter source region to copy from:  ")
  _, err = fmt.Scanln(&srcRegion)
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println("Copying", srcAmi, "in region", srcRegion)
}
