package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func ShowLogs(container string) error {
	if container == "" {
		return fmt.Errorf("container name is required")
	}

	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker binary not found")
	}

	cmd := exec.Command("docker", "logs", container)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to get logs for container %s: %w", container, err)
	}

	return nil
}
