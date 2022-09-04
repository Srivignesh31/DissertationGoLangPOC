package main
import (
    "fmt"
    "os/exec"
)
func main() {
  cmdStr := "docker-compose up"
  out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()  
  fmt.Printf("%s", out)
}   