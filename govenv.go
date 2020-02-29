package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Getwd()
	cmd := fmt.Sprintf("export GOPATH='%s'", path)
	file, err := os.Create("activate")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(cmd)
}
