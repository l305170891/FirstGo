package main

import (
	"os"
	"os/exec"
)

func main3() {
	c := exec.Command("netstat", "-ano")
	c.Stdout = os.Stdout
	c.Run()
}
