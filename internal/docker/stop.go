package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func StopContainer(container string) error {
	if container == "" {
		return fmt.Errorf("container name is required")
	}

	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker not found")
	}

	cmd := exec.Command("docker", "stop", container)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop container %s: %w", container, err)
	}

	return nil
}
