package main
import (
    "fmt"
    "os/exec"
)
func main() {
//     run a kali linux vm which has remote access
  cmdStr := "docker run -p 25900:5900 -p 25901:5901 --name kali -itd srivignesh/kali_gui:v2"
  out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()  
  fmt.Printf("%s", out)
}   
