package main
import (
    "fmt"
    "os/exec"
)
func main() {
  cmdStr := "docker run -itd srivignesh/kali_gui:v2"
  out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()  
  fmt.Printf("%s", out)
}   
