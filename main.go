package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("controller-container", "LocalAgent.SupportedThresholdOperator?")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command execution error: %s\n", err)
		fmt.Printf("Error output: %s\n", stderr.String())
		return
	}

	fmt.Printf("Output: %s\n", stdout.String())
}
