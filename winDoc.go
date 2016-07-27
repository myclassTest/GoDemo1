package main

import (
    "os"
    "os/exec"
)

func main() {
    cmd := "ping"
    host := "www.baidu.com"
    c := exec.Command(cmd, host, "-n", "6")
    c.Stdout = os.Stdout
    c.Run()
}
