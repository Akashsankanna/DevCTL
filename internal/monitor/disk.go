package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func DiskUsage() error {
	usage, err := disk.Usage("C:\\")
	if err != nil {
		return err
	}

	fmt.Printf("Total Disk: %.2f GB\n", float64(usage.Total)/(1024*1024*1024))
	fmt.Printf("Used Disk : %.2f GB\n", float64(usage.Used)/(1024*1024*1024))
	fmt.Printf("Free Disk : %.2f GB\n", float64(usage.Free)/(1024*1024*1024))
	fmt.Printf("Usage     : %.2f%%\n", usage.UsedPercent)

	return nil
}
