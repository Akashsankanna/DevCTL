package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

func ShowCPU() error {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return fmt.Errorf("failed to get CPU usage: %w", err)
	}

	cores, err := cpu.Counts(true)
	if err != nil {
		return fmt.Errorf("failed to get CPU core count: %w", err)
	}

	fmt.Printf("CPU Usage : %.2f%%\n", percentages[0])
	fmt.Printf("CPU Cores : %d\n", cores)

	return nil
}
