package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "go test -coverprofile=coverage.out ../../pkg/Array && go tool cover -html=coverage.out -o coverage.html")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
