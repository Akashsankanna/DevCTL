package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func RAMUsage() error {
	v, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	fmt.Printf("Total RAM: %.2f GB\n", float64(v.Total)/(1024*1024*1024))
	fmt.Printf("Used RAM : %.2f GB\n", float64(v.Used)/(1024*1024*1024))
	fmt.Printf("Usage    : %.2f%%\n", v.UsedPercent)

	return nil
}
