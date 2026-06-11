package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func StartContainer(container string) error {
	if container == "" {
		return fmt.Errorf("container name is required")
	}

	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker not found")
	}

	cmd := exec.Command("docker", "start", container)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start container %s: %w", container, err)
	}

	return nil
}
