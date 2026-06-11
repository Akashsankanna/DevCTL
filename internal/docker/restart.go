package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func RestartContainer(container string) error {
	if container == "" {
		return fmt.Errorf("container name is required")
	}

	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker not found")
	}

	cmd := exec.Command("docker", "restart", container)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to restart container %s: %w", container, err)
	}

	return nil
}
