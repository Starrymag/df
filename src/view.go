package main

import (
	"fmt"
	"math"
	"strconv"
)

func printRes(m []Mount) {
	fmt.Printf("Файл. система           1К-блоков Использовано Доступно Использовано%% Смонтировано в\n")
	for _, mount := range m {
		device := mount.Device
		mountPoint := mount.Mountpoint
		total := mount.Total
		free := mount.Free
		used := mount.Used
		usagePercent := "-"
		if total != 0 {
			usagePercent = strconv.FormatFloat(math.Ceil(((float64(used)/float64(total))) * 100), 'f', 0, 64) + "%"
		}
		// do block multiple 1K size
		blocks := mount.Blocks * (mount.BlockSize / 1024)
		fmt.Printf("%-23s %9d %12d %8d %13s %-25s\n", device, blocks, used, free, usagePercent, mountPoint)	
	}
}