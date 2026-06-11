package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func ListContainers() error {
	cmd := exec.Command("docker", "ps")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute docker ps: %w", err)
	}

	return nil
}
